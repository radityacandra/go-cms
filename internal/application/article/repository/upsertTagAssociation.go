package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (r *Repository) UpsertTagAssociation(ctx context.Context, input types.UpsertTagAssociationInput) error {
	var id string
	err := r.Db.QueryRowxContext(ctx, `
		SELECT
			id
		FROM public.tag_associations
		WHERE
			tag1_id = $1
			AND tag2_id = $2
			AND is_deleted = FALSE
	`, input.Tag1Id, input.Tag2Id).Scan(&id)
	if err != nil {
		_, err := r.Db.NamedExecContext(ctx, `
			INSERT INTO public.tag_associations(id, tag1_id, tag2_id, score, created_by, created_at)
			VALUES(:id, :tag1_id, :tag2_id, :score, :created_by, :created_at)
		`, &input)

		return err
	}

	_, err = r.Db.NamedExecContext(ctx, `
		UPDATE public.tag_associations
		SET
			score = :score,
			updated_at = :created_at,
			updated_by = :created_by
		WHERE
			tag1_id = :tag1_id
			AND tag2_id = :tag2_id
			AND is_deleted = FALSE
	`, &input)

	return err
}
