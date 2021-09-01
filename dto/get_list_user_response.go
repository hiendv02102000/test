package dto

type GetListUserResponse struct {
	NPage    int    `json:"n_page"`
	PageList []Page `json:"page_list"`
}
type Page struct {
	NUser    int           `json:"n_user"`
	UserList []UserProfile `json:"user_list"`
}
type UserProfile struct {
	Firstname   string `json:"fisrt_name"`
	Lastname    string `json:"last_name"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Decription  string `json:"decription"`
}
