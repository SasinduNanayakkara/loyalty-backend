package services

import (
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
	"github.com/sasinduNanayakkara/loyalty-backend/app/utils"
)

type LoyaltyAppServiceInterface interface {
	CreateNewLoyaltyCustomer(customer models.Customer, sessionId string) (string, error)
	CreateNewLoyaltyAccount(loyaltyCustomerId string, phoneNumber string, sessionId string) (*models.LoyaltyAccountResponseModel, error)
}

type CustomerService struct {
	repo           repositories.CustomerRepository
	loyaltyService LoyaltyAppServiceInterface
}

func NewCustomerService(repo repositories.CustomerRepository, loyaltyService LoyaltyAppServiceInterface) *CustomerService {
	return &CustomerService{repo: repo, loyaltyService: loyaltyService}
}

func (s *CustomerService) CreateNewCustomer(customer models.Customer, sessionId string) error {
	
	customer.ID = sessionId
	customer.Password = utils.GenerateHashedPassword(customer.Password)

	if error := s.repo.CreateNewCustomerRepository(customer, sessionId); error != nil {
		return error
	}

	loyaltyCustomer, err := s.loyaltyService.CreateNewLoyaltyCustomer(customer, sessionId)
	if err != nil {
		return err
	}

	var loyaltyAccount *models.LoyaltyAccountResponseModel

	if loyaltyCustomer != "" {
		loyaltyAccount, err = s.loyaltyService.CreateNewLoyaltyAccount(loyaltyCustomer, customer.Phone_number, sessionId)
		if err != nil {
			return err
		}
	}

	if loyaltyAccount != nil {
		if err := s.repo.CreateNewLoyaltyCustomer(loyaltyAccount, customer.ID, sessionId); err != nil {
			return err
		}
	}

	return nil
}