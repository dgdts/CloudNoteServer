package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	global_init "github.com/dgdts/CloudNoteServer/init"
	"github.com/dgdts/CloudNoteServer/pkg/config"
	"github.com/dgdts/CloudNoteServer/pkg/utils"
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

	binPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	rootPath := filepath.Dir(binPath)
	configPath := fmt.Sprintf("%s/configs/%s/conf.yaml", rootPath, os.Getenv("ENV"))
	if utils.IsDevEnv() {
		configPath = fmt.Sprintf("./configs/%s/conf.yaml", utils.GetEnv()) //本地临时调式，如果真正run服务，用上面一行的代码
	}

	// 1. read and parse config
	err := config.InitConfigFromLocal(configPath)

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
