package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigDB struct {
	MongoURL string
}

func (config *ConfigDB) InitialDB() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURL))
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("olap").Collection("product-visit")
	return coll
}
