package init

import (
	"github.com/dgdts/CloudNoteServer/pkg/config"
	"github.com/dgdts/CloudNoteServer/pkg/mongo"
)

func initMongo(config *config.GlobalConfig) {
	mongoConfig := &mongo.MongoClient{
		Path:        config.Mongo.Path,
		Username:    config.Mongo.Username,
		Password:    config.Mongo.Password,
		MaxPoolSize: config.Mongo.MaxPoolSize,
		MinPoolSize: config.Mongo.MinPoolSize,
		Database:    config.Mongo.Database,
	}
	mongo.RegisterConnection(mongoConfig)
}
