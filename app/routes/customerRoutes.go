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
	transactionRepo := repositories.NewTransactionRepository(db)

	var customerRepoInterface repositories.CustomerRepository = *customerRepo
	var loyaltyAppServiceInterface services.LoyaltyAppServiceInterface = loyaltyAppService
	var transactionRepoInterface repositories.TransactionRepository = *transactionRepo
	customerService := services.NewCustomerService(customerRepoInterface, loyaltyAppServiceInterface, transactionRepoInterface)
	customerController := controllers.NewCustomerController(customerService)

	{
		router.POST("/customer", customerController.CreateNewCustomer)
		router.POST("/login", customerController.CustomerLogin)
		router.GET("/customer/:loyaltyId", customerController.GetCustomerLoyaltyAccount)
		router.GET("/customer/history/:customerId", customerController.GetCustomerTransactionHistory)
	}

}