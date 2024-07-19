package pkg

import (
	"context"
	"fmt"
	"service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client     mongo.Client
	Collection mongo.Collection
}
var ctx = context.Background()

func NewConnection(cfg *config.Config) (*Mongo, error)  {

	url := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDB.Host, cfg.MongoDB.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	mycol := client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)
	return &Mongo{Client: *client, Collection: *mycol}, nil
}