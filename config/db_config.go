package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client
var DEFAULT_DB_NAME = "test_db"

func init() {
	LoadEnvironment()
	InitDatabase()
}

func InitDatabase() (*mongo.Client, error) {
	db_url := GetEnvProperty("database_url")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db_url))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")

	return client, err
}

func GetDatabaseCollection(dbName *string, collectionName string) *mongo.Collection {
	if *dbName == "" || dbName == nil {
		log.Fatal("db name is not provided, proceeding with test_db")
		dbName = &DEFAULT_DB_NAME
	}
	if dbClient == nil {
		var err error
		dbClient, err = InitDatabase()
		if err != nil {
			log.Fatal("client initialization error ", err)
		}
	}
	collection := dbClient.Database(*dbName).Collection(collectionName)

	return collection
}
