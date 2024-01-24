package mongotool

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoManager 管理 MongoDB 连接
type MongoManager struct {
	client   *mongo.Client
	config   MongoConfig
	database *mongo.Database
}

// 单例实例和初始化锁
var (
	instance *MongoManager
	once     sync.Once
)

// MongoConfig 配置 MongoDB 连接
type MongoConfig struct {
	URI            string
	Database       string
	ConnectTimeout time.Duration
}

// NewMongoManager 创建或返回一个已存在的 MongoDB 连接实例
func NewMongoManager(config MongoConfig) (*MongoManager, error) {
	var err error
	once.Do(func() {
		var client *mongo.Client
		client, err = mongo.NewClient(options.Client().ApplyURI(config.URI))
		if err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			return
		}

		// 测试连接
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			return
		}

		instance = &MongoManager{
			client:   client,
			config:   config,
			database: client.Database(config.Database),
		}
	})
	return instance, err
}

// GetDatabase 返回 MongoDB 数据库实例
func (m *MongoManager) GetDatabase() *mongo.Database {
	return m.database
}

// Close 断开 MongoDB 连接
func (m *MongoManager) Close() {
	if m.client != nil {
		_ = m.client.Disconnect(context.Background())
	}
}
