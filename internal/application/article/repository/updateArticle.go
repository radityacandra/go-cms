package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/model"
)

func (r *Repository) UpdateArticle(ctx context.Context, input model.Article) error {
	db := r.GetTransaction(ctx)

	_, err := db.NamedExecContext(ctx, `
		UPDATE public.articles
		SET
			content = :content,
			title = :title,
			author_id = :author_id,
			parent_id = :parent_id,
			status = :status,
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE
			id = :id
			AND is_deleted = false
	`, &input)

	return err
}
