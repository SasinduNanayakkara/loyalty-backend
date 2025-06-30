package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sasinduNanayakkara/loyalty-backend/app/routes"
	"github.com/sasinduNanayakkara/loyalty-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Loyalty Backend is up and running",
		})
	})
	api := router.Group("/api/v1")
	routes.CustomerRoutes(api, config.DB)
	routes.TransactionRoutes(api, config.DB)

	router.Run()
}
