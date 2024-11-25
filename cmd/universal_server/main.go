package main

import (
	"context"
	"flag"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	global_init "github.com/dgdts/UniversalServer/init"
	"github.com/dgdts/UniversalServer/pkg/config"
)

func main() {
	// 1. read and parse config
	configFilePath := flag.String("config", "../../conf/dev/conf.yaml", "config file path")
	absPath, err := filepath.Abs(*configFilePath)
	if err != nil {
		panic(err)
	}

	config, err := config.InitConfigFromLocal(absPath)
	if err != nil {
		panic(err)
	}

	// 2. init services according to config
	h := global_init.InitServer(config)

	// 3. add a ping-pong handler for health check
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "pong")
	})

	// 4. start server
	h.Spin()
}
