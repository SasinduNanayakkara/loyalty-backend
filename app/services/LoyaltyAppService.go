package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sasinduNanayakkara/loyalty-backend/app/dtos"
	"github.com/sasinduNanayakkara/loyalty-backend/app/models"
)

type LoyaltyAppService struct {
}

func NewLoyaltyAppService() *LoyaltyAppService {
	return &LoyaltyAppService{}
}

var loyaltyBaseUrl string
var loyaltyAccessToken string
var squareVersion string
var loyaltyProgramId string
var loyaltyLocationId string
var sourceId string = "cnon:card-nonce-ok"
var rewardTierId string = "647ed5d4-512f-4c01-9fd6-67b63c4bb0a2"

func init() {
	loyaltyBaseUrl = os.Getenv("LOYALTY_API_URL")
	if loyaltyBaseUrl == "" {
		loyaltyBaseUrl = "https://connect.squareupsandbox.com/v2"
	}
	loyaltyAccessToken = os.Getenv("LOYALTY_ACCESS_TOKEN")
	if loyaltyAccessToken == "" {
		loyaltyAccessToken = "EAAAl46zVNOhYm5YeuDebadRdnrjDPyFtLBvEypBI8hPiNQa8JPPmTT_TOygGg6c"
	}
	squareVersion = os.Getenv("SQUARE_VERSION")
	if squareVersion == "" {
		squareVersion = "2025-06-18"
	}
	loyaltyProgramId = os.Getenv("LOYALTY_PROGRAM_ID")
	if loyaltyProgramId == "" {
		loyaltyProgramId = "7e3874a3-6f99-42b4-8a4b-a3c69af5c106"
	}

	loyaltyLocationId = os.Getenv("LOYALTY_LOCATION_ID")
	if loyaltyLocationId == "" {
		loyaltyLocationId = "L0B21CBE1A66C"
	}
}

func (s *LoyaltyAppService) CreateNewLoyaltyCustomer(customer models.Customer, sessionId string) (string, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"email_address":   customer.Email,
		"nickname":        customer.Name,
		"phone_number":    customer.Phone_number,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	httpReq, error := http.NewRequest("POST", loyaltyBaseUrl+"/customers", bytes.NewBuffer(bodyBytes))

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	if error != nil {
		log.Printf("%s : Error creating new loyalty customer request: %v", sessionId, error)
		return "", error
	}

	client := &http.Client{}
	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return "", err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return "", err
	}

	log.Printf("%s : Loyalty API response: %s", sessionId, string(responseBody))

	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error creating loyalty customer: %s", sessionId, responseBody)
		return "", fmt.Errorf("error creating loyalty customer: %s", responseBody)
	}

	var customerResponse models.LoyaltyCustomerResponse

	if err := json.Unmarshal(responseBody, &customerResponse); err != nil {
		log.Printf("%s : Error unmarshalling loyalty customer response: %v", sessionId, err)
		return "", err
	}

	return customerResponse.Customer.ID, nil
}

func (s *LoyaltyAppService) CreateNewLoyaltyAccount(customerLoyaltyId string, phoneNumber string, sessionId string) (*dtos.LoyaltyAccountResponseDto, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"loyalty_account": map[string]interface{}{
			"program_id": loyaltyProgramId,
			"mapping": map[string]interface{}{
				"phone_number": phoneNumber,
			},
			"customer_id": customerLoyaltyId,
		},
	}

	bodyBytes, err := json.Marshal(body)
	log.Printf("%s : Creating new loyalty account for customer ID with phone number: %s", sessionId, bodyBytes)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", loyaltyBaseUrl+"/loyalty/accounts", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("%s : Error creating new loyalty account request: %v", sessionId, err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)

	client := &http.Client{}
	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return nil, err
	}

	log.Printf("%s : Loyalty account creation API response: %s", sessionId, string(responseBody))

	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error creating loyalty account: %s", sessionId, responseBody)
		return nil, fmt.Errorf("error creating loyalty account: %s", responseBody)
	}

	var accountResponse dtos.LoyaltyAccountResponseDto

	if err := json.Unmarshal(responseBody, &accountResponse); err != nil {
		log.Printf("%s : Error unmarshalling loyalty account response: %v", sessionId, err)
		return nil, err
	}

	return &accountResponse, nil
}

