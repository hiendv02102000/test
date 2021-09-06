package dto

type DeleteUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}
