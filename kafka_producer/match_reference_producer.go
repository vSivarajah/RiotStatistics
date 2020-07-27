package kafka_producer

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func ProduceMatchReferenceDTO() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
}
