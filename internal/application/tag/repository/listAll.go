package repository

import (
	"context"

	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

func (r *Repository) ListAll(ctx context.Context) (types.ListActiveTagOutput, error) {
	rows, err := r.Db.QueryxContext(ctx, `
		SELECT id, name
		FROM public.tags
		WHERE is_deleted = FALSE
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output := types.ListActiveTagOutput{}
	for rows.Next() {
		var item types.ListTagRepoItem
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		output = append(output, api.TagListGetResponseItem(item))
	}

	return output, nil
}
