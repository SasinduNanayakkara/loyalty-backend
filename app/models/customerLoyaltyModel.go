package models

type CustomerLoyalty struct {
	CustomerID  string `json:"customer_id" gorm:"not null"`
	LoyaltyID   string `json:"loyalty_id" gorm:"not null;unique"`
	Balance      int    `json:"points" gorm:"default:0"`
	}

func (CustomerLoyalty) TableName() string {
	return "CUSTOMER_LOYALTY"
}