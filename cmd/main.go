package main

import (
	"Testovoe/internal/configs"
	"Testovoe/internal/grpcserver/gserver"
	"Testovoe/internal/repository"
	"Testovoe/internal/service"
	"Testovoe/internal/tracing"
	pbrate "Testovoe/protos/gen"
	"flag"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "app/.env", "path to config fail")
}
func main() {
	tp, err := tracing.InitTracer("https://tj:14268/api/traces", "Rates Service")
	if err != nil {
		log.Fatal("init tracer", err)
	}
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal("failed logger", zap.Error(err))
	}

	pool, err := configs.Load(configPath, logger)
	if err != nil {
		logger.Fatal("Error config ", zap.Error(err))
	}

	defer pool.Close()

	newRepo := repository.NewRepo(pool, logger, tp)
	newService := service.NewService(newRepo, logger, tp)
	newServer := gserver.NewServer(newService, logger, tp)

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		logger.Fatal("Failed to listen on port ",
			zap.String("port", "50050"),
			zap.Error(err))
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcprometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpcprometheus.StreamServerInterceptor),
	)

	pbrate.RegisterGetRatesServer(grpcServer, newServer)

	logger.Info("Starting gRPC server ...")
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve gRPC ", zap.Error(err))
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2121", nil))
}
