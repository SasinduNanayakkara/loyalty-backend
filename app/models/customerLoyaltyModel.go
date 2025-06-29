package models

type CustomerLoyalty struct {
	ID          string `json:"id" gorm:"primaryKey"`
	CustomerID  string `json:"customer_id" gorm:"not null"`
	LoyaltyID   string `json:"loyalty_id" gorm:"not null;unique"`
	Balance      int    `json:"points" gorm:"default:0"`
	CreatedAt   string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string `json:"updated_at" gorm:"autoUpdateTime"`
}