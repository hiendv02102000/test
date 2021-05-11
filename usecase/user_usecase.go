package usecase

import (
	"errors"
	"test/dto"
	"test/entity"
	"test/repository"
)

type UserUsecase struct {
	Repo repository.UserRepository
}

func (u *UserUsecase) GetUser(id int) (dto.GetUserResponse, error) {

	user, err := u.Repo.FindUser(entity.Users{
		ID: id,
	})

	if err != nil {
		return dto.GetUserResponse{}, err
	}
	return dto.GetUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
func (u *UserUsecase) CreateUser(req dto.CreateUserRequest) error {
	user, err := u.Repo.FindUser(entity.Users{
		Email: req.Email,
	})

	if err.Error() == "record not found" {
		err = u.Repo.DB.Create(&entity.Users{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		})
		return err
	}

	if user.ID != 0 {
		return errors.New("Email is already exist")
	}
	return nil

}
