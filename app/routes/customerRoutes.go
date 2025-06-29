package routes

import (

	"github.com/gin-gonic/gin"
	"github.com/sasinduNanayakkara/loyalty-backend/app/controllers"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
	"github.com/sasinduNanayakkara/loyalty-backend/app/services"
	"gorm.io/gorm"
)

func CustomerRoutes(router *gin.RouterGroup, db *gorm.DB) {

	customerRepo := repositories.NewCustomerRepository(db)
	loyaltyAppService := services.NewLoyaltyAppService()

	var customerRepoInterface repositories.CustomerRepository = *customerRepo
	customerService := services.NewCustomerService(customerRepoInterface, loyaltyAppService)
	customerController := controllers.NewCustomerController(customerService)

	{
		router.POST("/customer", customerController.CreateNewCustomer)
		router.POST("/login", customerController.CustomerLogin)
	}

}