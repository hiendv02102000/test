package repository

import (
	"test/db"
	"test/entity"
)

type UserRepository struct {
	DB db.Database
}

func (u *UserRepository) FindUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.Find(condition, &user)

	return user, err
}
func (u *UserRepository) CreateUser(user entity.Users) (entity.Users, interface{}, error) {
	uID, err := u.DB.Create(&user)
	if err != nil {
		return entity.Users{}, uID, err
	}
	return user, uID, nil
}
