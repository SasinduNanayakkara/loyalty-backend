package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/sasinduNanayakkara/loyalty-backend/app/controllers"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
	"github.com/sasinduNanayakkara/loyalty-backend/app/services"
)


func TransactionRoutes(router *gin.RouterGroup, db *gorm.DB) {

	transactionRepository := repositories.NewTransactionRepository(db)
	loyaltyAppService := services.NewLoyaltyAppService()

	transactionService := services.NewTransactionService(transactionRepository, loyaltyAppService)
	transactionController := controllers.NewTransactionController(*transactionService)

	router.POST("/transaction", transactionController.CreateTransaction)
}