package ds

import (
	"context"
	"fmt"
	"loan-back-services/conf"
	"loan-back-services/pkg/logger"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadMongo() (*mongo.Client, error) {
	uri := conf.MongoDSN()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	logger.Sugar.Info("Successfully connect to mongodb.")

	return client, nil
}

func LoadMongo2() (*mongo.Client, error) {
	fUri := os.Getenv("MONGO_URI")
	address := os.Getenv("MONGO_ADDRESS")
	database := os.Getenv("MONGO_DATABASE")
	userName := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	maxPoolSize, _ := strconv.Atoi(os.Getenv("MONGO_MAXPOOLSIZE"))

	uri := "mongodb://sample.host:27017/?maxPoolSize=20&w=majority"
	if fUri != "" {
		uri = fUri
	} else {
		if userName != "" && password != "" {
			uri = fmt.Sprintf("mongodb://%s:%s@%s/%s?maxPoolSize=%d",
				userName, password, address, database, maxPoolSize)
		} else {
			uri = fmt.Sprintf("mongodb://%s/%s/?maxPoolSize=%d",
				address, database, maxPoolSize)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	logger.Sugar.Info("Successfully connected to MongoDB")

	return client, nil
}
