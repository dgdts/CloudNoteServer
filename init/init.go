package init

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/dgdts/UniversalServer/biz/biz_config"
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/dgdts/UniversalServer/pkg/global_id"
)

func InitServer(config *config.GlobalConfig) *server.Hertz {
	// 1. init logger
	initLogger(config.Log)

	// 2. init server
	s := initServer(config)

	// 3. init middleware
	initMiddleware(s, config)

	// 4. init biz config with nacos
	biz_config.GetBizConfigInstance().Init(config)

	// 5. init redis
	initRedis(config)

	// 6. init global id generator
	global_id.InitWithRedis(uint64(biz_config.GetBizConfigInstance().BusinessID))

	// 7. init mongodb
	initMongo(config)

	// 8. init kafka

	// 9. init cron and start cron job
	// 10. init hertz client
	// 11. init local cache from redis, and start goroutine to sync cache to redis periodically
	return s
}
