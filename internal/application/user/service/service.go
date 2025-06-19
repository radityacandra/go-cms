package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/repository"
	"github.com/radityacandra/go-cms/internal/application/user/types"
)

type IService interface {
	RegisterUser(ctx context.Context, input types.RegisterUserInput) (types.RegisterUserOutput, error)
	AssignRole(ctx context.Context, userId, roleName string) error
	DetailUser(ctx context.Context, userId string) (types.DetailUserOutput, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
