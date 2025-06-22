package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/model"
)

func (r *Repository) FindTagById(ctx context.Context, tagId string) (*model.Tag, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT id, name, trending_score, usage_count, created_by, created_at
		FROM public.tags
		WHERE
			is_deleted = FALSE
			AND id = $1
	`, tagId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var tag model.Tag
	if err := row.StructScan(&tag); err != nil {
		return nil, err
	}

	rows, err := r.Db.QueryxContext(ctx, `
		SELECT id, tag_id, article_id, created_at
		FROM public.article_tags
		WHERE
			is_deleted = FALSE
			AND tag_id = $1
	`, tagId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item model.TagArticle
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		tag.TagArticles = append(tag.TagArticles, item)
	}

	return &tag, nil
}
