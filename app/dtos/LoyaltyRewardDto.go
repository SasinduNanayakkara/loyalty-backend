package dtos

type LoyaltyRewardDto struct {
	CustomerId       string `json:"customer_id" validate:"required"`
	OrderId 		string `json:"order_id" validate:"required"`
	LoyaltyAccountId string `json:"loyalty_account_id" validate:"required"`
	Amount     int    `json:"amount" validate:"required"`
	Description string `json:"description" validate:"required"`
	Currency   string `json:"currency" validate:"required"`
	Quantity   string    `json:"quantity" validate:"required"`
}