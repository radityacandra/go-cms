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

	return err
}
