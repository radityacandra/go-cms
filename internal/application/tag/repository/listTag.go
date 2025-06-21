package repository

import (
	"context"
	"fmt"

	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

func (r *Repository) ListTag(ctx context.Context, input types.ListTagInput) (types.ListTagOutput, error) {
	baseQuery := `
		SELECT id, name
		FROM public.tags
		WHERE is_deleted = FALSE
	`

	offset := (input.Page - 1) * input.PageSize

	rows, err := r.Db.QueryxContext(ctx, fmt.Sprintf(`
		%s
		LIMIT %d
		OFFSET %d
	`, baseQuery, input.PageSize, offset))
	if err != nil {
		return types.ListTagOutput{}, err
	}
	defer rows.Close()

	output := types.ListTagOutput{
		Data: []api.TagListGetResponseItem{},
		Pagination: api.PaginationSchema{
			Page:     int64(input.Page),
			PageSize: int64(input.PageSize),
		},
	}
	for rows.Next() {
		var item types.ListTagRepoItem
		if err := rows.StructScan(&item); err != nil {
			return types.ListTagOutput{}, err
		}

		output.Data = append(output.Data, api.TagListGetResponseItem(item))
	}

	var totalData int64
	row := r.Db.QueryRowxContext(ctx, fmt.Sprintf(`
		SELECT COUNT(1) total_data FROM (%s) a
	`, baseQuery))
	if row.Err() != nil {
		return types.ListTagOutput{}, row.Err()
	}

	if err := row.Scan(&totalData); err != nil {
		return types.ListTagOutput{}, err
	}
	output.Pagination.TotalData = totalData

	return output, nil
}
