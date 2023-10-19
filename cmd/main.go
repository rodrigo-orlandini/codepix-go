package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/application/grpc"
	db "github.com/rodrigo-orlandini/codepix-go/infrastructure/database"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGRPCServer(database, 50051)
}
