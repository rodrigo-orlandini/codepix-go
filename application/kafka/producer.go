package kafka

import (
	"fmt"
	"os"

	confkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer() *confkafka.Producer {
	kafkaConfig := &confkafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	producer, err := confkafka.NewProducer(kafkaConfig)

	if err != nil {
		panic(err)
	}

	return producer
}

func Publish(msg string, topic string, producer *confkafka.Producer, deliveryChannel chan confkafka.Event) error {
	message := &confkafka.Message{
		TopicPartition: confkafka.TopicPartition{Topic: &topic, Partition: confkafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChannel chan confkafka.Event) {
	for e := range deliveryChannel {
		switch event := e.(type) {
		case *confkafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", event.TopicPartition)
			} else {
				fmt.Println("Delivery message to:", event.TopicPartition)
			}
		}
	}
}
