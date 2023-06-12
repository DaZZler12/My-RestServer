package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Brand     string    `json:"brand" bson:"brand"`
	Model     string    `json:"model" bson:"model"`
	Item_Name string    `json:"item_name" bson:"item_name"`
	Year      int64     `json:"year" bson:"year"`
	Price     float64   `json:"price" bson:"price"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
