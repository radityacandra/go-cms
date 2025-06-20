package repository

import (
	"context"
	"fmt"

	"github.com/radityacandra/go-cms/internal/application/article/model"
)

type queryResult struct {
	model.Article
	AuthorName string `db:"author_name"`
}

func (r *Repository) FindArticleByIdAndOptionalStatus(ctx context.Context, articleId string, status string) (*model.Article, error) {
	queryArgs := []interface{}{articleId}

	var statusStmt string
	if status != "" {
		statusStmt = "AND a.status = $2"
		queryArgs = append(queryArgs, status)
	}

	row := r.Db.QueryRowxContext(ctx, fmt.Sprintf(`
		SELECT 
			a.id, a.content, a.title, a.author_id, a.parent_id, a.status, a.created_by, a.created_at, au.full_name author_name
		FROM public.articles a
			JOIN public.users au
				ON a.author_id = au.id
		WHERE
			a.is_deleted = FALSE
			AND au.is_deleted = FALSE
			AND a.parent_id IS NULL
			AND a.id = $1
			%s
	`, statusStmt), queryArgs...)
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

	// find revision histories
	rows, err := r.Db.QueryxContext(ctx, `
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
	`, articleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item queryResult
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		article.ArticleRevisions = append(article.ArticleRevisions, model.Article{
			Id:        item.Id,
			Content:   item.Content,
			Title:     item.Title,
			ParentId:  item.ParentId,
			Status:    item.Status,
			AuthorId:  item.AuthorId,
			CreatedBy: item.CreatedBy,
			CreatedAt: item.CreatedAt,
			Author: model.Author{
				Id:   item.AuthorId,
				Name: item.AuthorName,
			},
		})
	}

	return &article, nil
}
