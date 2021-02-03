package main

import (
	"github.com/LucianoPierdona/CodePix/application/grpc"
	"github.com/LucianoPierdona/CodePix/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}