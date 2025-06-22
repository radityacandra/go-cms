package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/radityacandra/go-cms/internal/application/article/types"
	"github.com/radityacandra/go-cms/pkg/database"
)

type IRepository interface {
	ListArticle(ctx context.Context, input types.ListArticleInput) (types.ListArticleOutput, error)
	CreateArticle(ctx context.Context, input model.Article) error
	FindArticleByIdAndOptionalStatus(ctx context.Context, articleId, status string) (*model.Article, error)
	UpdateArticle(ctx context.Context, input model.Article) error
	FindArticleRevisionByIdAndArticleId(ctx context.Context, articleId, id string) (*model.Article, error)
	CountArticleContainingTags(ctx context.Context, tagIds []string) (int64, error)
	UpsertTagAssociation(ctx context.Context, input types.UpsertTagAssociationInput) error
	database.ITransaction
}

type Repository struct {
	database.ITransaction
	Db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		Db:           db,
		ITransaction: db,
	}
}
