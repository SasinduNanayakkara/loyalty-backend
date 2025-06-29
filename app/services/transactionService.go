package services

import (
	"log"

	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
)





type TransactionServiceInterface interface {
	AccumulateLoyaltyPoints(orderId string, loyaltyId string, sessionId string) (dtos.AccumulateLoyaltyResponseDto, error)
	CreateNewOrder(transactionDto dtos.TransactionDto, sessionId string) (string, error)
	MakePayment(transactionDto dtos.TransactionDto, sessionId string) error
}

type TransactionService struct {
	repo *repositories.TransactionRepository
	loyaltyService TransactionServiceInterface
}

func NewTransactionService(repo *repositories.TransactionRepository, loyaltyService TransactionServiceInterface) *TransactionService {
	return &TransactionService{repo: repo, loyaltyService: loyaltyService}
}

func (s *TransactionService) CreateTransaction(transactionDto dtos.TransactionDto, sessionId string) (models.Transaction, error) {


	//create order
	customerLoyaltyId, err := s.repo.GetCustomerLoyaltyId(transactionDto.CustomerId, sessionId)
	if err != nil {
		return models.Transaction{}, err
	}

	transactionDto.LoyaltyAccountId = customerLoyaltyId
	if transactionDto.LoyaltyAccountId == "" {
		return models.Transaction{}, nil
	}
	orderId, err := s.loyaltyService.CreateNewOrder(transactionDto, sessionId)
	if err != nil {
		log.Printf("%s : Error creating new order: %v", sessionId, err)
		return models.Transaction{}, err
	}
	transactionDto.OrderId = orderId

	//make payment
	err = s.loyaltyService.MakePayment(transactionDto, sessionId)
	if err != nil {
		log.Printf("%s : Error making payment: %v", sessionId, err)
		return models.Transaction{}, err
	}

	//accumulate loyalty points
	var loyaltyResponse dtos.AccumulateLoyaltyResponseDto
	loyaltyResponse, err = s.loyaltyService.AccumulateLoyaltyPoints(orderId, customerLoyaltyId, sessionId)
	if err != nil {
		log.Printf("%s : Error accumulating loyalty points: %v", sessionId, err)
		return models.Transaction{}, err
	}

	balance, err := loyaltyResponse.AccumulatedPoints.Points, nil
	if err != nil {
		log.Printf("%s : Error getting accumulated points: %v", sessionId, err)
		return models.Transaction{}, err
	}

	transactionModel := &models.Transaction{
		ID:          sessionId,
		CustomerId:  transactionDto.CustomerId,
		Amount:      transactionDto.Amount,
		Description: transactionDto.Description,
		Currency:    transactionDto.Currency,
		Quantity:    transactionDto.Quantity,
		LoyaltyAccountId: transactionDto.LoyaltyAccountId,
		OrderId:   orderId,
	}
	_, err = s.repo.CreateTransaction(*transactionModel, sessionId)
	if err != nil {
		log.Printf("%s : Error creating transaction: %v", sessionId, err)
		return models.Transaction{}, err
	}
	err = s.repo.UpdateLoyaltyBalance(customerLoyaltyId, balance, sessionId)
	if err != nil {
		log.Printf("%s : Error updating loyalty balance: %v", sessionId, err)
		return models.Transaction{}, err
	}

	return *transactionModel, nil
}
