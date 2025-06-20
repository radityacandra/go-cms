package repository

import (
	"context"
	"fmt"

	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (r *Repository) ListArticle(ctx context.Context, input types.ListArticleInput) (types.ListArticleOutput, error) {
	queryArgs := []interface{}{}

	var statusStmt string
	if input.Status != "all" {
		statusStmt = "AND a.status = $1"
		queryArgs = append(queryArgs, input.Status)
	}

	offset := (input.Page - 1) * input.PageSize

	output := types.ListArticleOutput{
		Data: []types.ListArticleItem{},
		Pagination: types.Pagination{
			Page:     input.Page,
			PageSize: input.PageSize,
		},
	}

	baseSql := fmt.Sprintf(`
		SELECT 
			a.id,
			a.title,
			a.content,
			a.author_id,
			a.status,
			u.full_name author_name
		FROM public.articles a
			JOIN public.users u
				ON a.author_id = u.id
		WHERE
			a.is_deleted = false
			AND u.is_deleted = false
			AND a.parent_id IS NULL
			%s
	`, statusStmt)

	rows, err := r.Db.QueryxContext(ctx, fmt.Sprintf(`
		%s
		LIMIT %d
		OFFSET %d
	`, baseSql, input.PageSize, offset), queryArgs...)
	if err != nil {
		return types.ListArticleOutput{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item types.ListArticleItem
		if err := rows.StructScan(&item); err != nil {
			return types.ListArticleOutput{}, err
		}

		output.Data = append(output.Data, item)
	}

	row := r.Db.QueryRowxContext(ctx, fmt.Sprintf(`
		SELECT
			COUNT(1) total_data
		FROM (
			%s
		) a`, baseSql), queryArgs...)
	if row.Err() != nil {
		return types.ListArticleOutput{}, row.Err()
	}

	var totalData int64
	if err := row.Scan(&totalData); err != nil {
		return types.ListArticleOutput{}, err
	}

	output.Pagination.TotalData = totalData

	return output, nil
}
