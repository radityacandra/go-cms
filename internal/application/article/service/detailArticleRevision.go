package service

import (
	"context"

	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) DetailArticleRevision(ctx context.Context, articleId, revisionId string) (types.DetailArticleRevisionOutput, error) {
	article, err := s.Repository.FindArticleRevisionByIdAndArticleId(ctx, articleId, revisionId)
	if err != nil {
		return types.DetailArticleRevisionOutput{}, err
	}

	return types.DetailArticleRevisionOutput{
		Id:      article.Id,
		Title:   article.Title,
		Content: article.Content,
		Author: api.AuthorSchema{
			Id:   article.AuthorId,
			Name: article.Author.Name,
		},
	}, nil
}
