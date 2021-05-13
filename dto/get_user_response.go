package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetUserResponse struct {
	UID      primitive.ObjectID `json:"uid"`
	Username string             `json:"user_name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
