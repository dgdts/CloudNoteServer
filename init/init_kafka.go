package init

import (
	"github.com/dgdts/CloudNoteServer/pkg/config"
	"github.com/dgdts/CloudNoteServer/pkg/kafka"
)

func _initAndRunKafkaConsumer(config *config.GlobalConfig) {
	producerConfigsMap := make(map[string]*kafka.KafkaProducer)

	for _, producerConfig := range config.Kafka.Producer {
		producerConfigsMap[producerConfig.Key] = &kafka.KafkaProducer{
			Address:       producerConfig.Address,
			Topic:         producerConfig.Topic,
			Async:         producerConfig.Async,
			Ack:           producerConfig.Ack,
			CompressCodec: producerConfig.CompressCodec,
			SaslConfig:    producerConfig.SaslConfig,
			TlsConfig:     producerConfig.TlsConfig,
		}
	}

	consumerConfigsMap := make(map[string]*kafka.KafkaConsumer)
	for _, consumerConfig := range config.Kafka.Consumer {
		consumerConfigsMap[consumerConfig.Key] = &kafka.KafkaConsumer{
			Key:        consumerConfig.Key,
			Address:    consumerConfig.Address,
			Group:      consumerConfig.Group,
			Topic:      consumerConfig.Topic,
			Partition:  consumerConfig.Partition,
			Offset:     consumerConfig.Offset,
			SaslConfig: consumerConfig.SaslConfig,
			TlsConfig:  consumerConfig.TlsConfig,
		}
	}

	kafkaConfig := &kafka.KafkaConfig{
		Producer: producerConfigsMap,
		Consumer: consumerConfigsMap,
	}

	kafka.RegisterKafkaConfig(kafkaConfig)
}
