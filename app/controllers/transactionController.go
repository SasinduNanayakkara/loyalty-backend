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
		log.Printf("%s : Error binding JSON: %v", sessionId, err)
		c.JSON(400, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Printf("%s : JSON binding successful", sessionId)
	log.Printf("%s : CustomerId: '%s' (length: %d)", sessionId, transactionDto.CustomerId, len(transactionDto.CustomerId))
	log.Printf("%s : Amount: %d", sessionId, transactionDto.Amount)
	log.Printf("%s : Description: '%s'", sessionId, transactionDto.Description)

	transaction, err := tc.transactionService.CreateTransaction(transactionDto, sessionId)
	if err != nil {
		log.Printf("%s : Error creating transaction: %v", sessionId, err)
		c.JSON(400, gin.H{"error": "Failed to create new transaction" + ": " + err.Error()})
		return
	}

	log.Printf("%s : New transaction created with ID: %s", sessionId, transaction.ID)

	c.JSON(200, gin.H{"message": "Transaction created successfully", "data": transaction})
}

func (tc *TransactionController) RedeemLoyaltyPoints(c *gin.Context) {
	var sessionId = utils.GenerateSessionId()
	var redeemDto dtos.LoyaltyRewardDto
	if err := c.ShouldBindJSON(&redeemDto); err != nil {
		log.Printf("%s : Error binding JSON: %v", sessionId, err)
		c.JSON(400, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Printf("%s : CustomerId: '%s' (length: %d)", sessionId, redeemDto.CustomerId, len(redeemDto.CustomerId))

	rewardPoints, err := tc.transactionService.RedeemLoyaltyPoints(redeemDto, sessionId)
	if err != nil {
		log.Printf("%s : Error redeeming loyalty points: %v", sessionId, err)
		c.JSON(400, gin.H{"error": "Failed to redeem loyalty points" + ": " + err.Error()})
		return
	}

	log.Printf("%s : Loyalty points redeemed successfully: %d", sessionId, rewardPoints)

	c.JSON(200, gin.H{"message": "Loyalty points redeemed successfully", "data": rewardPoints})
}
