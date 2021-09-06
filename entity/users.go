package entity

import "time"

var AutoUID int = 1

// Users struct
type Users struct {
	ID             int        `gorm:"column:id;primary_key;auto_increment"`
	Username       string     `gorm:"column:user_name;"`
	Email          string     `gorm:"column:email;not null"`
	Password       string     `gorm:"column:password;not null"`
	Token          string     `gorm:"column:token"`
	TokenExpriedAt *time.Time `gorm:"column:token_expried_at"`
}
