package init

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/dgdts/UniversalServer/pkg/config"
)

func InitServer(config *config.GlobalConfig) *server.Hertz {
	// 1. init logger
	initLogger(config.Log)

	// 2. init server
	s := initServer(config)

	return s
	// 3. init middleware
	// 4. init dynamic config with nacos
	// 5. init redis
	// 6. init global id generator
	// 7. init mongodb
	// 8. init kafka
	// 9. init cron and start cron job
	// 10. init hertz client
	// 11. init local cache from redis, and start goroutine to sync cache to redis periodically
}
