package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/model"
)

func (r *Repository) FindRoleByName(ctx context.Context, roleName string) (*model.Role, error) {
	row := r.Db.QueryRowxContext(ctx, `SELECT id, name FROM public.roles WHERE name = $1 AND is_deleted = false`, roleName)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var role model.Role
	if err := row.StructScan(&role); err != nil {
		return nil, err
	}

	return &role, nil
}
