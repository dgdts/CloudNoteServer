package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	global_init "github.com/dgdts/UniversalServer/init"
	"github.com/dgdts/UniversalServer/pkg/config"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../../conf/dev/conf.yaml"
	}

	// 1. read and parse config
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		panic(err)
	}

	err = config.InitConfigFromLocal(absPath)
	if err != nil {
		panic(err)
	}

	// 2. init services according to config
	h := global_init.InitServer(config.GetGlobalStaticConfig())

	// 3. add a ping-pong handler for health check
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "pong")
	})

	// 4. start server
	h.Spin()
}
