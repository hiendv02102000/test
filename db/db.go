package db

import (
	"context"
	"test/entity"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DBName string
}

func NewDB() (Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return Database{}, err
	}
	return Database{Client: client,
		DBName: "Test",
	}, nil
}

func (db *Database) Find(condition interface{}, value entity.BaseEntity) error {
	collection := db.Client.Database(db.DBName).Collection(value.CollectionName())
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctx, condition).Decode(value)
	return err
}
func (db *Database) Create(value entity.BaseEntity) (interface{}, error) {
	collection := db.Client.Database(db.DBName).Collection(value.CollectionName())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	res, err := collection.InsertOne(ctx, value)

	return res.InsertedID, err
}
