package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sasinduNanayakkara/loyalty-backend/app/controllers"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
	"github.com/sasinduNanayakkara/loyalty-backend/app/services"
)

func CustomerRoutes(router *gin.RouterGroup, db *sql.DB) {

	customerRepo := repositories.NewCustomerRepository(db)
	loyaltyAppService := services.NewLoyaltyAppService()

	var customerRepoInterface repositories.CustomerRepository = *customerRepo
	customerService := services.NewCustomerService(customerRepoInterface, loyaltyAppService)
	customerController := controllers.NewCustomerController(customerService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/customer", customerController.CreateNewCustomer)
	}

}