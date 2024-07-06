package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
