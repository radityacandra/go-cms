package repository

import "context"

func (r *Repository) CountArticleContainingTags(ctx context.Context, tagIds []string) (int64, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT
			COUNT(1) total_article
		FROM (
			SELECT
				article_id,
				COUNT(1) count_matched_tags
			FROM public.article_tags
			WHERE 
				is_deleted = FALSE
				AND tag_id = ANY($1)
			GROUP BY article_id
		) a
		WHERE a.count_matched_tags = $2
	`, tagIds, len(tagIds))
	if row.Err() != nil {
		return 0, row.Err()
	}

	var countArticle int64
	if err := row.Scan(&countArticle); err != nil {
		return 0, err
	}

	return countArticle, nil
}
