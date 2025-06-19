package types

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrFailedToGenerateHash = errors.New("failed to save password")
	ErrUserConflicting      = errors.New("user with given username already exist")
)
