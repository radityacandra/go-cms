package service

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/types"
)

func (s *Service) DetailUser(ctx context.Context, userId string) (types.DetailUserOutput, error) {
	user, err := s.Repository.FindUserById(ctx, userId)
	if err != nil {
		return types.DetailUserOutput{}, types.ErrUserNotFound
	}

	return types.DetailUserOutput{
		Id:       user.Id,
		Username: user.Username,
		Acls:     user.CollectAllAccess(),
	}, nil
}
