package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/vsivarajah/RiotStatistics/producer"
	"log"
	"os"
)

type kafkaService struct {
	prod sarama.SyncProducer
}

func New() producer.Sender {
	return &kafkaService{}
}

func (k *kafkaService) Init(ctx context.Context, cfg interface{}) error {
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
	prd, err := sarama.NewSyncProducer([]string{""}, config)
	if err != nil {
		return err
	}

	k.prod = prd
	return nil
}

func (k *kafkaService) Send(ctx context.Context, message interface{}) error {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: "",
		Value: sarama.StringEncoder(message.(string)),
	}

	p, o, err := k.prod.SendMessage(msg)
	if err != nil {
		return err
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
	return nil
}
