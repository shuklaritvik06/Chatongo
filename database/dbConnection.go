package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func GetDB() *mongo.Client {
	return db
}

func DBConnect() {
	uri := os.Getenv("DB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	db = client
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MongoDB!")
}
