package domain

import (
	"time"

	"github.com/martketplace-vkr/user/pkg/api/grpc/v1/client"
)

type Address struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	Country    string    `db:"country"`
	City       string    `db:"city"`
	Street     string    `db:"street"`
	PostalCode string    `db:"postal_code"`
	CreatedAt  time.Time `db:"created_at"`
}

type AddressList []Address

func AddressFromProto(addr *client.Address) *Address {
	return &Address{
		ID:         addr.Id,
		UserID:     addr.UserId,
		Country:    addr.Country,
		City:       addr.City,
		Street:     addr.Street,
		PostalCode: addr.PostalCode,
		// CreatedAt:  addr.CreatedAt,
	}
}

func (a *Address) ToProto() *client.Address {
	return &client.Address{
		Id:         a.ID,
		UserId:     a.UserID,
		Country:    a.Country,
		City:       a.City,
		Street:     a.Street,
		PostalCode: a.PostalCode,
	}
}

func (adrs *AddressList) ToProto() []*client.Address {
	res := make([]*client.Address, 0, len(*adrs))

	for _, address := range *adrs {
		res = append(res, address.ToProto())
	}

	return res
}
