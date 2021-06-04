package usecase

import "test/dto"

type UserUseCaseInterface interface {
	GetUserTokenLogin(req dto.LoginRequest) (dto.LoginResponse, error)
	GetProfile(id int) (dto.GetUserResponse, error)
	CreateUser(req dto.CreateUserRequest) error
	DeleteUser(req dto.DeleteUserRequest) error
	PatchUpdateUser(req dto.UserUpdateRequest) error
	GetProfileList() (dto.GetListUserResponse, error)
}
