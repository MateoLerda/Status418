package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	GetAll() (*[]models.User, error)
	GetById(id int) (*models.User, error)
	Create(models.User) (*mongo.InsertOneResult, error)
	Update(models.User) (*mongo.UpdateResult, error)
	Delete(id int) (*mongo.DeleteResult, error)
}

type UserRepository struct {
	db DB
}

func NewUserRepository(db DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) GetAll() (*[]models.User, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{}
	data, err := ur.db.GetClient().Database(DBNAME).Collection("Users").Find(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to get users")
		return nil, err
	}
	var users []models.User
	err = data.All(context.TODO(), &users)
	if err != nil {
		err = errors.New("failed to get users")
		return nil, err
	}
	return &users, nil

}

func (ur UserRepository) GetById(userId string) (*models.User, error) {
	DBNAME := os.Getenv("DB_NAME")
	filter := bson.M{"user_id": userId}
	data := ur.db.GetClient().Database(DBNAME).Collection("Users").FindOne(context.TODO(), filter)

	var user models.User
	err := data.Decode(&user)
	if err != nil {
		err = errors.New("failed to get the user")
		return nil, err
	}
	return &user, nil

}

func (ur UserRepository) Create(user *models.User) (*mongo.InsertOneResult, error) {
	DBNAME := os.Getenv("DB_NAME")

	res, err := ur.db.GetClient().Database(DBNAME).Collection("Users").InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return res, nil
}
func (ur UserRepository) Update(user *models.User) (*mongo.UpdateResult, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{"user_id": user.UserId}
	update := bson.M{
		"$set": user,
	}
	res, err := ur.db.GetClient().Database(DBNAME).Collection("Users").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		err = errors.New("failed to update the user")
		return nil, err
	}
	return res, nil
}

func (ur UserRepository) Delete(userId string) (*mongo.DeleteResult, error) {
	DBNAME := os.Getenv("DB_NAME")

	filter := bson.M{
		"user_id": userId,
	}
	res, err := ur.db.GetClient().Database(DBNAME).Collection("Users").DeleteOne(context.TODO(), filter)
	if err != nil {
		err = errors.New("failed to delete the user")
		return nil, err
	}
	return res, nil
}
