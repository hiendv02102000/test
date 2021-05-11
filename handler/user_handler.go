package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"test/db"
	"test/dto"
	"test/repository"
	"test/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	usecase usecase.UserUsecase
}

func NewHTTPHandler(db db.Database) *HTTPHandler {
	usersRepository := repository.UserRepository{DB: db}
	usersUsecase := usecase.UserUsecase{Repo: usersRepository}
	return &HTTPHandler{usecase: usersUsecase}
}
func (h *HTTPHandler) GetUserProfile(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.usecase.GetUser(userID)
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
