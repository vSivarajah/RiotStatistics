package kafka_producer

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "senz"
)

var (
	KafkaService kafkaInterface = &kafkaService{}
)

type kafkaService struct{}

type kafkaInterface interface {
	Init() (sarama.SyncProducer, error)
	Send(message string, producer sarama.SyncProducer)
}

func (k *kafkaService) Init() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func (k *kafkaService) Send(message string, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
