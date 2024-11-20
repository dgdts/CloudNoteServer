package biz_config

import (
	"sync"

	"github.com/dgdts/UniversalServer/pkg/config"
)

type BizConfig struct {
	AppID                 string `yaml:"app_id"`
	InitRedisCacheTimeout int    `yaml:"init_redis_cache_timeout"`
	BusinessID            int    `yaml:"business_id"`
	// OSS Platform Config
	// Sop Platform Config
}

var (
	once     sync.Once
	instance *BizConfig
)

func GetBizConfigInstance() *BizConfig {
	once.Do(func() {
		instance = &BizConfig{}
	})
	return instance
}

func (b *BizConfig) Init(config *config.GlobalConfig) {
	b.AppID = "universal_server"
	b.InitRedisCacheTimeout = 10
}
