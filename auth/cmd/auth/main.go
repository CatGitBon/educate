package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/vctrl/currency-service/auth/internal/config"
	"github.com/vctrl/currency-service/auth/internal/handler"
	"github.com/vctrl/currency-service/auth/internal/repository"
	"github.com/vctrl/currency-service/auth/internal/service"
	"github.com/vctrl/currency-service/pkg/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	configPath := flag.String("config", "./config", "path to the config file")

	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Временно отключаем базу данных для тестирования
	// db, _, err := db.NewDatabaseConnection(cfg.Database)
	// if err != nil {
	// 	log.Fatalf("error init database connection: %v", err)
	// }

	// repo, err := repository.NewAuth(db)
	// if err != nil {
	// 	log.Fatalf("error init exchange rate repository: %v", err)
	// }

	// srv := service.NewAuth(repo, logger)

	// Создаем mock сервис без базы данных
	srv := service.NewAuth(repository.Auth{}, logger)

	authServer := handler.NewAuthServer(&srv, logger)

	go func() {
		if err := startGRPCServer(cfg, authServer); err != nil {
			log.Fatalf("Error starting GRPC server: %s", err)
		}
	}()

	// Держим сервис запущенным
	select {}
}

func startGRPCServer(cfg config.AppConfig, srv *handler.AuthServer) error {
	lis, err := net.Listen("tcp", ":"+cfg.Service.ServerPort)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, srv)

	log.Printf("gRPC server is listening on :%s", cfg.Service.ServerPort)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
