package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/segmentio/kafka-go/compress/gzip"
	"github.com/segmentio/kafka-go/compress/lz4"
	"github.com/segmentio/kafka-go/compress/snappy"
	"github.com/segmentio/kafka-go/compress/zstd"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"

	kafkaGo "github.com/segmentio/kafka-go"
)

const (
	SaslTypePlain = "PLAIN"
	SaslTypeScram = "SCRAM"

	SaslScramAlgorithmSha256 = "SCRAM-SHA-256"
	SaslScramAlgorithmSha512 = "SCRAM-SHA-512"
)

var kafkaConfig *KafkaConfig

type KafkaConfig struct {
	Producer map[string]*KafkaProducer
	Consumer map[string]*KafkaConsumer
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

type KafkaProducer struct {
	Address       []string
	Topic         string
	Async         bool
	Ack           int
	CompressCodec string

	SaslConfig *SaslConfig
	TlsConfig  *TlsConfig

	writer *kafkaGo.Writer
}

type KafkaConsumer struct {
	Key       string   `json:"key"`
	Address   []string `json:"address"` // kafka地址
	Group     string   `json:"group"`   // groupId
	Topic     string   `json:"topic"`
	Partition int      `json:"partition"`
	Offset    int64    `json:"offset"`

	SaslConfig *SaslConfig
	TlsConfig  *TlsConfig

	reader *kafkaGo.Reader
}

func RegisterKafkaConfig(config *KafkaConfig) {
	kafkaConfig = config
	initProducer(config.Producer)
	initConsumer(config.Consumer)
}

func initProducer(producerConfigs map[string]*KafkaProducer) {
	for key, producerConfig := range producerConfigs {
		config := kafkaGo.WriterConfig{
			Brokers:      producerConfig.Address,
			Topic:        producerConfig.Topic,
			Async:        producerConfig.Async,
			RequiredAcks: producerConfig.Ack,
			Dialer:       getKafkaDialer(producerConfig.SaslConfig, producerConfig.TlsConfig),
		}

		if producerConfig.CompressCodec != "" {
			switch producerConfig.CompressCodec {
			case "snappy":
				config.CompressionCodec = new(snappy.Codec)
			case "gzip":
				config.CompressionCodec = new(gzip.Codec)
			case "lz4":
				config.CompressionCodec = new(lz4.Codec)
			case "zstd":
				config.CompressionCodec = new(zstd.Codec)
			}
		}
		producerConfigs[key].writer = kafkaGo.NewWriter(config)
	}
}

func getKafkaDialer(saslConfig *SaslConfig, tlsConfig *TlsConfig) *kafkaGo.Dialer {
	if saslConfig == nil || saslConfig.SaslType == "" {
		return nil
	}

	dialer := kafkaGo.DefaultDialer
	if saslConfig.SaslType == SaslTypePlain {
		dialer.SASLMechanism = plain.Mechanism{
			Username: saslConfig.SaslUsername,
			Password: saslConfig.SaslPassword,
		}
	} else if saslConfig.SaslType == SaslTypeScram {
		algorithm := scram.SHA256
		if saslConfig.SaslScramAlgo == SaslScramAlgorithmSha256 {
			algorithm = scram.SHA256
		} else if saslConfig.SaslScramAlgo == SaslScramAlgorithmSha512 {
			algorithm = scram.SHA512
		}

		var err error
		dialer.SASLMechanism, err = scram.Mechanism(algorithm, saslConfig.SaslUsername, saslConfig.SaslPassword)
		if err != nil {
			panic(err)
		}
	}

	// tls config
	if tlsConfig != nil && tlsConfig.TlsEnable {
		caCert, err := os.ReadFile(tlsConfig.CaFile)
		if err != nil {
			panic(err)
		}

		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			panic("failed to parse CA certificate")
		}

		dialer.TLS = &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
		}

		tlsClientConfig := tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
		}

		if tlsConfig.CertFile != "" {
			if tlsConfig.KeyFile == "" {
				tlsConfig.KeyFile = tlsConfig.CertFile
			}
			cert, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
			if err != nil {
				panic(err)
			}
			tlsClientConfig.Certificates = []tls.Certificate{cert}
		}
		dialer.TLS = &tlsClientConfig
	}

	return dialer
}

func initConsumer(consumerConfigs map[string]*KafkaConsumer) {
	for key, consumerConfig := range consumerConfigs {
		config := kafkaGo.ReaderConfig{
			Brokers:     consumerConfig.Address,
			GroupID:     consumerConfig.Group,
			Topic:       consumerConfig.Topic,
			Partition:   consumerConfig.Partition,
			StartOffset: consumerConfig.Offset,
			Dialer:      getKafkaDialer(consumerConfig.SaslConfig, consumerConfig.TlsConfig),
		}
		consumerConfigs[key].reader = kafkaGo.NewReader(config)
	}
}

func GetProducer(key string) *KafkaProducer {
	return kafkaConfig.Producer[key]
}

func GetConsumer(key string) *KafkaConsumer {
	return kafkaConfig.Consumer[key]
}
