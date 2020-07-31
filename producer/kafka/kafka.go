package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vsivarajah/RiotStatistics/pkg/config"
	"github.com/vsivarajah/RiotStatistics/producer"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaService struct {
	prod *kafka.Producer
}

func New(conf *config.Config) (producer.Sender, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka.BootstrapServers,
	})

	if err != nil {
		return nil, err
	}

	return &kafkaService{prod: p}, nil
}

func (k *kafkaService) Send(ctx context.Context, message interface{}) error {
	// publish sync
	message_2, _ := json.Marshal(message)

	deliveryChan := make(chan kafka.Event)
	topic := "matchDetails"
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

	/*

		message_2, _ := json.Marshal(message)
		msg := &sarama.ProducerMessage{
			Topic: "vignesh",

			Value: sarama.ByteEncoder(message_2),
		}
		p, o, err := k.prod.Produce()
		if err != nil {
			return err
		}

		// publish async
		//producer.Input() <- &sarama.ProducerMessage{

		fmt.Println("Partition: ", p)
		fmt.Println("Offset: ", o)
		return nil
	*/
}
