package usecase

import (
	"errors"
	"test/dto"
	"test/entity"
	"test/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	Repo repository.UserRepository
}

func (u *UserUsecase) GetUser(id primitive.ObjectID) (dto.GetUserResponse, error) {

	user, _ := u.Repo.FindUser(entity.Users{
		ID: id,
	})

	return dto.GetUserResponse{
		UID:      user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
func (u *UserUsecase) CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user, err := u.Repo.FindUser(entity.Users{
		Email: req.Email,
	})

	if user.ID.IsZero() {
		uID, err := u.Repo.DB.Create(&entity.Users{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		})
		return dto.CreateUserResponse{ID: uID}, err
	} else {
		if err == nil {
			return dto.CreateUserResponse{}, errors.New("Email is already exist")
		}

	}

	return dto.CreateUserResponse{}, err

}
