package service

import (
	"context"

	"github.com/radityacandra/go-cms/api"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) CreateArticleRevision(ctx context.Context, input types.CreateArticleRevisionInput) (string, error) {
	article, err := s.Repository.FindArticleByIdAndOptionalStatus(ctx, input.ArticleId, "")
	if err != nil || article.Status == "archived" {
		return "", types.ErrArticleNotFound
	}

	createInput := api.ArticleCreatePostRequest{
		Content: article.Content,
		Title:   article.Title,
		Status:  article.Status,
	}

	if input.Content != nil {
		createInput.Content = *input.Content
	}

	if input.Title != nil {
		createInput.Title = *input.Title
	}

	if input.Status != nil {
		createInput.Status = *input.Status
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

	revisionArticleId, err := s.CreateArticle(txCtx, types.CreateArticleInput{
		ArticleCreatePostRequest: createInput,
		UserId:                   input.UserId,
	})
	if err != nil {
		return "", err
	}

	// de-activate previous main article
	article.DeactivateArticle(revisionArticleId, input.UserId)
	err = s.Repository.UpdateArticle(txCtx, *article)
	if err != nil {
		return "", err
	}

	return revisionArticleId, nil
}
