package dto

type CreateUserRequest struct {
	Username string `json:"user_name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
