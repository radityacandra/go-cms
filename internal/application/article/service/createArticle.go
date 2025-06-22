package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	tagModel "github.com/radityacandra/go-cms/internal/application/tag/model"
)

func (s *Service) CreateArticle(ctx context.Context, input types.CreateArticleInput) (string, error) {
	articleId := uuid.NewString()
	article := model.NewArticle(articleId, input.Content, input.Title, input.UserId,
		input.Status, input.UserId, time.Now().UnixMilli())

	if input.Tags != nil && len(*input.Tags) > 0 {
		for _, tagInput := range *input.Tags {
			tag := tagModel.NewTag(uuid.NewString(), tagInput, input.UserId, time.Now().UnixMilli())

			article.Tags = append(article.Tags, model.ArticleTag{
				Id:        uuid.NewString(),
				ArticleId: articleId,
				TagId:     tag.Id,
				CreatedBy: input.UserId,
				CreatedAt: time.Now().UnixMilli(),
				Tag:       *tag,
			})
		}
	}

	txCtx, err := s.Repository.BeginTransaction(ctx)
	if err != nil {
		return "", err
	}

	// if transaction is created inside the function, let it commit inside
	if !s.Repository.IsTransaction(ctx) {
		defer func() {
			s.Repository.CommitOrRollbackTransaction(txCtx, err)
		}()
	}

	err = s.Repository.CreateArticle(txCtx, *article)
	if err != nil {
		return "", err
	}

	go func() error {
		ctx := context.Background()
		ctxDeadline, cancel := context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()

		err := s.CalculateTagAssociations(ctxDeadline, *article)

		return err
	}()

	return articleId, nil
}
