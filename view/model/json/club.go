package json

type ClubUploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}

type RewardUploadURLResponse struct {
	UploadURL string `json:"upload_urls"`
}

type RewardRegistrationResponse struct {
}

// ConvertRewardInfo ...
type ConvertRewardInfo struct {
	RewardName          string `json:"reward_name"`
	RewardAmount        int64  `json:"reward_amount"`
	RewardDetail        string `json:"reward_detail"`
	RewardThumbnail     string `json:"reward_thumbnail"`
	DistributionStatus  string `json:"distribution_status"`
	RewardID            int64  `json:"reward_id"`
	Section             string `json:"section"`
	AchievementAmount   int64  `json:"achievement_amount"`
	Created             string `json:"created"`
	ShipmentSource      string `json:"shipment_source"`
	DistributionPeriod  string `json:"distribution_period"`
	AchievementNumber   uint64 `json:"achievement_number"`
	RewardStockQuantity int64  `json:"reward_stock_quantity"`
	DownloadURL         string `json:"downloadURL"`
	RewardContent       string `json:"reward_content"`
}

type RewardListResponse struct {
	Result RewardResult `json:"result"`
}

// OperatorResult ...
type RewardResult struct {
	RewardList []*ConvertRewardInfo `json:"reward_list"`
}
