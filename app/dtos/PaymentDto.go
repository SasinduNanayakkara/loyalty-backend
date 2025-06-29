package dtos

type PaymentDto struct {
	CustomerId       string `json:"customer_id" validate:"required"`
	Amount           int    `json:"amount" validate:"required"`
	Currency		 string `json:"currency" validate:"required"`
	SourceId		 string `json:"source_id" validate:"required"`
	OrderId		  	 string `json:"order_id" validate:"required"`
}