package types

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordMissmatch   = errors.New("username or password is incorrect")
	ErrFailedGenerateToken = errors.New("failed to generate token")
)
