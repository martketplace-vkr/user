package client

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
	"github.com/martketplace-vkr/user/domain"
	"github.com/martketplace-vkr/user/internal/service/client/dto"
)

type repository struct {
	ctxGetter *trmsqlx.CtxGetter
	db        *sqlx.DB
}

func New(db *sqlx.DB, ctxGetter *trmsqlx.CtxGetter) *repository {
	return &repository{
		db:        db,
		ctxGetter: ctxGetter,
	}
}

func (r *repository) UpsertUser(ctx context.Context, req dto.UpdateUserRequest) (user domain.User, err error) {
	query := `
	insert into "user"."user" as u (
		id,
		email,
		first_name,
		last_name,
		avatar_url
	) values (
		$1,
		$2,
		$3,
		$4,
		$5
	)
	on conflict (id) do update
	set
		email = coalesce(excluded.email, u.email),
		first_name = coalesce(excluded.first_name, u.first_name),
		last_name = coalesce(excluded.last_name, u.last_name),
		avatar_url = coalesce(excluded.avatar_url, u.avatar_url)
	returning 
		id,
		email,
		first_name,
		last_name,
		avatar_url,
		created_at

	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&user,
		query,
		req.UserID,
		req.Email,
		req.FirstName,
		req.SecondName,
		req.AvatarUrl,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SelectUser(ctx context.Context, userID int64) (user domain.User, err error) {
	query := `
		select 
			id, 
			email,
			first_name, 
			last_name,
			avatar_url,
			created_at
		from "user"."user"
		where id = $1
	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		&user,
		query,
		userID,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
