package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/model"
	"github.com/radityacandra/go-cms/internal/application/user/types"
	"github.com/radityacandra/go-cms/pkg/database"
)

type IRepository interface {
	FindUserByUsername(ctx context.Context, input types.FindUserByUsernameInput) (*model.User, error)
	InsertUser(ctx context.Context, input model.User) error
	FindRoleByName(ctx context.Context, roleName string) (*model.Role, error)
	InsertUserRole(ctx context.Context, userRole model.UserRole) error
	FindUserById(ctx context.Context, userId string) (*model.User, error)
	database.ITransaction
}

type Repository struct {
	Db *database.DB
	database.ITransaction
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		Db:           db,
		ITransaction: db,
	}
}
