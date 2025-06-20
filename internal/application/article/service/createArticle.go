package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) CreateArticle(ctx context.Context, input types.CreateArticleInput) (string, error) {
	articleId := uuid.NewString()
	article := model.NewArticle(articleId, input.Content, input.Title, input.UserId,
		input.Status, input.UserId, time.Now().UnixMilli())

	txCtx, err := s.Repository.BeginTransaction(ctx)
	if err != nil {
		return "", err
	}

	// if transaction is created inside the function, let it commit inside
	if !s.Repository.IsTransaction(ctx) {
		defer func(err error) {
			s.Repository.CommitOrRollbackTransaction(txCtx, err)
		}(err)
	}

	if err := s.Repository.CreateArticle(txCtx, *article); err != nil {
		return "", err
	}

	return articleId, nil
}
