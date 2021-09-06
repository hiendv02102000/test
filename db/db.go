package db

import (
	"test/entity"

	"github.com/jinzhu/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDB() (Database, error) {
	dsn := "go_test:go_test@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	return Database{
		DB: db,
	}, err

}
func (db *Database) MigrationDB() error {
	return db.DB.AutoMigrate(entity.Users{}).Error
}
func (db *Database) First(condition interface{}, value interface{}) error {
	return db.DB.First(value, condition).Error
}
func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}
func (db *Database) Delete(value interface{}) error {
	return db.DB.Delete(value).Error
}
func (db *Database) Update(oldVal, newVal, model interface{}) error {
	return db.DB.Model(model).Where(oldVal).Update(newVal).Error
}
