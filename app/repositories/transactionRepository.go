package repositories

import (
	"log"

	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transactionModel models.Transaction, sessionId string) (string, error) {
	if err := r.db.Create(&transactionModel).Error; err != nil {
		log.Printf("%s : Error creating transaction: %v", sessionId, err)
		return "", err
	}
	log.Printf("%s : Transaction created with ID: %s", sessionId, transactionModel.ID)
	return transactionModel.ID, nil
}

func (r *TransactionRepository) GetCustomerLoyaltyId(customerId string, sessionId string) (string, error) {
	var loyaltyId string
	err := r.db.Table("CUSTOMER_LOYALTY").Select("LOYALTY_ID").Where("CUSTOMER_ID = ?", customerId).Scan(&loyaltyId).Error
	if err != nil {
		return "", err
	}
	return loyaltyId, nil
}

func (r *TransactionRepository) UpdateLoyaltyBalance(loyaltyId string, points int, sessionId string) error {
	
	result := r.db.Exec("UPDATE CUSTOMER_LOYALTY SET BALANCE = BALANCE + ? WHERE LOYALTY_ID = ?", points, loyaltyId)
	err := result.Error
	if err != nil {
		log.Printf("%s : Error updating loyalty balance: %v", sessionId, err)
		return err
	}
	log.Printf("%s : Loyalty balance updated for ID: %s", sessionId, loyaltyId)
	return nil
}

func (r *TransactionRepository) ReduceLoyaltyPoints(loyaltyId string, sessionId string) (int, error) {
	var points int
	result := r.db.Exec("UPDATE CUSTOMER_LOYALTY SET BALANCE = BALANCE - ? WHERE LOYALTY_ID = ?", points, loyaltyId)
	if err := result.Error; err != nil {
		log.Printf("%s : Error reducing loyalty points: %v", sessionId, err)
		return 0, err
	}
	return points, nil
}

func (r *TransactionRepository) GetCustomerTransactionHistory(customerId string, sessionId string) (*models.Transaction, error) {
	var transactionHistory models.Transaction
	err := r.db.Table("TRANSACTIONS").Where("CUSTOMER_ID = ?", customerId).Find(&transactionHistory).Error
	if err != nil {
		log.Printf("%s : Error fetching transaction history: %v", sessionId, err)
		return nil, err
	}
	return &transactionHistory, nil
}