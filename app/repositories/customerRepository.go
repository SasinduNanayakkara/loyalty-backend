package repositories

import (
	"database/sql"
	"log"

	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) CreateNewCustomerRepository(customer models.Customer, sessionId string) error {

	query := "INSERT INTO CUSTOMER (ID, NAME, EMAIL, PHONE) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, customer.ID, customer.Name, customer.Email, customer.PhoneNumber)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	log.Printf("%s : New customer created with ID: %d", sessionId, userId)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) CreateNewLoyaltyCustomer(loyaltyCustomer *models.LoyaltyAccountResponseModel, customerId string, sessionId string) error { 
	query := "INSERT INTO CUSTOMER_LOYALTY (ID, LOYALTY_CUSTOMER_ID, BALANCE) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, customerId, loyaltyCustomer.LoyaltyAccount.ID, loyaltyCustomer.LoyaltyAccount.Balance)
	if err != nil {
		return err
	}
	return nil
}