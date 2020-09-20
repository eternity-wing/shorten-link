package model

import (
	"context"
	"github.com/eternity-wing/short_link/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Link struct {
	ID int `json:”id,omitempty” bson:”id,omitempty”`
	URL string `json:”url,omitempty” bson:”url,omitempty”`
	UserID int `json:”user_id,omitempty” bson:”user_id,omitempty”`
	ShortLink string
}



const linkCollectionName string = "link"



func GetLink(filter bson.M) (*Link, error) {
	col := getLinkCollection()

	var link Link
	if err := col.FindOne(context.Background(), filter).Decode(&link); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &link ,nil
}

func NewLink(link *Link) (*Link, error) {
	ID, err := GetNextSequenceValue(linkCollectionName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	link.ID = ID

	col := getLinkCollection()
	_, err = col.InsertOne(context.Background(), link)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}


	return link ,nil
}

func getLinkCollection() *mongo.Collection {
	return database.Mongo.GetCollection(linkCollectionName)
}