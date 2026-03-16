package domain

import (
	"time"

	"github.com/martketplace-vkr/user/pkg/api/grpc/v1/client"
)

type User struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	FirstName *string   `db:"first_name"`
	LastName  *string   `db:"last_name"`
	AvatarUrl *string   `db:"avatar_url"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *User) ToProto() *client.User {
	return &client.User{
		Id:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		AvatarUrl: u.AvatarUrl,
	}
}
