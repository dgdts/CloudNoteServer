package config

type ConfigSource interface {
	Load(path string) (*GlobalConfig, error)
	Watch() (<-chan *GlobalConfig, error)
	Close() error
}
