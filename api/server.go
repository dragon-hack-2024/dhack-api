package api

import (
	"dhack-api/config"
	"dhack-api/db"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

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

		v1.GET("/challenges/:id", server.GetChallengeByID)
		v1.GET("/challenges", server.GetChallengeList)
		v1.POST("/challenges", server.CreateChallenge)
		v1.PUT("/challenges/:id", server.UpdateChallenge)
		v1.DELETE("/challenges/:id", server.DeleteChallenge)

		v1.GET("/stats/:id", server.GetStatByID)
		v1.GET("/stats", server.GetStatList)
		v1.POST("/stats", server.CreateStat)
		v1.PUT("/stats/:id", server.UpdateStat)
		v1.DELETE("/stats/:id", server.DeleteStat)
		v1.GET("/stats/weekly/:id", server.GetWeeklyProgress)
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
