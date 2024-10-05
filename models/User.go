package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId       primitive.ObjectID    `bson:"user_id,omitempty"`
	Name         string       `bson:"first_name"`
	LastName     string       `bson:"last_name"`
	Email        string       `bson:"user_email"`
	Password     string       `bson:"user_password"`
	CreationDate string       `bson:"creation_date"`
	UpdateDate   string       `bson:"update_date"`
}
