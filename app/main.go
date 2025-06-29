package app

import (
	"github.com/gin-gonic/gin"
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

	router.Run()
}