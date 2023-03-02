package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, name, password, url, replicaSet string) (*mongo.Client, error) {
	credential := options.Credential{
		Username: name,
		Password: password,
	}
	clientOptions := options.Client().ApplyURI(url).SetAuth(credential)

	if replicaSet != "" {
		clientOptions.SetReplicaSet(replicaSet)
		clientOptions.SetDirect(false)
	}
	return mongo.Connect(ctx, clientOptions)
}
