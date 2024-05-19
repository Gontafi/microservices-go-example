package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"orders/config"
	pb "orders/gen/go/product"
	"orders/internal/api"
	"orders/internal/db"
	"orders/internal/repository"
	"orders/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func RunServer(cfg *config.Config) error {
	log.Println("connecting to db")

	database, err := db.Connect(cfg)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d",
			config.AppConfig.ProductAppHost,
			config.AppConfig.ProductGRPCPort),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}
	defer conn.Close()

	productClient := pb.NewProductServiceClient(conn)

	repo := repository.NewRepository(database)
	services := service.NewService(repo, productClient)
	handlers := api.NewRESTAPI(services)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	log.Println("running REST api server on port", config.AppConfig.ServerPort)

	go func() {
		err := handlers.Router.Listen(fmt.Sprintf(":%d", config.AppConfig.ServerPort))
		if err != nil {
			log.Println("Failed to listen port 8080", err)
		}
	}()

	log.Println("running gRPC server on port", config.AppConfig.GRPCPort)

	<-stop
	log.Println("graceful shutdown")

	err = handlers.Router.Shutdown()
	if err != nil {
		log.Println("Error while trying to shutdown")
		return err
	}

	log.Println("application stop")
	return nil
}
