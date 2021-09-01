package dto

type CreateUserRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Decription string `json:"decription" binding:"oneof='client' 'admin'"`
}
