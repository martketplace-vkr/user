package dto

type UpdateUserRequest struct {
	UserID     int64
	FirstName  *string
	SecondName *string
	AvatarUrl  *string
}
