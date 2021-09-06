package dto

type RegisterUserRequest struct {
	Username string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required email"`
	Password string `json:"password" binding:"required min=8"`
}
