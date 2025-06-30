package dtos


type LoyaltyAccountResponseDto struct {
	LoyaltyAccount LoyaltyAccount `json:"loyalty_account"`
}

type LoyaltyAccount struct {
	ID             string                `json:"id"`
	ProgramId      string                `json:"program_id"`
	Balance        int                   `json:"balance"`
	LifetimePoints int                   `json:"lifetime_points"`
	CustomerId     string                `json:"customer_id"`
	Mapping        LoyaltyAccountMapping `json:"mapping"`
}

type LoyaltyAccountMapping struct {
	ID          string    `json:"id"`
	PhoneNumber string    `json:"phone_number"`
}
