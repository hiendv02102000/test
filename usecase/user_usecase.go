package usecase

import (
	"errors"
	"test/dto"
	"test/entity"
	"test/middleware"
	"test/repository"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserUsecaseInterface interface {
	GetUser(id int) (dto.GetUserResponse, error)
	CreateUser(req dto.RegisterUserRequest) error
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
}
type userUsecase struct {
	repo repository.UserRepositoryInterface
}

func (u *userUsecase) GetUser(id int) (dto.GetUserResponse, error) {
	user, err := u.repo.FirstUser(entity.Users{
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
func (u *userUsecase) CreateUser(req dto.RegisterUserRequest) error {
	user, err := u.repo.FirstUser(entity.Users{
		Email: req.Email,
	})
	if err.Error() == "record not found" {
		_, err = u.repo.CreateUser(entity.Users{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		})
		return err
	}
	if user.ID != 0 {
		return errors.New("Email is already exist")
	}
	return err
}
func (u *userUsecase) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.repo.FirstUser(entity.Users{
		Email:    req.Email,
		Password: req.Password,
	})
	if gorm.IsRecordNotFoundError(err) {
		return dto.LoginResponse{}, errors.New("email or password is invalid")
	}
	if err != nil {
		return dto.LoginResponse{}, err
	}

	timeNow := time.Now()
	timeExpriedAt := timeNow.Add(time.Hour * 168)
	uuid := uuid.Must(uuid.NewV4(), nil)
	tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
		UUID:       uuid,
		Authorized: true,
		ExpriedAt:  timeExpriedAt,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}
	newUser := entity.Users{
		Token:          tokenString,
		TokenExpriedAt: &timeExpriedAt,
	}
	err = u.repo.UpdateUser(user, newUser)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{Token: tokenString}, nil
}
func NewUserUsecase(r repository.UserRepositoryInterface) *userUsecase {
	return &userUsecase{repo: r}
}