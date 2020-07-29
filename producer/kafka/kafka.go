package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/vsivarajah/RiotStatistics/producer"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type kafkaService struct {
	prod *kafka.Producer
}

func New() producer.Sender {
	return &kafkaService{}
}

func (k *kafkaService) Init(ctx context.Context, cfg interface{}) error {
	// setup sarama log to stdout

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	k.prod = p
	return nil
}

func (k *kafkaService) Send(ctx context.Context, message interface{}) error {
	// publish sync
	message_2, _ := json.Marshal(message)
	deliveryChan := make(chan kafka.Event)
	topic := "vigitorres"

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
