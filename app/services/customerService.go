package services

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
	"github.com/sasinduNanayakkara/loyalty-backend/app/repositories"
	"github.com/sasinduNanayakkara/loyalty-backend/app/utils"
	"github.com/sasinduNanayakkara/loyalty-backend/config"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	config.LoadEnv()
}
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

func (s *CustomerService) CustomerLogin(loginDto dtos.LoginDto, sessionId string) (*dtos.LoginResponseDto, error) {

	customer, err := s.repo.GetCustomerByEmail(loginDto.Email, sessionId)

	if err != nil {
		return nil, err
	}

	log.Printf("%s : Customer fetched successfully: %s", sessionId, customer.Password)


	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(loginDto.Password))

	if err != nil {
		log.Printf("%s : Invalid email or password: %s", sessionId, customer.Password)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": customer.ID,
		"email": customer.Email,
		"name": customer.Name,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Printf("%s : Error signing JWT token: %v", sessionId, err)
		return nil, err
	}

	loginResponse := &dtos.LoginResponseDto{
		Token:    tokenString,
		Customer: *customer,
	}

	log.Printf("%s : Customer login successful: %s", sessionId, loginResponse)
	return loginResponse, nil
}