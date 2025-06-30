package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/app/services"
	"github.com/sasinduNanayakkara/loyalty-backend/app/utils"
)

type CustomerController struct {
	customerService *services.CustomerService
}

func NewCustomerController(customerService *services.CustomerService) *CustomerController {
	return &CustomerController{customerService: customerService}
}


func (cc *CustomerController) CreateNewCustomer(c *gin.Context) {

	var sessionId = utils.GenerateSessionId()

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if error := cc.customerService.CreateNewCustomer(customer, sessionId); error != nil {
		c.JSON(400, gin.H{"error": "Failed to create new customer"})
		return
	}

	log.Printf("%s : New customer created with ID: %s", sessionId, customer.ID)

	c.JSON(200, gin.H{"message": "New customer created successfully"})
}

func (cc *CustomerController) CustomerLogin(c *gin.Context) {

	var sessionId = utils.GenerateSessionId()

	var loginDto dtos.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	loginResponse, err := cc.customerService.CustomerLogin(loginDto, sessionId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to login"})
		return
	}

	log.Printf("%s : Customer login successful: %s", sessionId, loginResponse)

	c.JSON(200, gin.H{"message": "Customer login successful", "data": loginResponse})
}

func (cc *CustomerController) GetCustomerLoyaltyAccount(c *gin.Context) {
	var sessionId = utils.GenerateSessionId()

	loyaltyId := c.Param("loyaltyId")
	if loyaltyId == "" {
		c.JSON(400, gin.H{"error": "Loyalty ID is required"})
		return
	}

	loyaltyAccount, err := cc.customerService.GetCustomerLoyaltyAccount(loyaltyId, sessionId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get customer loyalty account"})
		return
	}

	log.Printf("%s : Customer loyalty account retrieved successfully: %v", sessionId, loyaltyAccount)

	c.JSON(200, gin.H{"message": "Customer loyalty account retrieved successfully", "data": loyaltyAccount})
}

func (cc *CustomerController) GetCustomerTransactionHistory(c *gin.Context) {
	var sessionId = utils.GenerateSessionId()

	customerId := c.Param("customerId")
	if customerId == "" {
		c.JSON(400, gin.H{"error": "Customer ID is required"})
		return
	}

	transactionHistory, err := cc.customerService.GetCustomerTransactionHistory(customerId, sessionId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get customer transaction history"})
		return
	}

	log.Printf("%s : Customer transaction history retrieved successfully: %v", sessionId, transactionHistory)

	c.JSON(200, gin.H{"message": "Customer transaction history retrieved successfully", "data": transactionHistory})
}