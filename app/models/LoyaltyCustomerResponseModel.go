package models

type LoyaltyCustomerResponse struct {
	Customer struct {
		ID             string `json:"id"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		Nickname       string `json:"nickname"`
		EmailAddress   string `json:"email_address"`
		PhoneNumber    string `json:"phone_number"`
		Preferences    struct {
			EmailUnsubscribed bool `json:"email_unsubscribed"`
		} `json:"preferences"`
		CreationSource string `json:"creation_source"`
		Version        int    `json:"version"`
	} `json:"customer"`
}