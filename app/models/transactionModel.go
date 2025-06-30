package models

type Transaction struct {
	ID 	   string `json:"id" gorm:"primaryKey;column:ID"`
	CustomerId  string `json:"customer_id" gorm:"not null;column:CUSTOMER_ID"`
	Amount      int    `json:"amount" gorm:"not null;column:AMOUNT"`
	Description string `json:"description" gorm:"not null;column:DESCRIPTION"`
	Currency    string `json:"currency" gorm:"not null;column:CURRENCY"`
	Quantity    int    `json:"quantity" gorm:"not null;column:QUANTITY"`
	LoyaltyAccountId string `json:"loyalty_account_id" gorm:"not null;column:LOAYLTY_ACCOUNT_ID"`
	OrderId   string `json:"order_id" gorm:"not null;column:ORDER_ID"`
}

func (Transaction) TableName() string {
	return "TRANSACTION"
}