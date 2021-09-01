package repository

import (
	"test/db"
	"test/entity"
)

type UserRepository struct {
	DB db.Database
}

func (u *UserRepository) FirstUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(condition, &user)
	return user, err
}
func (u *UserRepository) FindUserList(condition entity.Users) (user []entity.Users, err error) {
	err = u.DB.Find(condition, &user)
	return
}
func (u *UserRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	return user, err
}
func (u *UserRepository) DeleteUser(user entity.Users) error {
	err := u.DB.Delete(&user)
	return err
}
func (u *UserRepository) UpdateUser(user, oldUser entity.Users) error {
	return u.DB.Update(entity.Users{}, &oldUser, &user)
}
func NewUserRepository(db db.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
