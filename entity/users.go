package entity

import "time"

// Users struct
type Users struct {
	ID             int        `gorm:"column:id;primary_key;auto_increment;not null"`
	Firstname      string     `gorm:"column:first_name;"`
	Lastname       string     `gorm:"column:last_name"`
	Address1       string     `gorm:"column:address1"`
	Address2       string     `gorm:"column:address2"`
	Email          string     `gorm:"column:email;not null"`
	Password       string     `gorm:"column:password;not null"`
	PhoneNumber    int        `gorm:"column:phone_number"`
	Decription     string     `gorm:"column:decription"`
	Token          *string    `gorm:"column:token"`
	TokenExpriedAt *time.Time `gorm:"column:token_expried_at"`
}
