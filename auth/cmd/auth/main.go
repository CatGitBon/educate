package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/vctrl/currency-service/auth/internal/config"
	"github.com/vctrl/currency-service/auth/internal/db"
	"github.com/vctrl/currency-service/auth/internal/handler"
	"github.com/vctrl/currency-service/auth/internal/repository"
	"github.com/vctrl/currency-service/auth/internal/service"
	"github.com/vctrl/currency-service/pkg/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "currency_requests_total",
			Help: "Total number of requests handled by the currency service",
		},
		[]string{"method"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "currency_request_duration_seconds",
			Help:    "Histogram of response times for requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	appUptime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "currency_service_uptime_seconds",
			Help: "Time since service start in seconds",
		},
	)
)

func init() {
	// Регистрируем метрики
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(appUptime)
}

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

	db, _, err := db.NewDatabaseConnection(cfg.Database)
	if err != nil {
		log.Fatalf("error init database connection: %v", err)
	}

	repo, err := repository.NewAuth(db)
	if err != nil {
		log.Fatalf("error init exchange rate repository: %v", err)
	}

	srv := service.NewAuth(repo, logger)

	authServer := handler.NewAuthServer(srv, logger, requestCount, requestDuration, appUptime)

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
