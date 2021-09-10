package dto

type GetUserResponse struct {
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
