package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var Mongo MongoInstance

func InitiateMongo() error {

	client, err := GetMongoDbConnection()

	if err != nil {
		return err
	}


	Mongo = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB_NAME")),
	}

	return nil
}

func GetMongoDbConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func (mg *MongoInstance) GetCollection(CollectionName string) *mongo.Collection {
	return mg.DB.Collection(CollectionName)
}

