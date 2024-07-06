package di

import (
	server "grpc-user-api-gateway/pkg/api"
	"grpc-user-api-gateway/pkg/api/handler"
	"grpc-user-api-gateway/pkg/client"
	"grpc-user-api-gateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	userClient := client.NewUserClient(cfg)
	userHandler := handler.NewUserHandler(userClient)
	serverHTTP := server.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
