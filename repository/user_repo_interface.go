package repository

import (
	"test/entity"
)

type UserRepositoryInterface interface {
	FirstUser(condition entity.Users) (entity.Users, error)
	FindUserList(condition entity.Users) (user []entity.Users, err error)
	CreateUser(user entity.Users) (entity.Users, error)
	DeleteUser(user entity.Users) error
	UpdateUser(user, oldUser entity.Users) error
}
