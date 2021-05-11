package entity

var AutoUID int = 1

// Users struct
type Users struct {
	ID       int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Username string `gorm:"column:user_name;"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}
