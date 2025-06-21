package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/model"
)

func (r *Repository) CreateTag(ctx context.Context, input model.Tag) error {
	_, err := r.Db.NamedExecContext(ctx, `
		INSERT INTO public.tags(id, name, popularity_score, created_by, created_at)
		VALUES(:id, :name, :popularity_score, :created_by, :created_at)`, &input)

	return err
}
