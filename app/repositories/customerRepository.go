package repositories

import (
	"database/sql"
	"log"

	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/config"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) CreateNewCustomerRepository(customer models.Customer, sessionId string) error {

	result := config.DB.Create(&customer)
	if result.Error != nil {
		log.Printf("%s : Error creating new customer: %v", sessionId, result.Error)
		return result.Error
	}
	log.Printf("%s : New customer created with ID: %s", sessionId, customer.ID)
	return nil

}

func (r *CustomerRepository) CreateNewLoyaltyCustomer(loyaltyCustomer *models.LoyaltyAccountResponseModel, customerId string, sessionId string) error {
	
	customerLoyalty := models.CustomerLoyalty{
		CustomerID: customerId,
		LoyaltyID: loyaltyCustomer.LoyaltyAccount.ID,		
		Balance: loyaltyCustomer.LoyaltyAccount.Balance,
	}
	result := config.DB.Create(&customerLoyalty)
	if result.Error != nil {
		log.Printf("%s : Error creating new loyalty customer: %v", sessionId, result.Error)
		return result.Error
	}
	log.Printf("%s : New loyalty customer created with ID: %s", sessionId, customerLoyalty.LoyaltyID)
	return nil
}
