package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoClientConfig struct {
	Url             string
	Name            string
	PassWord        string
	ReplicaSet      string
	MaxPoolSize     uint64
	MinPoolSize     uint64
	MaxConnIdleTime time.Duration
}

func NewMongoClient(ctx context.Context, config MongoClientConfig) (*mongo.Client, error) {
	credential := options.Credential{
		Username: config.Name,
		Password: config.PassWord,
	}
	clientOptions := options.Client().ApplyURI(config.Url).SetAuth(credential).SetMaxConnIdleTime(config.MaxConnIdleTime)

	if config.MaxPoolSize > 0 {
		clientOptions.SetMaxPoolSize(config.MaxPoolSize)
	}

	if config.MinPoolSize > 0 {
		clientOptions.SetMinPoolSize(config.MinPoolSize)

	}

	if config.ReplicaSet != "" {
		clientOptions.SetReplicaSet(config.ReplicaSet)
		clientOptions.SetDirect(false)
	}
	return mongo.Connect(ctx, clientOptions)
}
