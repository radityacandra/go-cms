package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/tag/model"
	"github.com/radityacandra/go-cms/internal/application/tag/types"
)

func (s *Service) CreateTag(ctx context.Context, input types.CreateTagInput) (string, error) {
	tag := model.NewTag(uuid.NewString(), input.Name, input.UserId, time.Now().UnixMilli())

	err := s.Repository.CreateTag(ctx, *tag)
	if err != nil {
		return "", err
	}

	return tag.Id, nil
}
