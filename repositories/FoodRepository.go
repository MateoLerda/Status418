package repositories

import (
	"Status418/models"
	"context"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

type FoodRepositoryInterface interface {
	GetAll() ([]models.Food, error)
	GetByCode(id int) (models.Food, error)
	Create(models.Food) error
	Update(models.Food) error
	Delete(id int) error
}

type FoodRepository struct {
	db DB
}

func NewFoodRepository(db DB) *FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (fr FoodRepository) GetAll() (*[]models.Food, error) {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{}

	data, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").Find(context.TODO(), filtro)

	var foods []models.Food
	err = data.All(context.TODO(), &foods)
	if err != nil {
		err = errors.New("failed to get foods")
		return nil, err
	}
	return &foods, nil

	//MAPEO DE DATOS
	// var foods []models.Food
	// 	err = data.All(context.TODO(), &foods)

	// 	if err != nil {
	// 		err = errors.New("failed to get foods")
	// 		return nil, err
	// 	}

	// 	var foodDTOs []dto.FoodDTO
	// 	for _, food := range foods {
	// 		foodDTO := dto.FoodDTO{
	// 			Type:            food.Type,
	// 			Moment:          food.Moment,
	// 			Name:            food.Name,
	// 			UnitPrice:       food.UnitPrice,
	// 			CurrentQuantity: food.CurrentQuantity,
	// 			MinimumQuantity: food.MinimumQuantity,
	// 			UserId:          food.UserId,
	// 		}
	// 		foodDTOs = append(foodDTOs, foodDTO)
	// 	}

}

func (fr FoodRepository) GetByCode(id int) (*models.Food, error) {
	envRes := godotenv.Load(".env")
	if envRes != nil {
		return nil, envRes
	}
	DBNAME := os.Getenv("DB_NAME")

	filtro := bson.M{"food_code": id}
	data := fr.db.GetClient().Database(DBNAME).Collection("Foods").FindOne(context.TODO(), filtro)
	var food models.Food

	err := data.Decode(&food)
	if err != nil {
		err = errors.New("failed to get food")
	}
	return &food, err

}

func (fr FoodRepository) Create(food *models.Food) error {
	envRes := godotenv.Load(".env")
	if envRes != nil{
		return envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	_, err:= fr.db.GetClient().Database(DBNAME).Collection("Foods").InsertOne(context.TODO(), food)

	if err != nil {
		err = errors.New("failed to create food")
		return err
	}
	return nil
}

func (fr FoodRepository) Update(food *models.Food) error {
	envRes := godotenv.Load(".env")
	if envRes != nil{
		return envRes
	}
	DBNAME := os.Getenv("DB_NAME")
	filtro := bson.M{"food_code": food.Code}
	_, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").UpdateOne(context.TODO(), filtro, food)
	if err != nil {
		err = errors.New("failed to update food")
		return err
	}
	return nil
}

func (fr FoodRepository) Delete(id int) error {
	envRes := godotenv.Load(".env")
	if envRes != nil{
		return envRes
	}
	DBNAME:= os.Getenv("DB_NAME")
	filtro := bson.M{"food_code": id}
	_, err := fr.db.GetClient().Database(DBNAME).Collection("Foods").DeleteOne(context.TODO(), filtro)
	if err != nil {
		err = errors.New("failed to delete food")
		return err
	}
	return nil
}
