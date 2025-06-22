package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/radityacandra/go-cms/internal/application/article/repository"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

type IService interface {
	ListArticle(ctx context.Context, input types.ListArticleInput) (types.ListArticleOutput, error)
	CreateArticle(ctx context.Context, input types.CreateArticleInput) (string, error)
	DetailArticle(ctx context.Context, articleId, userId string) (types.DetailArticleOutput, error)
	CreateArticleRevision(ctx context.Context, input types.CreateArticleRevisionInput) (string, error)
	DetailArticleRevision(ctx context.Context, articleId, revisionId string) (types.DetailArticleRevisionOutput, error)
	CalculateTagAssociations(ctx context.Context, article model.Article) error
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
