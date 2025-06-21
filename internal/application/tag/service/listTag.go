package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

func (s *Service) ListTag(ctx context.Context, input types.ListTagInput) (types.ListTagOutput, error) {
	return s.Repository.ListTag(ctx, input)
}
