package models

type User struct {
	UserId int `bson:"user_id,omitempty"`
	Name string `bson:"first_name"`
	LastName string `bson:"last_name"`
	Email string `bson:"user_email"`
	Password string `bson:"user_password"`
}