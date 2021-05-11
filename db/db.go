package db

import (
	"test/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDB() (Database, error) {
	dsn := "go_test:go_test@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Migrator().AutoMigrate(entity.Users{})

	return Database{
		DB: db,
	}, err
}

func (db *Database) Find(condition interface{}, value interface{}) error {
	return db.DB.First(value, condition).Error
}
func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}
