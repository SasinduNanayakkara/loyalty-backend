package dtos



type AccumulateLoyaltyResponseDto struct {
	Events []LoyaltyEvent `json:"events"`
}

type LoyaltyEvent struct {
	ID               string                  `json:"id"`
	Type             string                  `json:"type"`
	AccumulatePoints AccumulatePointsDetails `json:"accumulate_points"`
	LoyaltyAccountId string                  `json:"loyalty_account_id"`
	LocationId       string                  `json:"location_id"`
	Source           string                  `json:"source"`
}

type AccumulatePointsDetails struct {
	LoyaltyProgramId string `json:"loyalty_program_id"`
	Points           int    `json:"points"`
	OrderId          string `json:"order_id"`
}
