package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/model"
)

func (r *Repository) FindArticleRevisionByIdAndArticleId(ctx context.Context, articleId, id string) (*model.Article, error) {
	row := r.Db.QueryRowxContext(ctx, `
		WITH RECURSIVE revision_histories AS (
			SELECT 
				*
			FROM public.articles
			WHERE
				parent_id = $1
				AND is_deleted = FALSE
			
			UNION ALL
			
			SELECT 
				articles.*
			FROM public.articles, revision_histories
			WHERE
				revision_histories.id = articles.parent_id
				AND articles.is_deleted = FALSE
		)
		SELECT
			rh.id, rh.content, rh.title, rh.author_id, rh.parent_id, rh.status, rh.created_by, rh.created_at, u.full_name author_name
		FROM revision_histories rh
			JOIN public.users u
				ON rh.author_id = u.id
		WHERE
			u.is_deleted = FALSE
			AND rh.id = $2
	`, articleId, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var result queryResult
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}

	article := model.Article{
		Id:        result.Id,
		Content:   result.Content,
		Title:     result.Title,
		ParentId:  result.ParentId,
		Status:    result.Status,
		AuthorId:  result.AuthorId,
		CreatedBy: result.CreatedBy,
		CreatedAt: result.CreatedAt,
		Author: model.Author{
			Id:   result.AuthorId,
			Name: result.AuthorName,
		},
	}

	return &article, nil
}
