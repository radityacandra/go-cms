package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/user/model"
)

func (s *Service) AssignRole(ctx context.Context, userId, roleName string) error {
	role, err := s.Repository.FindRoleByName(ctx, roleName)
	if err != nil {
		return err
	}

	txCtx, err := s.Repository.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	// if transaction is created inside the function, let it commit inside
	if !s.Repository.IsTransaction(ctx) {
		defer func() {
			s.Repository.CommitOrRollbackTransaction(txCtx, err)
		}()
	}

	userRole := model.UserRole{
		Id:        uuid.NewString(),
		RoleId:    role.Id,
		UserId:    userId,
		CreatedAt: time.Now().UnixMilli(),
		CreatedBy: userId,
	}
	err = s.Repository.InsertUserRole(txCtx, userRole)

	return err
}
