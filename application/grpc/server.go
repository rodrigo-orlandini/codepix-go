package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGRPCServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, listenError := net.Listen("tcp", address)

	if listenError != nil {
		log.Fatal("It can't start gRPC server.", listenError)
	}

	log.Printf("gRPC server has been started at port %d", port)
	serverError := grpcServer.Serve(listener)

	if serverError != nil {
		log.Fatal("It can't start gRPC server.", listenError)
	}
}
