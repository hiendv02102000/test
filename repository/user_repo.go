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
func (u *UserRepository) FindUserList(condition entity.Users) (user []entity.Users, err error) {

	err = u.DB.FindAll(condition, &user)
	return
}
func (u *UserRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}
func (u *UserRepository) DeleteUser(user entity.Users) error {
	err := u.DB.Delete(&user)
	return err
}
func (u *UserRepository) UpdateUser(user, oldUser entity.Users) error {
	err := u.DB.Update(entity.Users{}, &oldUser, &user)
	return err
}