func (s *LoyaltyAppService) CreateNewOrder(transactionDto dtos.TransactionDto, sessionId string) (string, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"order": map[string]interface{}{
			"location_id": loyaltyLocationId,
			"customer_id": transactionDto.LoyaltyAccountId,
			"line_items": []map[string]interface{}{
				{
					"quantity": transactionDto.Quantity,
					"name":     transactionDto.Description,
					"base_price_money": map[string]interface{}{
						"amount":   transactionDto.Amount,
						"currency": transactionDto.Currency,
					},
				},
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	log.Printf("%s : Creating new order for transaction: %s", sessionId, bodyBytes)
	if err != nil {
		log.Printf("%s : Error marshalling order request body: %v", sessionId, err)
		return "", err
	}

	httpReq, err := http.NewRequest("POST", loyaltyBaseUrl+"/orders", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("%s : Error creating new order request: %v", sessionId, err)
		return "", err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	client := &http.Client{}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return "", err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return "", err
	}

	log.Printf("%s : Order creation API response: %s", sessionId, string(responseBody))
	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error creating order: %s", sessionId, responseBody)
		return "", fmt.Errorf("error creating order: %s", responseBody)
	}

	var orderResponse dtos.LoyaltyOrderWrapper
	if err := json.Unmarshal(responseBody, &orderResponse); err != nil {
		log.Printf("%s : Error unmarshalling order response: %v", sessionId, err)
		return "", err
	}
	log.Printf("%s : New order created with ID: %s", sessionId, orderResponse.Order.Id)

	return orderResponse.Order.Id, nil
}

func (s *LoyaltyAppService) AccumulateLoyaltyPoints(orderId string, loyaltyId string, sessionId string) (dtos.AccumulateLoyaltyResponseDto, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"accumulate_points": map[string]interface{}{
			"order_id": orderId,
		},
		"location_id": loyaltyLocationId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("%s : Error marshalling accumulate points request body: %v", sessionId, err)
		return dtos.AccumulateLoyaltyResponseDto{}, err
	}
	httpReq, err := http.NewRequest("POST", loyaltyBaseUrl+"/loyalty/accounts/"+loyaltyId+"/accumulate", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("%s : Error creating accumulate points request: %v", sessionId, err)
		return dtos.AccumulateLoyaltyResponseDto{}, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	client := &http.Client{}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return dtos.AccumulateLoyaltyResponseDto{}, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return dtos.AccumulateLoyaltyResponseDto{}, err
	}

	log.Printf("%s : Accumulate points API response: %s", sessionId, string(responseBody))
	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error accumulating loyalty points: %s", sessionId, responseBody)
		return dtos.AccumulateLoyaltyResponseDto{}, fmt.Errorf("error accumulating loyalty points: %s", responseBody)
	}

	var accumulateResponse dtos.AccumulateLoyaltyResponseDto
	if err := json.Unmarshal(responseBody, &accumulateResponse); err != nil {
		log.Printf("%s : Error unmarshalling accumulate points response: %v", sessionId, err)
		return dtos.AccumulateLoyaltyResponseDto{}, err
	}

	return accumulateResponse, nil
}

func (s *LoyaltyAppService) MakePayment(transactionDto dtos.TransactionDto, sessionId string) error {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"source_id":       sourceId,
		"order_id":        transactionDto.OrderId,
		"customer_id":     transactionDto.LoyaltyAccountId,
		"amount_money": map[string]interface{}{
			"amount":   transactionDto.Amount,
			"currency": transactionDto.Currency,
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("%s : Error marshalling payment request body: %v", sessionId, err)
		return err
	}

	httpReq, err := http.NewRequest("POST", loyaltyBaseUrl+"/payments", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("%s : Error creating new payment request: %v", sessionId, err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	client := &http.Client{}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return err
	}

	log.Printf("%s : Payment API response: %s", sessionId, string(responseBody))
	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error making payment: %s", sessionId, responseBody)
		return fmt.Errorf("error making payment: %s", responseBody)
	}

	return nil
}

func (s *LoyaltyAppService) CreateLoyaltyReward(customerLoyaltyId string, orderId string, sessionId string) (int, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"reward": map[string]interface{}{
			"loyalty_account_id": customerLoyaltyId,
			"reward_tier_id":     rewardTierId,
			"order_id":           orderId,
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("%s : Error marshalling loyalty reward request body: %v", sessionId, err)
		return 0, err
	}

	httpReq, err := http.NewRequest("POST", loyaltyBaseUrl+"/loyalty/rewards", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("%s : Error creating loyalty reward request: %v", sessionId, err)
		return 0, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	client := &http.Client{}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return 0, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return 0, err
	}

	log.Printf("%s : Loyalty reward API response: %s", sessionId, string(responseBody))
	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error creating loyalty reward: %s", sessionId, responseBody)
		return 0, fmt.Errorf("error creating loyalty reward: %s", responseBody)
	}

	loyaltyRewardResponse := dtos.LoyaltyRewardResponseDto{}
	if err := json.Unmarshal(responseBody, &loyaltyRewardResponse); err != nil {
		log.Printf("%s : Error unmarshalling loyalty reward response: %v", sessionId, err)
		return 0, err
	}
	log.Printf("%s : Loyalty reward created with ID: %s", sessionId, loyaltyRewardResponse.Reward.ID)

	return loyaltyRewardResponse.Reward.Points, nil
}

func (s *LoyaltyAppService) GetLoyaltyAccount(loyaltyId string, sessionId string) (*dtos.LoyaltyAccountResponseDto, error) {

	httpReq, err := http.NewRequest("GET", loyaltyBaseUrl+"/loyalty/accounts/"+loyaltyId, nil)
	if err != nil {
		log.Printf("%s : Error creating get loyalty account request: %v", sessionId, err)
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+loyaltyAccessToken)
	httpReq.Header.Set("Square-Version", squareVersion)
	client := &http.Client{}

	response, err := client.Do(httpReq)
	if err != nil {
		log.Printf("%s : Error sending request to loyalty API: %v", sessionId, err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return nil, err
	}

	log.Printf("%s : Get loyalty account API response: %s", sessionId, string(responseBody))
	if response.StatusCode != http.StatusOK {
		log.Printf("%s : Error getting loyalty account: %s", sessionId, responseBody)
		return nil, fmt.Errorf("error getting loyalty account: %s", responseBody)
	}

	loyaltyAccountResponse := dtos.LoyaltyAccountResponseDto{}
	if err := json.Unmarshal(responseBody, &loyaltyAccountResponse); err != nil {
		log.Printf("%s : Error unmarshalling loyalty account response: %v", sessionId, err)
		return nil, err
	}

	return &loyaltyAccountResponse, nil
}
