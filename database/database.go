package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg *MongoInstance

var lock = &sync.Mutex{}

func GetMongoInstance() *MongoInstance {
	if mg == nil {
		lock.Lock()
		defer lock.Unlock()
		InitiateMongo()
	}

	if mg == nil {
		InitiateMongo()
	}
	return mg
}

func InitiateMongo() {
	client := GetMongoDbConnection()
	mg = &MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB_NAME")),
	}
}

func GetMongoDbConnection() *mongo.Client {
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
	return client
}

func (mg *MongoInstance) GetCollection(CollectionName string) *mongo.Collection {
	return mg.DB.Collection(CollectionName)
}
