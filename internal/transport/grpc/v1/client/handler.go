package client

import (
	"context"

	"github.com/martketplace-vkr/user/domain"
	"github.com/martketplace-vkr/user/internal/service/client/dto"
	"github.com/martketplace-vkr/user/pkg/api/grpc/v1/client"
)

type Handler struct {
	service service
	client.UserClientServiceServer
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetUser(ctx context.Context, request *client.GetUserRequest) (resp *client.User, err error) {
	user, err := h.service.GetUser(ctx, request.UserId)
	if err != nil {
		return resp, err
	}

	return user.ToProto(), nil
}

func (h *Handler) UpdateUser(ctx context.Context, request *client.UpdateUserRequest) (resp *client.User, err error) {
	user, err := h.service.UpsertUser(ctx, dto.UpdateUserRequestFromProto(request))
	if err != nil {
		return resp, err
	}

	return user.ToProto(), nil
}

func (h *Handler) AddUserAddress(ctx context.Context, req *client.AddUserAddressRequest) (resp *client.AddUserAddressResponse, err error) {
	address := domain.AddressFromProto(req.Address)

	err = h.service.AddUserAddress(ctx, address)
	if err != nil {
		return resp, err
	}

	return &client.AddUserAddressResponse{
		Address: address.ToProto(),
	}, nil
}

func (h *Handler) GetUserAddresses(ctx context.Context, req *client.GetUserAddressesRequest) (resp *client.GetUserAddressesResponse, err error) {
	addresses, err := h.service.GetUserAddresses(ctx, req.UserID)
	if err != nil {
		return resp, err
	}

	return &client.GetUserAddressesResponse{
		Addresses: addresses.ToProto(),
	}, nil
}
