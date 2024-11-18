package config

type Global struct {
	NameSpace     string `yaml:"namespace"`
	EnvName       string `yaml:"env_name"`
	LocalIP       string `yaml:"local_ip"`
	ContainerName string `yaml:"container_name"`
}

type Service struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
}

type Hertz struct {
	App             string    `yaml:"app"`
	Server          string    `yaml:"server"`
	BinPath         string    `yaml:"bin_path"`
	ConfPath        string    `yaml:"conf_path"`
	DataPath        string    `yaml:"data_path"`
	EnablePprof     bool      `yaml:"enable_pprof"`
	EnableGzip      bool      `yaml:"enable_gzip"`
	EnableAccessLog bool      `yaml:"enable_access_log"`
	Service         []Service `yaml:"service"`
}

type Log struct {
	LogMode       string `yaml:"log_mode"`
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
	LogCompress   bool   `yaml:"log_compress"`
}

type Registry struct {
	Name            string   `yaml:"name"`
	RegistryAddress []string `yaml:"registry_address"`
	Namespace       string   `yaml:"namespace"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
}

type Selector struct {
	Name       string   `yaml:"name"`
	ServerAddr []string `yaml:"server_addr"`
	Namespace  string   `yaml:"namespace"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
}

type Config struct {
	Name       string   `yaml:"name"`
	ServerAddr []string `yaml:"server_addr"`
	Namespace  string   `yaml:"namespace"`
	Username   string   `yaml:"username"`
	Password   string   `yaml:"password"`
}

type Prometheus struct {
	Enable bool   `yaml:"enable"`
	Addr   string `yaml:"addr"`
	Path   string `yaml:"path"`
}

type SaslConfig struct {
	SaslType      string `yaml:"sasl_type"`
	SaslUsername  string `yaml:"sasl_username"`
	SaslPassword  string `yaml:"sasl_password"`
	SaslScramAlgo string `yaml:"sasl_scram_algo"`
}

type TlsConfig struct {
	TlsEnable          bool   `yaml:"tls_enable"`
	CaFile             string `yaml:"ca_file"`
	CertFile           string `yaml:"cert_file"`
	KeyFile            string `yaml:"key_file"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify"`
}

type ProducerConfig struct {
	Key           string      `yaml:"key"`
	Address       []string    `yaml:"address"`
	Async         bool        `yaml:"async"`
	Topic         string      `yaml:"topic"`
	Ack           int         `yaml:"ack"`
	CompressCodec string      `yaml:"compress_codec"`
	SaslConfig    *SaslConfig `yaml:"sasl_config"`
	TlsConfig     *TlsConfig  `yaml:"tls_config"`
}

type ConsumerConfig struct {
	Key        string      `yaml:"key"`
	Address    []string    `yaml:"address"`
	Group      string      `yaml:"group"`
	Topic      string      `yaml:"topic"`
	Partition  int         `yaml:"partition"`
	Offset     int64       `yaml:"offset"`
	SaslConfig *SaslConfig `yaml:"sasl_config"`
	TlsConfig  *TlsConfig  `yaml:"tls_config"`
}

type Kafka struct {
	Producer []*ProducerConfig `yaml:"producer"`
	Consumer []*ConsumerConfig `yaml:"consumer"`
}

type (
	MySQL struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		Database        string `yaml:"database"`
		MultiStatements int    `yaml:"multi_statements"`
		Charset         string `yaml:"charset"`
		MaxOpenConns    int    `yaml:"max_open_conns"`
		MaxIdleConns    int    `yaml:"max_idle_conns"`
	}

	MySQLConfig struct {
		Masters []*MySQL `yaml:"masters"`
		Slaves  []*MySQL `yaml:"slaves"`
	}
)

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"pool_size"`
	IdleTimeout int    `yaml:"idle_timeout"`
	DB          int    `yaml:"db"`
}

type MongoConfig struct {
	Path        string `yaml:"path"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxPoolSize int    `yaml:"max_pool_size"`
	MinPoolSize int    `yaml:"min_pool_size"`
}

type GlobalConfig struct {
	Global     *Global                 `yaml:"global"`     // Global config
	Hertz      *Hertz                  `yaml:"hertz"`      // Hertz Server config
	Log        *Log                    `yaml:"log"`        // Log config
	Registry   *Registry               `yaml:"registry"`   // Registry center config
	Selector   *Selector               `yaml:"selector"`   // Selector config
	Config     *Config                 `yaml:"config"`     // Config center config
	Prometheus *Prometheus             `yaml:"prometheus"` // Prometheus config
	Kafka      *Kafka                  `yaml:"kafka"`      // Kafka config
	MySQL      map[string]*MySQLConfig `yaml:"mysql"`      // MySQL config
	Redis      map[string]*RedisConfig `yaml:"redis"`      // Redis config
	Mongo      map[string]*MongoConfig `yaml:"mongo"`      // Mongo config
}
