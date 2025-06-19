package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/auth/types"
	"github.com/radityacandra/go-cms/internal/application/user/repository"
)

type IService interface {
	Login(ctx context.Context, input types.LoginInput) (types.LoginOutput, error)
}

type Service struct {
	Repository repository.IRepository
	PrivateKey string
}

func NewService(repository repository.IRepository, privateKey string) *Service {
	return &Service{
		Repository: repository,
		PrivateKey: privateKey,
	}
}
