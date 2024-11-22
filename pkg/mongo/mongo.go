package mongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOption "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Path        string `yaml:"path"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxPoolSize int    `yaml:"max_pool_size"`
	MinPoolSize int    `yaml:"min_pool_size"`
	Database    string `yaml:"database"`

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
		clientOptions := mongoOption.Client().ApplyURI(mc.getURI())
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

func SelectDB(dbName string) *mongo.Database {
	return getMongoClientManagerInstance().connection.client.Database(dbName)
}

type options struct {
	dbName                 string
	mongoCollectionOptions []*mongoOption.CollectionOptions
}

type CollectionOption func(opt *options)

func WithDBName(dbName string) CollectionOption {
	return func(opt *options) {
		opt.dbName = dbName
	}
}

func WithCollectionOptions(mongoCollectionOptions []*mongoOption.CollectionOptions) CollectionOption {
	return func(opt *options) {
		opt.mongoCollectionOptions = append(opt.mongoCollectionOptions, mongoCollectionOptions...)
	}
}

func RawCollection(collectionName string, opts ...CollectionOption) *mongo.Collection {
	mongoOptions := &options{
		dbName: getMongoClientManagerInstance().connection.Database,
	}

	if len(opts) > 0 {
		for _, o := range opts {
			o(mongoOptions)
		}
	}
	return SelectDB(mongoOptions.dbName).Collection(collectionName, mongoOptions.mongoCollectionOptions...)
}
