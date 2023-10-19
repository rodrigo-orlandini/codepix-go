package kafka

import (
	"fmt"

	confkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Database        *gorm.DB
	Producer        *confkafka.Producer
	DeliveryChannel chan confkafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *confkafka.Producer, deliveryChannel chan confkafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:        database,
		Producer:        producer,
		DeliveryChannel: deliveryChannel,
	}
}

func (processor *KafkaProcessor) Consume() {
	kafkaConfig := &confkafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}

	consumer, err := confkafka.NewConsumer(kafkaConfig)

	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started.")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.value))
		}
	}
}
