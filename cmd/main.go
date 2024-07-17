package main

import (
	"Testovoe/internal/configs"
	"Testovoe/internal/grpcserver/gserver"
	"Testovoe/internal/repository"
	"Testovoe/internal/service"
	pbrate "Testovoe/protos/gen"
	"flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "app/.env", "path to config fail")
}
func main() {
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
	defer logger.Sync()

	newRepo := repository.NewRepo(pool, logger)
	newService := service.NewService(newRepo, logger)
	newServer := gserver.NewServer(newService, logger)

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		logger.Fatal("Failed to listen on port ",
			zap.String("port", "50050"),
			zap.Error(err))
	}
	grpcServer := grpc.NewServer()

	pbrate.RegisterGetRatesServer(grpcServer, newServer)

	logger.Info("Starting gRPC server ...")
	if err = grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve gRPC ", zap.Error(err))
	}
}
