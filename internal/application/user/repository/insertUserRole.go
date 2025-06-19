package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/model"
)

func (r *Repository) InsertUserRole(ctx context.Context, userRole model.UserRole) error {
	db := r.Db.GetTransaction(ctx)
	_, err := db.NamedExecContext(ctx, `INSERT INTO public.user_roles(id, role_id, user_id, created_by, created_at)
		VALUES(:id, :role_id, :user_id, :created_by, :created_at)`, &userRole)

	return err
}
