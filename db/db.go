package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

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

func (db *Database) Find(condition interface{}, value interface{}) error {
	return db.DB.First(value, condition).Error

}
func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}
func (db *Database) MigrationDB() error {
	driver, err := mysql.WithInstance(db.DB.DB(), &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:test/db/migrations",
		"mysql", driver)
	if err != nil {
		fmt.Println(err.Error() + "12345")
		return err
	}

	// return m.Force(-1)
	err = m.Up()
	if err != nil {
		currentVer, dirty, _ := m.Version()
		if dirty {
			if currentVer == 1 {
				_ = m.Force(-1)
			} else {
				_ = m.Force(int(currentVer) - 1)
			}
		}

		return err
	}

	return nil
}
