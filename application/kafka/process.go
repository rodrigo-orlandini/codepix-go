package kafka

import (
	"fmt"
	"os"

	confkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/application/dto"
	"github.com/rodrigo-orlandini/codepix-go/application/factory"
	"github.com/rodrigo-orlandini/codepix-go/application/usecase"
	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
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
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	consumer, err := confkafka.NewConsumer(kafkaConfig)

	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaTransactionsTopic"), os.Getenv("kafkaTransactionConfirmationTopic")}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started.")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			processor.processMessage(msg)
		}
	}
}

func (processor *KafkaProcessor) processMessage(msg *confkafka.Message) {
	transactionTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionTopic:
		processor.processTransaction(msg)
	case transactionConfirmationTopic:
		processor.processTransactionConfirmation(msg)
	default:
		fmt.Println("Invalid topic.", string(msg.Value))
	}
}

func (processor *KafkaProcessor) processTransaction(msg *confkafka.Message) error {
	transaction := dto.NewTransactionDTO()
	err := transaction.ParseJson(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(processor.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)

	if err != nil {
		fmt.Println("Error registering transaction", err)
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code

	transaction.ID = createdTransaction.ID
	transaction.Status = entity.TransactionPending
	transactionJson, err := transaction.ToJSON()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, processor.Producer, processor.DeliveryChannel)
	if err != nil {
		return err
	}

	return nil
}

func (processor *KafkaProcessor) processTransactionConfirmation(msg *confkafka.Message) error {
	transaction := dto.NewTransactionDTO()
	err := transaction.ParseJson(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(processor.Database)

	if transaction.Status == entity.TransactionConfirmed {
		err = processor.confirmTransaction(transaction, transactionUseCase)

		if err != nil {
			return err
		}
	} else if transaction.Status == entity.TransactionCompleted {
		_, err := transactionUseCase.Complete(transaction.ID)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (processor *KafkaProcessor) confirmTransaction(transaction *dto.TransactionDTO, transactionUseCase *usecase.TransactionUseCase) error {
	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)
	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJSON()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, processor.Producer, processor.DeliveryChannel)
	if err != nil {
		return err
	}

	return nil
}
