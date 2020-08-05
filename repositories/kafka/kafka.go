package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vsivarajah/RiotStatistics/api"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	repo "github.com/vsivarajah/RiotStatistics/repositories"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaService struct {
	prod *kafka.Producer
}

func New(conf *config.Config) (repo.DbRepository, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka.BootstrapServers,
	})

	if err != nil {
		return nil, err
	}

	return &kafkaService{prod: p}, nil
}

func (k *kafkaService) Send(ctx context.Context, match *api.Match) error {
	// publish sync
	message_2, _ := json.Marshal(match.MatchDTO)

	deliveryChan := make(chan kafka.Event)
	topic := "matchdetails"
	err := k.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message_2),
	}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
	return nil
}

func (k *kafkaService) Get(ctx context.Context, key int) *api.Match {
	return nil
}
