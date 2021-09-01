package dto

type GetUserResponse struct {
	Firstname   string `json:"fisrt_name"`
	Lastname    string `json:"last_name"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"phone_number"`
	Decription  string `json:"decription"`
}
