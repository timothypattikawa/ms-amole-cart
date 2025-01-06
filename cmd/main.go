package main

import (
	"log"

	"github.com/timothypattikawa/amole-services/cart-service/api"
	"github.com/timothypattikawa/amole-services/cart-service/api/grpc/client"
	"github.com/timothypattikawa/amole-services/cart-service/internal/config"
	"github.com/timothypattikawa/amole-services/cart-service/internal/handler"
	"github.com/timothypattikawa/amole-services/cart-service/internal/repository"
	"github.com/timothypattikawa/amole-services/cart-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	v := config.LoadViper()
	conf := config.NewConfig(v)

	conn, err := grpc.NewClient("localhost:1200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	prpc := client.NewProductClientgRPC(v, conn)

	dbconn := conf.GetDatabseConfig("postgres").GetDatabaseConnPool()

	repo := repository.NewCartRepository(dbconn)
	service := service.NewCartService(repo, v, dbconn, prpc)
	handler := handler.NewCartHandler(service)

	api.RunServer(handler, *conf.GetServerConfig())
}
