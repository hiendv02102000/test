package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

var AutoUID int = 1

// Users struct
type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func (u *Users) CollectionName() string {
	return "Users"
}
