package db

import (
	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"

	// import source file
	_ "github.com/golang-migrate/migrate/v4/source/file"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
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
func (db *Database) MigrateDB() error {
	driver, err := mysql.WithInstance(db.DB.DB(), &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+"test/migrations",
		"mysql", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}
func (db *Database) MigrateDBWithGorm() {
	db.DB.AutoMigrate()
}
func (db *Database) First(condition interface{}, value interface{}) error {
	err := db.DB.First(value, condition).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}
func (db *Database) Find(condition interface{}, value interface{}) error {
	err := db.DB.Find(value, condition).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}
func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error

	return err
}
func (db *Database) Delete(value interface{}) error {
	return db.DB.Delete(value).Error
}
func (db *Database) Update(model interface{}, oldVal interface{}, newVal interface{}) error {
	return db.DB.Model(model).Where(oldVal).Updates(newVal).Error
}
