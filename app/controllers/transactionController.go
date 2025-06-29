package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/services"
	"github.com/sasinduNanayakkara/loyalty-backend/app/utils"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {

	var sessionId = utils.GenerateSessionId()
	var transactionDto dtos.TransactionDto
	if err := c.ShouldBindJSON(&transactionDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	transaction, err := tc.transactionService.CreateTransaction(transactionDto, sessionId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to create new transaction" + ": " + err.Error()})
		return
	}

	log.Printf("%s : New transaction created with ID: %s", sessionId, transaction.ID)

	c.JSON(200, gin.H{"message": "Transaction created successfully", "data": transaction})
}