package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoClientConfig struct {
	Url             string
	Name            string
	PassWord        string
	ReplicaSet      string
	MaxPoolSize     int
	MinPoolSize     uint64
	MaxConnIdleTime time.Duration
}

func NewMongoClient(config MongoClientConfig) *mongo.Client {

	credential := options.Credential{
		Username: config.Name,
		Password: config.PassWord,
	}
	clientOptions := options.Client().ApplyURI(config.Url).SetAuth(credential)

	if config.ReplicaSet != "" {
		clientOptions.SetReplicaSet(config.ReplicaSet)
		clientOptions.SetDirect(false)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	ctxping, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctxping, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}
