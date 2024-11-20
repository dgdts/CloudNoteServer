package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

type LocalConfigSource struct {
	path     string
	watcher  *fsnotify.Watcher
	stopChan chan struct{}
	mutex    sync.RWMutex
}

func NewLocalConfigSource() *LocalConfigSource {
	return &LocalConfigSource{
		stopChan: make(chan struct{}),
	}
}

func (l *LocalConfigSource) Load(path string) (*GlobalConfig, error) {
	l.mutex.Lock()
	l.path = path
	l.mutex.Unlock()

	return l.loadFile(path)
}

func (l *LocalConfigSource) loadFile(path string) (*GlobalConfig, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &GlobalConfig{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

func (l *LocalConfigSource) Watch() (<-chan *GlobalConfig, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %w", err)
	}

	l.mutex.Lock()
	l.watcher = watcher
	l.mutex.Unlock()

	dir := filepath.Dir(l.path)
	if err := watcher.Add(dir); err != nil {
		watcher.Close()
		return nil, fmt.Errorf("failed to watch directory: %w", err)
	}

	configChan := make(chan *GlobalConfig)

	go l.watchConfig(configChan)

	return configChan, nil
}

func (l *LocalConfigSource) watchConfig(configChan chan *GlobalConfig) {
	defer close(configChan)

	var debounceTimer *time.Timer
	for {
		select {
		case event, ok := <-l.watcher.Events:
			if !ok {
				return
			}

			l.mutex.RLock()
			targetPath := l.path
			l.mutex.RUnlock()

			if event.Name == targetPath && (event.Op&(fsnotify.Write|fsnotify.Create) != 0) {
				if debounceTimer != nil {
					debounceTimer.Stop()
				}
				debounceTimer = time.AfterFunc(100*time.Millisecond, func() {
					if config, err := l.loadFile(targetPath); err == nil {
						configChan <- config
					}
				})
			}

		case err, ok := <-l.watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("watch error: %v\n", err)

		case <-l.stopChan:
			return
		}
	}
}

func (l *LocalConfigSource) Close() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	close(l.stopChan)

	if l.watcher != nil {
		return l.watcher.Close()
	}

	return nil
}
