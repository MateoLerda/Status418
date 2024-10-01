package repositories

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)



func (mongoDB *MongoDB) GetClient() *mongo.Client {
	return mongoDB.MongoClient
}

func NewMongoDB() *MongoDB {
	instancia := &MongoDB{}
	instancia.Connect()

	return instancia
}

func (mongoDB *MongoDB) Connect() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	connectionLink := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(connectionLink)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil {
		 panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	
	return nil
	
}

func (mongoDB *MongoDB) Disconnect()  {
	err := mongoDB.MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
