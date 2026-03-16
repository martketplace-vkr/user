package client

import (
	"context"

	"github.com/martketplace-vkr/user/domain"
	"github.com/martketplace-vkr/user/internal/service/client/dto"
)

type (
	txManager interface {
		Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
	}
	repository interface {
		userRepository
		addressRepository
	}

	userRepository interface {
		UpsertUser(ctx context.Context, req dto.UpdateUserRequest) (user domain.User, err error)
		SelectUser(ctx context.Context, userID int64) (user domain.User, err error)
	}
	addressRepository interface {
		InsertAddress(ctx context.Context, address *domain.Address) (err error)
		SelectUserAddresses(ctx context.Context, userID int64) (addresses []domain.Address, err error)
	}
)
