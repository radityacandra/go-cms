package service

import (
	"context"
	"errors"

	"github.com/radityacandra/go-cms/internal/application/auth/types"
	userTypes "github.com/radityacandra/go-cms/internal/application/user/types"
	"github.com/radityacandra/go-cms/pkg/hash"
	"github.com/radityacandra/go-cms/pkg/jwt"
)

func (s *Service) Login(ctx context.Context, input types.LoginInput) (types.LoginOutput, error) {
	// find user by username
	user, err := s.Repository.FindUserByUsername(ctx, userTypes.FindUserByUsernameInput{
		Username: input.Username,
	})
	if err != nil {
		err = errors.Join(err, types.ErrUserNotFound)
		return types.LoginOutput{}, err
	}

	// validate password
	if ok := hash.MatchHash(input.Password, user.Password); !ok {
		err = types.ErrPasswordMissmatch
		return types.LoginOutput{}, err
	}

	// generate jwt
	token, exp, err := jwt.BuildToken(map[string]interface{}{
		"scopes": user.CollectAllAccess(),
		"sub":    user.Id,
	}, s.PrivateKey)
	if err != nil {
		err = errors.Join(err, types.ErrFailedGenerateToken)
		return types.LoginOutput{}, err
	}

	return types.LoginOutput{
		Token:     token,
		ExpiredAt: exp,
	}, nil
}
