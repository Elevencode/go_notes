package database

import (
	"context"
	"fmt"
	"go_notes/envs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func InitDatabase() error {
	env := &envs.ServerEnvs

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MONGO_INITDB_ROOT_USERNAME, env.MONGO_INITDB_ROOT_PASSWORD, env.MONGO_INITDB_HOST, env.MONGO_INITDB_PORT)
	log.Println("URI: " + mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	} else {
		MongoClient = mongo
	}

	mongoErr := MongoClient.Ping(ctx, readpref.Primary())
	if mongoErr != nil {
		return mongoErr
	}

	return nil
}
