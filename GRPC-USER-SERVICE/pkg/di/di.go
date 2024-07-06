package di

import (
	server "grpc-user-service/pkg/api"
	"grpc-user-service/pkg/api/service"
	"grpc-user-service/pkg/config"
	"grpc-user-service/pkg/db"
	"grpc-user-service/pkg/repository"
	"grpc-user-service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)

	ServiceServer := service.NewAuthServer(userUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, ServiceServer)
	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil
}
