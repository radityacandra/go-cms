package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/model"
)

func (r *Repository) CreateArticle(ctx context.Context, input model.Article) error {
	db := r.GetTransaction(ctx)

	_, err := db.NamedExecContext(ctx, `
		INSERT INTO public.articles(id, content, title, author_id, parent_id, status, created_by, created_at)
		VALUES(:id, :content, :title, :author_id, :parent_id, :status, :created_by, :created_at)`, &input)

	if len(input.Tags) > 0 {
		for i, articleTag := range input.Tags {
			// upsert tags
			var tagId string
			err = db.QueryRowxContext(ctx, `
				SELECT id
				FROM public.tags
				WHERE
					is_deleted = FALSE
					AND name = $1
			`, articleTag.Tag.Name).Scan(&tagId)
			if err != nil {
				_, err = db.NamedExecContext(ctx, `
					INSERT INTO public.tags(id, name, trending_score, created_by, created_at)
					VALUES(:id, :name, :trending_score, :created_by, :created_at)
				`, &articleTag.Tag)
				if err != nil {
					return err
				}

				tagId = articleTag.TagId
			} else {
				_, err = db.NamedExecContext(ctx, `
					UPDATE public.tags
					SET
						updated_at = :created_at,
						updated_by = :created_by,
						usage_count = usage_count + 1
					WHERE
						is_deleted = FALSE
						AND name = :name
				`, &articleTag.Tag)
				if err != nil {
					return err
				}
			}

			articleTag.TagId = tagId

			_, err := db.NamedExecContext(ctx, `
				INSERT INTO public.article_tags(id, article_id, tag_id, created_by, created_at)
				VALUES(:id, :article_id, :tag_id, :created_by, :created_at)`, &articleTag)
			if err != nil {
				return err
			}

			input.Tags[i].TagId = tagId
			input.Tags[i].Tag.Id = tagId
		}
	}

	return err
}
