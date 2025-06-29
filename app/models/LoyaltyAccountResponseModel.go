package models



type LoyaltyMapping struct {
	ID          string    `json:"id"`
	PhoneNumber string    `json:"phone_number"`
}

type LoyaltyAccount struct {
	ID             string         `json:"id"`
	Mapping        LoyaltyMapping `json:"mapping"`
	ProgramID      string         `json:"program_id"`
	Balance        int            `json:"balance"`
	LifetimePoints int            `json:"lifetime_points"`
	CustomerID     string         `json:"customer_id"`
}

type LoyaltyAccountResponseModel struct {
	LoyaltyAccount LoyaltyAccount `json:"loyalty_account"`
}