package mongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Path        string `yaml:"path"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxPoolSize int    `yaml:"max_pool_size"`
	MinPoolSize int    `yaml:"min_pool_size"`

	client *mongo.Client
	once   sync.Once
}

type mongoClientManager struct {
	connection *MongoClient
}

var mongoClientManagerInstance *mongoClientManager
var mongoClientManagerInstanceOnce sync.Once

func getMongoClientManagerInstance() *mongoClientManager {
	mongoClientManagerInstanceOnce.Do(func() {
		mongoClientManagerInstance = &mongoClientManager{
			connection: nil,
		}
	})
	return mongoClientManagerInstance
}

func (mc *MongoClient) connect() *mongo.Client {
	mc.once.Do(func() {
		clientOptions := options.Client().ApplyURI(mc.getURI())
		clientOptions.SetMaxPoolSize(uint64(mc.MaxPoolSize))
		clientOptions.SetMinPoolSize(uint64(mc.MinPoolSize))

		var err error
		mc.client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = mc.client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		hlog.Debug("connect mongo success, uri:" + mc.getURI())
	})

	return mc.client
}

func (mc *MongoClient) getURI() string {
	if mc.Username == "" {
		return fmt.Sprintf("mongodb://%s", mc.Path)
	}

	return fmt.Sprintf("mongodb://%s:%s@%s", mc.Username, mc.Password, mc.Path)
}

func RegisterConnection(configs *MongoClient) {
	getMongoClientManagerInstance().connection = configs
	getMongoClientManagerInstance().connection.connect()
}
