/*
Copyright © 2023 Rodrigo Orlandini <rodrigosorlandini@hotmail.com>
*/
package cmd

import (
	"os"

	confkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rodrigo-orlandini/codepix-go/application/grpc"
	"github.com/rodrigo-orlandini/codepix-go/application/kafka"
	db "github.com/rodrigo-orlandini/codepix-go/infrastructure/database"
	"github.com/spf13/cobra"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all services (gRPC and Kafka Consumer)",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGRPCServer(database, gRPCPortNumber)

		deliveryChannel := make(chan confkafka.Event)
		producer := kafka.NewKafkaProducer()
		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 50051, "gRPC server port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
