package config

var (
	GlobalStaticConfig *GlobalConfig
)

func GetGlobalStaticConfig() *GlobalConfig {
	return GlobalStaticConfig
}

func InitConfigFromLocal(path string) error {
	source := LocalConfigSource{}
	var err error
	GlobalStaticConfig, err = source.Load(path)
	return err
}

func InitConfigFromNacos(path string) error {
	panic("not implemented")
}

func InitConfigFromEtcd(path string) error {
	panic("not implemented")
}

func InitConfigFromConsul(path string) error {
	panic("not implemented")
}
