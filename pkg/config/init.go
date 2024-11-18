package config

func InitConfigFromLocal(path string) (*GlobalConfig, error) {
	source := LocalConfigSource{}
	return source.Load(path)
}

func InitConfigFromNacos(path string) (*GlobalConfig, error) {
	panic("not implemented")
}

func InitConfigFromEtcd(path string) (*GlobalConfig, error) {
	panic("not implemented")
}

func InitConfigFromConsul(path string) (*GlobalConfig, error) {
	panic("not implemented")
}
