package global_id

import (
	"strconv"
	"sync"
	"time"
)

const (
	maxCounter = uint64(65535)
)

type idGenerator struct {
	mu         sync.Mutex
	bizID      uint64
	machineID  uint64
	lastSecond uint64
	counter    uint64
}

type config struct {
	machineIDGetter machineIDGetter
}

type machineIDGetter interface {
	GetMachineID() (int, error)
}

func newIDGenerator(bizID uint64, c *config) *idGenerator {
	machineID, err := c.machineIDGetter.GetMachineID()
	if err != nil {
		panic(err)
	}

	return &idGenerator{
		bizID:     bizID,
		machineID: uint64(machineID),
	}
}

func (g *idGenerator) GenerateUniqueID() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	for {
		now := uint64(time.Now().Unix())
		if now != g.lastSecond {
			g.lastSecond = uint64(now)
			g.counter = 0
		} else if g.counter >= maxCounter {
			time.Sleep(time.Millisecond * 10)
			continue
		}
		g.counter++
		ret := g.lastSecond<<31 | g.bizID<<24 | g.machineID<<16 | g.counter
		return strconv.FormatUint(ret, 36)
	}
}
