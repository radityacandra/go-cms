package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/repository"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

type IService interface {
	ListTag(ctx context.Context, input types.ListTagInput) (types.ListTagOutput, error)
	CreateTag(ctx context.Context, input types.CreateTagInput) (string, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
