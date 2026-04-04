package client

import (
	"context"

	"github.com/martketplace-vkr/user/domain"
)

func (r *repository) InsertAddress(ctx context.Context, address *domain.Address) (err error) {
	query := `
		insert into "user".address(
			user_id,
			country,
			city,
			street,
			postal_code
		) values (
			$1,
			$2,
			$3,
			$4,
			$5 
		) returning 
		 	id, 
			user_id,
			country,
			city,
			street,
			postal_code
	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).GetContext(
		ctx,
		address,
		query,
		address.UserID,
		address.Country,
		address.City,
		address.Street,
		address.PostalCode,
	)
	if err != nil {
		return err
	}

	return err
}

func (r *repository) SelectUserAddresses(ctx context.Context, userID int64) (addresses []domain.Address, err error) {
	query := `
		select 
			id,
			user_id,
			country,
			city,
			street,
			postal_code
		from "user".address
		where user_id = $1
	`

	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(
		ctx,
		&addresses,
		query,
		userID,
	)
	if err != nil {
		return addresses, err
	}

	return addresses, nil
}
