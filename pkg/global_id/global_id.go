package global_id

import (
	"sync"

	"github.com/dgdts/CloudNoteServer/pkg/redis"
)

var glabal_id_generator *idGenerator
var once sync.Once

func InitWithRedis(bizID uint64) {
	once.Do(func() {
		glabal_id_generator = newIDGenerator(bizID, &config{
			machineIDGetter: &redisMachineIDGetter{
				client: redis.GetConnection(),
			},
		})
	})
}

func InitWithLocalMachine(bizID uint64) {
	once.Do(func() {
		glabal_id_generator = newIDGenerator(bizID, &config{
			machineIDGetter: &localMachineIDGetter{},
		})
	})
}

func InitWithCustomMachineID(bizID uint64, id int) {
	once.Do(func() {
		glabal_id_generator = newIDGenerator(bizID, &config{
			machineIDGetter: &customMachineIDGetter{
				id: id,
			},
		})
	})
}

func GenerateUniqueID() string {
	return glabal_id_generator.GenerateUniqueID()
}
