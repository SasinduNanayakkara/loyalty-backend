package dtos

type AccumulateLoyaltyResponseDto struct {
	Id 		string `json:"id"`
	AccumulatedPoints AccumulateLoyaltyResponse    `json:"accumulated_points"`
}

type AccumulateLoyaltyResponse struct {
	LoyaltyProgramId string `json:"loyalty_program_id"`
	Points 		 int    `json:"points"`
	OrderId 	 string `json:"order_id"`

}