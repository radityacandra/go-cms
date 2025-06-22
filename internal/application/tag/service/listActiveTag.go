package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

func (s *Service) ListActiveTag(ctx context.Context) (types.ListActiveTagOutput, error) {
	return s.Repository.ListAll(ctx)
}
