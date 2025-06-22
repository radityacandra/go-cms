package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/model"
)

func (r *Repository) UpdateTag(ctx context.Context, input model.Tag) error {
	_, err := r.Db.NamedExecContext(ctx, `
		UPDATE public.tags
		SET
			name = :name,
			trending_score = :trending_score,
			usage_count = :usage_count,
			updated_by = :updated_by,
			updated_at = :updated_at
		WHERE
			is_deleted = FALSE
			AND id = :id
	`, &input)

	return err
}
