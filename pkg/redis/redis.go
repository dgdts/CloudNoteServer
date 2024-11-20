package redis

import (
	"context"
	"fmt"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

const defaultClientName = "default"

type RedisClient struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"password"`
	Port        int    `yaml:"port"`
	IdleTimeout int    `yaml:"idle_timeout"`
	DB          int    `yaml:"db"`
	PoolSize    int    `yaml:"pool_size"`

	client *redis.Client
	once   sync.Once
}

type redisClientManager struct {
	connectionMap map[string]*RedisClient
}

var redisClientManagerInstance *redisClientManager
var redisClientManagerInstanceOnce sync.Once

func getRedisClientManagerInstance() *redisClientManager {
	redisClientManagerInstanceOnce.Do(func() {
		redisClientManagerInstance = &redisClientManager{
			connectionMap: make(map[string]*RedisClient),
		}
	})
	return redisClientManagerInstance
}

func (rcm *redisClientManager) updateConfigs(configs map[string]*RedisClient) {
	rcm.connectionMap = configs
}

func (rcm *redisClientManager) getClient(name string) *redis.Client {
	client, ok := rcm.connectionMap[name]
	if !ok {
		panic("cannot get redis name:" + name)
	}
	return client.connect()
}

func (rc *RedisClient) connect() *redis.Client {
	rc.once.Do(func() {
		rc.client = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%v:%d", rc.Host, rc.Port),
			Password:     rc.Password,
			DB:           rc.DB,
			PoolSize:     rc.PoolSize,
			MinIdleConns: rc.IdleTimeout,
			MaxIdleConns: rc.IdleTimeout,
		})

		_, err := rc.client.Ping(context.Background()).Result()
		if err != nil {
			panic("connect redis failed, host:" + rc.Host + ",errmsg:" + err.Error())
		}
	})
	return rc.client
}

func RegisterConnection(configs map[string]*RedisClient) {
	getRedisClientManagerInstance().updateConfigs(configs)
}

func GetConnection(redisName ...string) *redis.Client {
	var clientName string
	if len(redisName) == 0 {
		clientName = defaultClientName
	} else {
		clientName = redisName[0]
	}

	return getRedisClientManagerInstance().getClient(clientName)
}
