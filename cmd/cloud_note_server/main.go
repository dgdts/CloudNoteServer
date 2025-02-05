package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	global_init "github.com/dgdts/CloudNoteServer/init"
	"github.com/dgdts/CloudNoteServer/pkg/config"
)

func env() {
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "" {
		env = "dev"
	}
	_ = os.Setenv("ENV", env)
}

func main() {
	env()

	workspacePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(workspacePath, "conf", os.Getenv("ENV"), "conf.yaml")

	// 1. read and parse config
	err = config.InitConfigFromLocal(configPath)

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
