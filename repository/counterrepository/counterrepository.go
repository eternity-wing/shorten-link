package counterrepository

import (
	"context"
	"github.com/eternity-wing/short_link/database"
	"github.com/eternity-wing/short_link/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName string = "counter"

type Repository struct {
	col *mongo.Collection
}

func NewRepository() *Repository {
	return &Repository{
		col: database.GetMongoInstance().GetCollection(collectionName),
	}
}

func (repo *Repository) GetNextSequenceValue(name string) int {

	update := bson.M{
		"$inc": bson.M{
			"value": 1,
		},
	}

	var ctr model.Counter

	if err := repo.col.FindOneAndUpdate(context.Background(), bson.M{"id": name}, update).Decode(&ctr); err != nil {
		log.Fatal(err)
	}

	return ctr.Value

}

func (repo *Repository) Find(filter bson.M) *model.Counter {

	var ctr model.Counter
	if err := repo.col.FindOne(context.Background(), filter).Decode(&ctr); err != nil {
		return nil
	}
	return &ctr
}

func (repo *Repository) Create(ctr *model.Counter) *model.Counter {
	_, err := repo.col.InsertOne(context.Background(), ctr)
	if err != nil {
		log.Fatal(err)
	}

	return ctr
}
