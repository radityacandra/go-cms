package service

import (
	"context"
	"errors"

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
		return types.DetailArticleOutput{}, errors.Join(err, types.ErrArticleNotFound)
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
		Tags:              api.ArticleTagsSchema{},
	}

	for _, item := range article.ArticleRevisions {
		output.RevisionHistories = append(output.RevisionHistories, api.RevisionHistorySchema{
			Id:           item.Id,
			AuthorId:     item.AuthorId,
			AuthorName:   item.Author.Name,
			RevisionTime: item.CreatedAt,
		})
	}

	for _, articleTag := range article.Tags {
		output.Tags = append(output.Tags, articleTag.Tag.Name)
	}

	return output, nil
}
