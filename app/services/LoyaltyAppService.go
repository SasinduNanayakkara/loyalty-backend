package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func init() {
	loyaltyBaseUrl = os.Getenv("LOYALTY_API_URL")
	loyaltyAccessToken = os.Getenv("LOYALTY_ACCESS_TOKEN")
	squareVersion = os.Getenv("SQUARE_VERSION")
	loyaltyProgramId = os.Getenv("LOYALTY_PROGRAM_ID")
}

func (s *LoyaltyAppService) CreateNewLoyaltyCustomer(customer models.Customer, sessionId string) (string, error) {

	body := map[string]interface{}{
		"idempotency_key": sessionId,
		"email_address":   customer.Email,
		"nickname":        customer.Name,
		"phone_number":    customer.PhoneNumber,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	

	httpReq, error := http.NewRequest("POST", loyaltyBaseUrl+"/customers", bytes.NewBuffer(bodyBytes) )

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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s : Error reading response body: %v", sessionId, err)
		return "", err
	}

	if response.StatusCode != http.StatusOK{
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

func (s *LoyaltyAppService) CreateNewLoyaltyAccount(customerLoyaltyId string, phoneNumber string, sessionId string) (*models.LoyaltyAccountResponseModel, error) {

		body := map[string]interface{}{
			"idempotency_key": sessionId,
			"customer_id":     customerLoyaltyId,
			"loyalty_account": map[string]interface{}{
				"program_id": loyaltyProgramId,
				"mapping": map[string]interface{}{
					"phone_number": phoneNumber,
				},
			},
		}
	
		bodyBytes, err := json.Marshal(body)
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
	
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("%s : Error reading response body: %v", sessionId, err)
			return nil, err
		}
	
		if response.StatusCode != http.StatusOK {
			log.Printf("%s : Error creating loyalty account: %s", sessionId, responseBody)
			return nil, fmt.Errorf("error creating loyalty account: %s", responseBody)
		}
	
		var accountResponse models.LoyaltyAccountResponseModel
	
		if err := json.Unmarshal(responseBody, &accountResponse); err != nil {
			log.Printf("%s : Error unmarshalling loyalty account response: %v", sessionId, err)
			return nil, err
		}
	
		return &accountResponse, nil
	}