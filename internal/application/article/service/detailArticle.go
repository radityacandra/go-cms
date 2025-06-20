package service

import (
	"context"

	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) DetailArticle(ctx context.Context, articleId, userId string) (types.DetailArticleOutput, error) {
	status := ""
	if userId == "" {
		status = "published"
	}
	article, err := s.Repository.FindArticleByIdAndOptionalStatus(ctx, articleId, status)
	if err != nil {
		return types.DetailArticleOutput{}, err
	}

	output := types.DetailArticleOutput{
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
		Author: api.AuthorSchema{
			Id:   article.AuthorId,
			Name: article.Author.Name,
		},
		RevisionHistories: []api.RevisionHistorySchema{},
	}

	for _, item := range article.ArticleRevisions {
		output.RevisionHistories = append(output.RevisionHistories, api.RevisionHistorySchema{
			Id:           item.Id,
			AuthorId:     item.AuthorId,
			AuthorName:   item.Author.Name,
			RevisionTime: item.CreatedAt,
		})
	}

	return output, nil
}
