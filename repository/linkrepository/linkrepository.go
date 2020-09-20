package linkrepository

import (
	"context"
	"github.com/eternity-wing/short_link/database"
	"github.com/eternity-wing/short_link/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository struct {
	col *mongo.Collection
}

const collectionName string = "link"

func NewRepository() *Repository {
	return &Repository{
		col: database.GetMongoInstance().GetCollection(collectionName),
	}
}

func (repo *Repository) Find(filter bson.M) *model.Link {
	var link model.Link
	if err := repo.col.FindOne(context.Background(), filter).Decode(&link); err != nil {
		return nil
	}
	return &link
}

func (repo *Repository) Create(link *model.Link) *model.Link {

	if _, err := repo.col.InsertOne(context.Background(), link); err != nil {
		log.Fatal(err)
	}

	return link
}
