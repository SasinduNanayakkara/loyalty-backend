package models



type Customer struct {
	ID           string `json:"id" gorm:"primaryKey;column:ID"`
	Name         string `json:"name" gorm:"not null;column:NAME"`
	Email        string `json:"email" gorm:"not null;column:EMAIL"`
	Phone_number string `json:"phone_number" gorm:"not null;column:PHONE"`
	Password     string `json:"password" gorm:"not null;column:PASSWORD"`
	Status       string `json:"status" gorm:"default:'active';column:STATUS"`
}

func (Customer) TableName() string {
	return "CUSTOMERS"
}