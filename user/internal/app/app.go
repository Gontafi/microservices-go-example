package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user/config"
	"user/internal/api"
	"user/internal/app/grpcapp"
	"user/internal/db"
	"user/internal/repository"
	"user/internal/service"
)

func RunServer(cfg *config.Config) error {
	log.Println("connecting to db")

	database, err := db.Connect(cfg)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(database)
	services := service.NewService(repo)
	handlers := api.NewRESTAPI(services)

	grpcServer := grpcapp.NewServerRPC(config.AppConfig.GRPCPort)

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

	go func() {
		err := grpcServer.Run()
		if err != nil {
			log.Println("Failed to listen port 8081", err)
		}
	}()

	<-stop
	log.Println("graceful shutdown")

	grpcServer.Stop()
	err = handlers.Router.Shutdown()
	if err != nil {
		log.Println("Error while trying to shutdown")
		return err
	}

	log.Println("application stop")
	return nil
}
