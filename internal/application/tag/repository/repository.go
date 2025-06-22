package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/model"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
	"github.com/radityacandra/go-cms/pkg/database"
)

type IRepository interface {
	ListTag(ctx context.Context, input types.ListTagInput) (types.ListTagOutput, error)
	CreateTag(ctx context.Context, input model.Tag) error
	FindTagById(ctx context.Context, tagId string) (*model.Tag, error)
	UpdateTag(ctx context.Context, input model.Tag) error
	ListAll(ctx context.Context) (types.ListActiveTagOutput, error)
}

type Repository struct {
	Db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
