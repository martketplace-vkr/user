package dto

import "github.com/martketplace-vkr/user/pkg/api/grpc/v1/client"

type UpdateUserRequest struct {
	UserID     int64
	Email      string
	FirstName  *string
	SecondName *string
	AvatarUrl  *string
}

func UpdateUserRequestFromProto(req *client.UpdateUserRequest) UpdateUserRequest {
	return UpdateUserRequest{
		UserID:     req.UserId,
		Email:      req.Email,
		FirstName:  req.FirstName,
		SecondName: req.LastName,
		AvatarUrl:  req.AvatarUrl,
	}
}
