package dto

type DeleteUserRequest struct {
	UserEmail string `json:"email" binding:"email"`
}
