package init

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	hertz_config "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertz_utils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/dgdts/UniversalServer/pkg/config"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func initServer(config *config.GlobalConfig) *server.Hertz {
	if len(config.Hertz.Service) == 0 {
		hlog.Errorf("service address is empty")
		return nil
	}
	serverOptions := []hertz_config.Option{
		server.WithHostPorts(config.Hertz.Service[0].Address),
	}

	if config.Hertz.EnablePprof {
		serverOptions = append(serverOptions, server.WithHandleMethodNotAllowed(true))
	}

	// // generate hertz server hertz registry center config
	// registryConfig, err := generateHertzRegistryConfig(config)
	// if err != nil {
	// 	hlog.Errorf("generate hertz registry config error: %v", err)
	// 	return nil
	// }

	// serverOptions = append(serverOptions, *registryConfig)

	// // add prometheus config
	// if config.Prometheus != nil && config.Prometheus.Enable {
	// 	serverOptions = append(serverOptions, server.WithTracer(prometheus.NewServerTracer(config.Prometheus.Addr, config.Prometheus.Path)))
	// }

	s := server.New(serverOptions...)
	return s
}

func _generateHertzRegistryConfig(config *config.GlobalConfig) (*hertz_config.Option, error) {
	nacosRegistryServiceConfig := make([]constant.ServerConfig, 0)
	for _, addr := range config.Registry.RegistryAddress {
		ipAndPort := strings.Split(addr, ":")
		uintPort, err := strconv.Atoi(ipAndPort[1])
		if err != nil {
			hlog.Errorf("parse registry address %s error: %v", addr, err)
			continue
		}
		nacosRegistryServiceConfig = append(nacosRegistryServiceConfig, constant.ServerConfig{
			IpAddr: ipAndPort[0],
			Port:   uint64(uintPort),
		})
	}
	if config.Registry.Username == "" {
		config.Registry.Username = os.Getenv(NacosUsernameEnvKey)
	}
	if config.Registry.Password == "" {
		config.Registry.Password = os.Getenv(NacosPasswordEnvKey)
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         config.Registry.Namespace,
		TimeoutMs:           DefaultTimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              filepath.Join(os.TempDir(), "nacos", "log"),
		CacheDir:            filepath.Join(os.TempDir(), "nacos", "cache"),
		LogLevel:            config.Log.LogLevel,
		Username:            config.Registry.Username,
		Password:            config.Registry.Password,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: nacosRegistryServiceConfig,
		},
	)
	if err != nil {
		hlog.Errorf("create nacos client error: %v", err)
		return nil, err
	}

	if len(config.Hertz.Service) == 0 {
		hlog.Errorf("service address is empty")
		return nil, err
	}

	serviceConfig := config.Hertz.Service[0]

	serviceName := fmt.Sprintf("%s.%s.%s.%s",
		FrameName,
		config.Hertz.App,
		config.Hertz.Server,
		serviceConfig.Name)

	ret := server.WithRegistry(nacos.NewNacosRegistry(client), &registry.Info{
		ServiceName: serviceName,
		Addr:        hertz_utils.NewNetAddr("tcp", serviceConfig.Address),
		Weight:      DefaultWeight,
		Tags:        nil,
	})

	return &ret, nil
}
