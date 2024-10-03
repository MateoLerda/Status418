package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInterface interface {
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Create(models.User) error
	Update(models.User) error
	Delete(id int) error
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
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{}
	data, err := ur.db.GetClient().Database(DBNAME).Collection("Users").Find(context.TODO(), filtro)
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

func (ur UserRepository) GetById(id int) (*models.User, error) {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{"user_id": id}
	data := ur.db.GetClient().Database(DBNAME).Collection("Users").FindOne(context.TODO(), filtro)
	if data == nil {
		err := errors.New("Failed to get the user with ID ")
		return nil, err
	}
	var user models.User
	err := data.Decode(&user)
	if err != nil {
		err = errors.New("Failed to get the user with ID ")
		return nil, err
	}
	return &user, nil

}

func (ur UserRepository) Create(user *models.User) (*mongo.InsertOneResult, error) {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")

	result, err := ur.db.GetClient().Database(DBNAME).Collection("Users").InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (ur UserRepository) Update(user *models.User) error {
	
	return nil
}

func (ur UserRepository) Delete(id int) error {
	return nil
}
