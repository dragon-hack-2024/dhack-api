package api

import (
	"dhack-api/config"
	"dhack-api/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config, store *db.Store) (*Server, error) {
	gin.SetMode(config.GinMode)
	router := gin.Default()

	server := &Server{
		store:  store,
		config: config,
	}

	// Setup routing for server.
	v1 := router.Group("v1")
	{
		v1.GET("/users/:id", server.GetUserByID)
		v1.GET("/users", server.GetUserList)
		v1.POST("/users", server.CreateUser)
		v1.PUT("/users/:id", server.UpdateUser)
		v1.DELETE("/users/:id", server.DeleteUser)
	}

	// Setup health check routes.
	health := router.Group("health")
	{
		health.GET("/live", server.Live)
		health.GET("/ready", server.Ready)
	}

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
