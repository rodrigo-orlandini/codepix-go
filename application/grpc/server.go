package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/application/grpc/pb"
	"github.com/rodrigo-orlandini/codepix-go/application/usecase"
	"github.com/rodrigo-orlandini/codepix-go/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixKeyRepository := repository.PixKeyRepositoryDatabase{Database: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixKeyRepository}
	pixGRPCService := NewPixGRPCService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGRPCService)

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
