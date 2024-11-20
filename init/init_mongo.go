package init

import (
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/dgdts/UniversalServer/pkg/mongo"
)

func initMongo(config *config.GlobalConfig) {
	mongoConfig := &mongo.MongoClient{
		Path:     config.Mongo.Path,
		Username: config.Mongo.Username,
		Password: config.Mongo.Password,
	}
	mongo.RegisterConnection(mongoConfig)
}
