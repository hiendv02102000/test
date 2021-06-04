package handler

import (
	"fmt"
	"net/http"
	"test/db"
	"test/dto"
	"test/middleware"
	"test/repository"
	"test/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	usecase usecase.UserUseCaseInterface
}

func NewHTTPHandler(db db.Database) *HTTPHandler {
	usersRepository := repository.NewUserRepository(db)
	usersUsecase := usecase.NewUserUsecase(usersRepository)
	return &HTTPHandler{usecase: usersUsecase}
}

func (h *HTTPHandler) Login(c *gin.Context) {
	req := dto.LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	token, err := h.usecase.GetUserTokenLogin(req)

	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprint(err),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: token,
	}
	c.JSON(http.StatusOK, data)
}

func (h *HTTPHandler) GetUserProfile(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	userID := user.ID
	res, err := h.usecase.GetProfile(userID)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}

func (h *HTTPHandler) UpdateUser(c *gin.Context) {
	req := dto.UserUpdateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.usecase.PatchUpdateUser(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: "Success",
	}
	c.JSON(http.StatusOK, data)
}

func (h *HTTPHandler) CreateNewUser(c *gin.Context) {
	req := dto.CreateUserRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.usecase.CreateUser(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprint(err),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: "Success",
	}
	c.JSON(http.StatusOK, data)
}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	req := dto.DeleteUserRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.usecase.DeleteUser(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprint(err),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: "Success",
	}
	c.JSON(http.StatusOK, data)
}

func (h *HTTPHandler) GetProfileList(c *gin.Context) {
	res, err := h.usecase.GetProfileList()
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
