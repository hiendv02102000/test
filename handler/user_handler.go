package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/db"
	"test/dto"
	"test/repository"
	"test/usecase"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HTTPHandler struct {
	usecase usecase.UserUsecase
}

func NewHTTPHandler(db db.Database) *HTTPHandler {
	usersRepository := repository.UserRepository{DB: db}
	usersUsecase := usecase.UserUsecase{Repo: usersRepository}
	return &HTTPHandler{usecase: usersUsecase}
}
func (h *HTTPHandler) GetUserProfile(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	userID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(data)
		return
	}
	res, err := h.usecase.GetUser(userID)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
func (h *HTTPHandler) CreateNewUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	req := dto.CreateUserRequest{}
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Result: nil,
			Error:  err.Error(),
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(data)
		return
	}
	res, err := h.usecase.CreateUser(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprint(err),
		}
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
