package dtos

type TransactionDto struct {
	CustomerId string `json:"customer_id" validate:"required"`
	Amount     int    `json:"amount" validate:"required"`
	Description string `json:"description" validate:"required"`
	Currency   string `json:"currency" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
	LoyaltyAccountId string `json:"loyalty_account_id" validate:"required"`
	OrderId    string `json:"order_id" validate:"required"`
}