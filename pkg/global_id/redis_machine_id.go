package global_id

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dgdts/CloudNoteServer/pkg/utils"
	"github.com/redis/go-redis/v9"
)

var _ machineIDGetter = (*redisMachineIDGetter)(nil)

const (
	redisMachineIDKeyPrefix = "global_machine_id_"
	redisMachineIDTTL       = 24 * time.Hour
)

type redisMachineIDGetter struct {
	client  *redis.Client
	localIP string
}

func (r *redisMachineIDGetter) GetMachineID() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	localIP, err := utils.GetLocalIP()
	if err != nil {
		return 0, err
	}

	r.localIP = localIP

	ret, err := r.find(ctx)
	if err == nil {
		return ret, nil
	}

	if !errors.Is(err, redis.Nil) {
		return 0, err
	}

	ret, err = r.register(ctx)
	if err != nil {
		return 0, err
	}

	go r.keep(ret)

	return ret, nil
}

func (r *redisMachineIDGetter) keep(machineID int) {
	key := r.genKey(machineID)
	ticker := time.NewTicker(redisMachineIDTTL / 2)
	defer ticker.Stop()

	var err error
	var ok bool

	for {
		<-ticker.C
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		ok, err = r.client.ExpireAt(ctx, key, time.Now().Add(redisMachineIDTTL)).Result()
		cancel()
		if err != nil {
			log.Printf("keep machine id failed, key: %s, err: %v", key, err)
			continue
		}
		if !ok {
			log.Printf("keep machine id failed, key: %s, err: %v", key, err)
		}
	}
}

func (r *redisMachineIDGetter) genKey(last int) string {
	return fmt.Sprintf("%s%d", redisMachineIDKeyPrefix, last)
}

func (r *redisMachineIDGetter) register(ctx context.Context) (int, error) {
	last, err := lastToInt(r.localIP)
	if err != nil {
		return 0, err
	}

	for {
		ok, err := r.client.SetNX(ctx, r.genKey(last), r.localIP, redisMachineIDTTL).Result()
		if err != nil {
			return 0, err
		}
		if ok {
			return last, nil
		}
	}
}

func (r *redisMachineIDGetter) find(ctx context.Context) (int, error) {
	var cursor uint64
	var keys []string
	var err error
	var val string

	for {
		keys, cursor, err = r.client.Scan(ctx, cursor, redisMachineIDKeyPrefix+"*", 10).Result()
		if err != nil {
			return 0, err
		}

		for _, key := range keys {
			val, err = r.client.Get(ctx, key).Result()
			if err != nil {
				return 0, err
			}

			if val == r.localIP {
				ok, err := r.client.ExpireAt(ctx, key, time.Now().Add(redisMachineIDTTL)).Result()
				if err != nil {
					return 0, err
				}
				if !ok {
					return 0, redis.Nil
				}

				return strconv.Atoi(strings.ReplaceAll(key, redisMachineIDKeyPrefix, ""))
			}
		}
		if cursor == 0 {
			break
		}
	}
	return 0, redis.Nil
}

func lastToInt(ipStr string) (int, error) {
	ipSplit := strings.Split(ipStr, ".")
	if len(ipSplit) != 4 {
		return 0, errors.New("ip format error")
	}

	last := ipSplit[3]

	lastInt, err := strconv.Atoi(last)
	if err != nil {
		return 0, err
	}

	return lastInt, nil
}
