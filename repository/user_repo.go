package repository

import (
	"test/db"
	"test/entity"
)

type UserRepositoryInterface interface {
	FirstUser(condition entity.Users) (entity.Users, error)
	CreateUser(user entity.Users) (entity.Users, error)
	UpdateUser(oldUser, newUser entity.Users) error
	DeleteUser(user entity.Users) error
}

type userRepository struct {
	DB db.Database
}

func (u *userRepository) FirstUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(condition, &user)
	return user, err
}
func (u *userRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	return user, err
}
func (u *userRepository) UpdateUser(oldUser, newUser entity.Users) error {
	return u.DB.Update(oldUser, newUser, entity.Users{})
}
func (u *userRepository) DeleteUser(user entity.Users) error {
	return u.DB.Delete(user)
}
func NewUserRepository(db db.Database) *userRepository {
	return &userRepository{DB: db}
}
