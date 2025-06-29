package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
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