package dtos


type LoyaltyRewardResponseDto struct {
	Reward LoyaltyReward `json:"reward"`
}

type LoyaltyReward struct {
	ID               string    `json:"id"`
	Status           string    `json:"status"`
	LoyaltyAccountId string    `json:"loyalty_account_id"`
	RewardTierId     string    `json:"reward_tier_id"`
	Points           int       `json:"points"`
	OrderId          string    `json:"order_id"`
}
