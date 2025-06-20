package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) ListArticle(ctx context.Context, input types.ListArticleInput) (types.ListArticleOutput, error) {
	return s.Repository.ListArticle(ctx, input)
}
