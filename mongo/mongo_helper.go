package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
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

type MongoPool struct {
	capacity   int
	host       string
	name       string
	pw         string
	replicaSet string
	pool       chan *mongo.Client
	sync.Mutex
}

func NewMongoPool(config MongoClientConfig) (*MongoPool, error) {
	pool := &MongoPool{
		capacity:   config.MaxPoolSize,
		host:       config.Url,
		pw:         config.PassWord,
		pool:       make(chan *mongo.Client, config.MaxPoolSize),
		replicaSet: config.ReplicaSet,
	}

	for i := 0; i < config.MaxPoolSize; i++ {
		client, err := pool.createClient()
		if err != nil {
			return nil, err
		}
		pool.pool <- client
	}

	return pool, nil
}

func (p *MongoPool) createClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: p.name,
		Password: p.pw,
	}
	clientOptions := options.Client().ApplyURI(p.host).SetAuth(credential)

	if p.replicaSet != "" {
		clientOptions.SetReplicaSet(p.replicaSet)
		clientOptions.SetDirect(false)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (p *MongoPool) GetClient() (*mongo.Client, error) {
	select {
	case client := <-p.pool:
		return client, nil
	default:
		client, err := p.createClient()
		if err != nil {
			return nil, err
		}
		return client, nil
	}
}

func (p *MongoPool) ReleaseClient(client *mongo.Client) {
	p.Lock()
	defer p.Unlock()

	if len(p.pool) >= p.capacity {
		client.Disconnect(context.Background())
		return
	}

	p.pool <- client
}
