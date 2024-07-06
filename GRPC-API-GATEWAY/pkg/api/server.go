package server

import (
	"grpc-user-api-gateway/pkg/api/handler"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	r := gin.New()

	r.Use(gin.Logger())

	r.POST("/adduser", userHandler.AddUser)
	r.GET("/user", userHandler.GetUserByID)
	r.GET("/users", userHandler.GetUsersByIDs)
	r.GET("/search", userHandler.SearchUsers)

	return &ServerHTTP{engine: r}
}

func (s *ServerHTTP) Start() {
	log.Printf("Starting Server on 3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
