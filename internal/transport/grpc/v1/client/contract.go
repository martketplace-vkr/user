package client

import (
	"context"

	"github.com/martketplace-vkr/user/domain"
	"github.com/martketplace-vkr/user/internal/service/client/dto"
)

type (
	service interface {
		GetUser(ctx context.Context, req dto.UpdateUserRequest) (user domain.User, err error)
		UpsertUser(ctx context.Context, req dto.UpdateUserRequest) (user domain.User, err error)
		GetUserAddresses(ctx context.Context, userID int64) (addresses domain.AddressList, err error)
		AddUserAddress(ctx context.Context, address *domain.Address) (err error)
	}
)
