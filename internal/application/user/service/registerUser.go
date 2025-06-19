package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/user/model"
	"github.com/radityacandra/go-cms/internal/application/user/types"
	"github.com/radityacandra/go-cms/pkg/hash"
)

func (s *Service) RegisterUser(ctx context.Context, input types.RegisterUserInput) (types.RegisterUserOutput, error) {
	// find user by username
	_, err := s.Repository.FindUserByUsername(ctx, types.FindUserByUsernameInput{
		Username: input.Username,
	})
	if err == nil {
		err = types.ErrUserConflicting
		return types.RegisterUserOutput{}, err
	}

	if !errors.Is(err, types.ErrUserNotFound) {
		return types.RegisterUserOutput{}, err
	}

	// if not found, save to db
	password, err := hash.GenerateHash(input.Password)
	if err != nil {
		err = errors.Join(err, types.ErrFailedToGenerateHash)
		return types.RegisterUserOutput{}, err
	}

	userId := uuid.NewString()
	user := model.NewUser(userId, input.Username, password, time.Now().UnixMilli(), userId)

	txCtx, err := s.Repository.BeginTransaction(ctx)
	if err != nil {
		return types.RegisterUserOutput{}, err
	}

	// if transaction is created inside the function, let it commit inside
	if !s.Repository.IsTransaction(ctx) {
		defer func(err error) {
			s.Repository.CommitOrRollbackTransaction(txCtx, err)
		}(err)
	}

	err = s.Repository.InsertUser(txCtx, *user)
	if err != nil {
		return types.RegisterUserOutput{}, err
	}

	// assign default role
	err = s.AssignRole(txCtx, userId, "default")
	if err != nil {
		return types.RegisterUserOutput{}, err
	}

	return types.RegisterUserOutput{
		Id: user.Id,
	}, nil
}
