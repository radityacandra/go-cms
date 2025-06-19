package repository

import (
	"context"

	"github.com/radityacandra/go-cms/internal/application/user/model"
)

func (r *Repository) InsertUser(ctx context.Context, input model.User) error {
	db := r.Db.GetTransaction(ctx)

	_, err := db.NamedExecContext(ctx, `
		INSERT INTO public.users(id, username, password, created_at, created_by) 
		VALUES (:id, :username, :password, :created_at, :created_by)`,
		&input)

	return err
}
