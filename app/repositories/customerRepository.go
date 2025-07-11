package repositories

import (
	"database/sql"
	"log"

	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/config"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
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

func (r *CustomerRepository) CreateNewLoyaltyCustomer(loyaltyCustomer *dtos.LoyaltyAccountResponseDto, customerId string, sessionId string) error {

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

func (r *CustomerRepository) GetCustomerByEmail(email string, sessionId string) (*models.Customer, error) {
	var customer models.Customer
	result := config.DB.First(&customer, "email = ?", email)
	if result.Error != nil {
		log.Printf("%s : Error fetching customer by email: %v", sessionId, result.Error)
		return nil, result.Error
	}

	if customer.ID == "" {
		log.Printf("%s : No customer found with email: %s", sessionId, email)
		return nil, sql.ErrNoRows
	}

	log.Printf("%s : Customer found - ID: %s, Email: %s, Password length: %d", sessionId, customer.ID, customer.Email, len(customer.Password))
	return &customer, nil
}
