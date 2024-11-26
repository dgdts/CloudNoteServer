package init

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/dgdts/UniversalServer/biz/biz_config"
	"github.com/dgdts/UniversalServer/biz/router"
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/dgdts/UniversalServer/pkg/cron"
	"github.com/dgdts/UniversalServer/pkg/global_id"
	"github.com/dgdts/UniversalServer/pkg/minio"
)

func InitServer(config *config.GlobalConfig) *server.Hertz {
	// 1. init logger
	initLogger(config.Log)

	// 2. init server
	s := initServer(config)

	// 3. init middleware
	initMiddleware(s, config)

	// 4. init biz config with nacos
	// biz_config.GetBizConfigInstance().Init(config)

	// 5. init redis
	// initRedis(config)

	// 6. init global id generator
	global_id.InitWithLocalMachine(uint64(biz_config.GetBizConfigInstance().BusinessID))

	// 7. init mongodb
	initMongo(config)

	// 8. init kafka
	//initAndRunKafkaConsumer(config)

	// 9. init cron and start cron job
	cron.Start()

	// 10. init hertz client
	//initHertzClient(config)

	// 11. init local cache from redis, also start to sync redis from db

	// 12. register router
	router.GeneratedRegister(s)

	// 13. init minio
	minio.Init(config.Minio)

	return s
}
