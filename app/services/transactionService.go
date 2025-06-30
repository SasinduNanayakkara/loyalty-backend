package services

import (
	"log"
	"strconv"

	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
)

type TransactionServiceInterface interface {
	AccumulateLoyaltyPoints(orderId string, loyaltyId string, sessionId string) (dtos.AccumulateLoyaltyResponseDto, error)
	CreateNewOrder(transactionDto dtos.TransactionDto, sessionId string) (string, error)
	MakePayment(transactionDto dtos.TransactionDto, sessionId string) error
	CreateLoyaltyReward(customerLoyaltyId string, orderId string, sessionId string) (int, error)
}


type TransactionService struct {
	repo           *repositories.TransactionRepository
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

	balance, err := loyaltyResponse.Events[0].AccumulatePoints.Points, nil
	if err != nil {
		log.Printf("%s : Error getting accumulated points: %v", sessionId, err)
		return models.Transaction{}, err
	}
	log.Printf("%s : Accumulated points: %d", sessionId, balance)
	quantityInt, err := strconv.Atoi(transactionDto.Quantity)
	if err != nil {
		log.Printf("%s : Error converting quantity to int: %v", sessionId, err)
		return models.Transaction{}, err
	}

	transactionModel := &models.Transaction{
		ID:               sessionId,
		CustomerId:       transactionDto.CustomerId,
		Amount:           transactionDto.Amount,
		Description:      transactionDto.Description,
		Currency:         transactionDto.Currency,
		Quantity:         quantityInt,
		LoyaltyAccountId: transactionDto.LoyaltyAccountId,
		OrderId:          orderId,
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

func (s *TransactionService) RedeemLoyaltyPoints(redeemDto dtos.LoyaltyRewardDto, sessionId string) (int, error) {

	//create order
	customerLoyaltyId, err := s.repo.GetCustomerLoyaltyId(redeemDto.CustomerId, sessionId)
	if err != nil {
		return 0, err
	}

	redeemDto.LoyaltyAccountId = customerLoyaltyId
	if redeemDto.LoyaltyAccountId == "" {
		return 0, nil
	}
	transactionDto := dtos.TransactionDto{
		CustomerId:       redeemDto.CustomerId,
		Amount:           redeemDto.Amount,
		Description:      redeemDto.Description,
		Currency:         redeemDto.Currency,
		Quantity:         redeemDto.Quantity,
		LoyaltyAccountId: redeemDto.LoyaltyAccountId,
	}
	orderId, err := s.loyaltyService.CreateNewOrder(transactionDto, sessionId)
	if err != nil {
		log.Printf("%s : Error creating new order: %v", sessionId, err)
		return 0, err
	}
	transactionDto.OrderId = orderId

	//redeem loyalty points
	points, err := s.loyaltyService.CreateLoyaltyReward(customerLoyaltyId, orderId, sessionId)
	if err != nil {
		log.Printf("%s : Error creating loyalty reward: %v", sessionId, err)
		return 0, err
	}

	_, err = s.repo.ReduceLoyaltyPoints(customerLoyaltyId, sessionId)
	if err != nil {
		log.Printf("%s : Error reducing loyalty points: %v", sessionId, err)
		return 0, err
	}

	//make the rest of the payment
	err = s.loyaltyService.MakePayment(transactionDto, sessionId)
	if err != nil {
		log.Printf("%s : Error making payment: %v", sessionId, err)
		return 0, err
	}

	return points, nil
}
