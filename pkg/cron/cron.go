package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron
var jobs map[string]cron.EntryID
var mutex sync.Mutex

func init() {
	c = cron.New()
	jobs = make(map[string]cron.EntryID)
}

func Start() {
	c.Start()
}

type Worker interface {
	Name() string
	Run()
}

func AddJob(every time.Duration, worker Worker) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := jobs[worker.Name()]
	if ok {
		return fmt.Errorf("job %s already exists", worker.Name())
	}
	s := every.Seconds()
	if s < 1 {
		s = 1
	}
	spec := fmt.Sprintf("@every %.0fs", s)
	job := cron.FuncJob(func() {
		start := time.Now()
		hlog.Infof("job %s started", worker.Name())
		worker.Run()
		hlog.Infof("job %s finished in %.2f seconds", worker.Name(), time.Since(start).Seconds())
	})
	id, err := c.AddJob(spec, cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(job))
	if err != nil {
		return err
	}
	jobs[worker.Name()] = id
	return nil
}

func RemoveJob(name string) error {
	mutex.Lock()
	defer mutex.Unlock()
	id, ok := jobs[name]
	if !ok {
		return fmt.Errorf("job %s not found", name)
	}
	c.Remove(id)
	delete(jobs, name)
	return nil
}
