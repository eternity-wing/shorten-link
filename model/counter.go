package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eternity-wing/short_link/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const counterCollection string = "counter"

type Counter struct {
	ID    string `json:”id,omitempty” bson:”id,omitempty”`
	Value int    `json:”value,omitempty” bson:”value,omitempty”`
}

func GetNextSequenceValue(name string) (int, error) {
	col := getCounterCollection()

	update := bson.M{
		"$inc": bson.M{
			"value": 1,
		},
	}

	var ctr Counter

	if err := col.FindOneAndUpdate(context.Background(), bson.M{"id": name}, update).Decode(&ctr); err != nil {
		log.Fatal(err)
		return -1, err
	}

	return ctr.Value, nil

}

func GetCounter(filter bson.M) (*Counter, error) {
	col := getCounterCollection()

	var ctr Counter
	if err := col.FindOne(context.Background(), filter).Decode(&ctr); err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(ctr.Value, ctr.ID)
	return &ctr, nil
}

func NewCounter(ctr *Counter) (*Counter, error) {
	col := getCounterCollection()

	res, err := col.InsertOne(context.Background(), ctr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println(string(response))

	return ctr, nil
}

func getCounterCollection() *mongo.Collection {
	return database.Mongo.GetCollection(counterCollection)
}
