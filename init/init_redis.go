package init

import (
	"github.com/dgdts/CloudNoteServer/pkg/config"
	"github.com/dgdts/CloudNoteServer/pkg/redis"
)

func _initRedis(config *config.GlobalConfig) {
	redisConfigMap := make(map[string]*redis.RedisClient)

	for redisName, redisConfig := range config.Redis {
		redisConfigMap[redisName] = &redis.RedisClient{
			Host:        redisConfig.Host,
			Port:        redisConfig.Port,
			Password:    redisConfig.Password,
			PoolSize:    redisConfig.PoolSize,
			IdleTimeout: redisConfig.IdleTimeout,
			DB:          redisConfig.DB,
		}
	}

	redis.RegisterConnection(redisConfigMap)
}
