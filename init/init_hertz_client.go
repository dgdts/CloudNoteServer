package init

import (
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	DefaultHertzClientTimeoutMs = 5000
)

var HertzClient *client.Client

func _initHertzClient(config *config.GlobalConfig) {
	client, err := client.NewClient()
	if err != nil {
		panic(err)
	}

	serverConfigs := make([]constant.ServerConfig, 0)
	for _, addr := range config.Selector.ServerAddr {
		ipAndPort := strings.Split(addr, ":")
		uintPort, err := strconv.Atoi(ipAndPort[1])
		if err != nil {
			panic(err)
		}
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: ipAndPort[0],
			Port:   uint64(uintPort),
		})
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Selector.Namespace,
		TimeoutMs:           DefaultHertzClientTimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            config.Log.LogLevel,
	}

	nacosClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	resolver := nacos.NewNacosResolver(nacosClient)
	client.Use(sd.Discovery(resolver))

	HertzClient = client
}
