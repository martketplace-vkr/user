package client

import (
	"context"

	"github.com/martketplace-vkr/user/domain"
	"github.com/martketplace-vkr/user/internal/service/client/dto"
)

type service struct {
	txManager  txManager
	repository repository
}

func New(txManager txManager, repository repository) *service {
	return &service{
		txManager:  txManager,
		repository: repository,
	}
}

func (s *service) UpsertUser(ctx context.Context, req dto.UpdateUserRequest) (user domain.User, err error) {
	user, err = s.repository.UpsertUser(ctx, req)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUser(ctx context.Context, userID int64) (user domain.User, err error) {
	return s.repository.SelectUser(ctx, userID)
}

func (s *service) AddUserAddress(ctx context.Context, address *domain.Address) (err error) {
	err = s.repository.InsertAddress(ctx, address)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserAddresses(ctx context.Context, userID int64) (addresses domain.AddressList, err error) {
	return s.repository.SelectUserAddresses(ctx, userID)
}
