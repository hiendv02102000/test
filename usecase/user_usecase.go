package usecase

import (
	"errors"
	"test/dto"
	"test/entity"
	"test/middleware"
	"test/repository"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserUsecase struct {
	Repo repository.UserRepositoryInterface
}

func (u *UserUsecase) GetUserTokenLogin(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.Repo.FirstUser(entity.Users{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}
	if user.Password != req.Password || user.ID == 0 {
		return dto.LoginResponse{}, errors.New("Login fail")
	}
	timeNow := time.Now()

	timeExpriedAt := timeNow.Add(time.Hour * 2)
	// generate uuid
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
		Token:          &tokenString,
		TokenExpriedAt: &timeExpriedAt,
	}
	err = u.Repo.UpdateUser(newUser, user)

	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: tokenString,
	}, nil
}

func (u *UserUsecase) GetProfile(id int) (dto.GetUserResponse, error) {
	user, err := u.Repo.FirstUser(entity.Users{
		ID: id,
	})

	if err != nil {
		return dto.GetUserResponse{}, err
	}
	if user.ID == 0 {
		return dto.GetUserResponse{}, errors.New("Record not found")
	}
	return dto.GetUserResponse{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		Password:    user.Password,
		Address1:    user.Address1,
		Address2:    user.Address2,
		PhoneNumber: user.PhoneNumber,
		Decription:  user.Decription,
	}, nil
}

func (u *UserUsecase) CreateUser(req dto.CreateUserRequest) error {
	user, err := u.Repo.FirstUser(entity.Users{
		Email: req.Email,
	})
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("Email is already exist")
	}
	_, err = u.Repo.CreateUser(entity.Users{
		Email:      req.Email,
		Password:   req.Password,
		Decription: req.Decription,
	})
	return err

}
func (u *UserUsecase) DeleteUser(req dto.DeleteUserRequest) error {
	user, err := u.Repo.FirstUser(entity.Users{
		Email: req.UserEmail,
	})
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("Record not found")
	}
	err = u.Repo.DeleteUser(user)
	return err
}

//
func (u *UserUsecase) PatchUpdateUser(req dto.UserUpdateRequest) error {

	user, err := u.Repo.FirstUser(entity.Users{
		Email: req.Email,
	})

	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("User is not exist")
	}

	if user.Password == req.Password {
		return errors.New("Password is repeated")
	}

	newUser := entity.Users{
		Firstname:  req.Firstname,
		Password:   req.Password,
		Lastname:   req.Lastname,
		Address1:   req.Address1,
		Address2:   req.Address2,
		Decription: req.Decription,
	}
	err = u.Repo.UpdateUser(newUser, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetProfileList() (dto.GetListUserResponse, error) {

	userList, err := u.Repo.FindUserList(entity.Users{})
	if err != nil {
		return dto.GetListUserResponse{}, err
	}
	userProfileList := make([]dto.UserProfile, 0)
	for _, v := range userList {
		us := dto.UserProfile{
			Firstname:   v.Firstname,
			Lastname:    v.Lastname,
			Address1:    v.Address1,
			Address2:    v.Address2,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			Decription:  v.Decription,
		}
		userProfileList = append(userProfileList, us)
	}
	pageList := make([]dto.Page, 0)
	if len(userProfileList) > 5 {

		for i := 0; i < len(userProfileList)-5; i += 5 {
			p := dto.Page{
				NUser:    5,
				UserList: userProfileList[i : i+5],
			}
			pageList = append(pageList, p)

		}
		p := dto.Page{
			NUser:    (len(userProfileList)-1)/5 + 1,
			UserList: userProfileList[(len(userProfileList) - len(userProfileList)%5):],
		}
		pageList = append(pageList, p)
	} else {
		p := dto.Page{
			NUser:    len(userProfileList),
			UserList: userProfileList[0:],
		}
		pageList = append(pageList, p)
	}

	return dto.GetListUserResponse{
		NPage:    len(pageList),
		PageList: pageList,
	}, nil
}

func NewUserUsecase(repo repository.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{Repo: repo}
}
