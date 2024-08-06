package config

import "github.com/IBM/sarama"

// Kafka параметры работы и подключения
type Kafka struct {
	Brokers      []string `yaml:"brokers" env:"KAFKA_BROKER" env-default:"localhost:9092"`
	GroupID      string   `yaml:"group_id" env:"KAFKA_GROUP_ID" env-default:""`
	produceRetry int      `yaml:"produce_retry" env:"KAFKA_PRODUCE_RETRY" env-default:"3"`
}

// SaramaConfig генерирует конфиг sarama
func (k Kafka) SaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = k.produceRetry
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	return config
}
