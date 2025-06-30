package dtos

type LoyaltyOrderResponseDto struct {
	Id 		string `json:"id"`
	Location_Id string `json:"location_id"`
	Line_Items []struct {}
}

type LoyaltyOrderWrapper struct {
	Order LoyaltyOrderResponseDto `json:"order"`
}