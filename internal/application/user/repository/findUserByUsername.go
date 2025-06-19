package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/radityacandra/go-cms/internal/application/user/model"
	"github.com/radityacandra/go-cms/internal/application/user/types"
)

func (r *Repository) FindUserByUsername(ctx context.Context, input types.FindUserByUsernameInput) (*model.User, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT
			id, username, password, created_at, created_by, updated_at, updated_by
		FROM public.users
		WHERE
			username = $1
			AND is_deleted = false
	`, input.Username)

	var user model.User
	if err := row.StructScan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Join(err, types.ErrUserNotFound)
		}

		return nil, err
	}

	rows, err := r.Db.QueryxContext(ctx, `
		SELECT id, role_id, user_id, created_at, created_by FROM public.user_roles WHERE user_id = $1 AND is_deleted = false
	`, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.UserRole
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		user.UserRoles = append(user.UserRoles, item)
	}

	for i, userRole := range user.UserRoles {
		rows, err := r.Db.QueryxContext(ctx, `
			SELECT id, role_id, access FROM public.role_acls WHERE role_id = $1 AND is_deleted = false
		`, userRole.RoleId)
		if err != nil {
			rows.Close()
			return nil, err
		}

		for rows.Next() {
			var item model.RoleAcl
			if err := rows.StructScan(&item); err != nil {
				rows.Close()
				return nil, err
			}

			user.UserRoles[i].RoleAcls = append(user.UserRoles[i].RoleAcls, item)
		}

		rows.Close()
	}

	return &user, nil
}
