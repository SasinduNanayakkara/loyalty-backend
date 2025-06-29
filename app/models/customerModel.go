package models



type Customer struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Phone_number string `json:"phone_number" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Status    string `json:"status" gorm:"default:'active'"`
}

func (Customer) TableName() string {
	return "CUSTOMERS"
}