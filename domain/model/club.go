package model

type ClubInfo struct {
	ClubID           string `json:"club_id"`
	ClubName         string `json:"club_name"`
	ClubOwnerID      string `json:"club_owner_id"`
	Description      string `json:"description"`
	ClubImageURL     string `json:"club_image_url"`
	CoverImageURL    string `json:"cover_image_url"`
	ThemeColor       string `json:"theme_color"`
	CategoryCode     string `json:"category_code"`
	ReleaseStatus    string `json:"release_status"`
	ReleaseStartDate string `json:"release_start_date"`
	ReleaseEndDate   string `json:"release_end_date"`
	MaxMembers       string `json:"max_members"`
	ClubParticipants string `json:"club_participants"`
	MaxMembershipNo  string `json:"max_membersip_no"`
}

type ClubUploadURL struct {
	Type   int    `json:"type"`
	ClubID string `json:"clubID"`
}

// ClubUploadURLResponse ...
type ClubUploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}

type ClubUpdateResponse struct {
	ClubID        uint64 `json:"clubID"`
	ReleaseStatus int64  `json:"releaseStatus"`
}

// RewardUploadURL
type RewardUploadURL struct {
	Type     int    `json:"type"`
	RewardID string `json:"rewardID"`
}

type RewardInfo struct {
	ClubID                string `json:"clubID"`
	RewardName            string `json:"rewardName"`
	DistributionStartDate string `json:"distributionStartDate"`
	DistributionEndDate   string `json:"distributionEndDate"`
	DistributionStatus    string `json:"distributionStatus"`
	TotalOrder            string `json:"totalOrder"`
	RewardDownloadURL     string `json:"downloadURL"`
	Description           string `json:"rewardContent"`
	RewardImageURL        string `json:"rewardImage"`
	RewardThumbnailURL    string `json:"rewardThumbnail"`
}

// RewardUploadURLResponse ...
type RewardUploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}

type RewardUpdateResponse struct {
	RewardID uint64 `json:"rewardID"`
}

type RewardRegisterInfo struct {
	ClubID                string `json:"clubID"`
	RewardName            string `json:"rewardName"`
	RewardType            string `json:"rewardType"`
	AchieveAmount         string `json:"achieveAmount"`
	MadeBy                string `json:"madeBy"`
	ShipmentSource        string `json:"shipmentSource"`
	DistributionStartDate string `json:"distributionStartDate"`
	DistributionEndtDate  string `json:"distributionEndtDate"`
	DistributionStatus    string `json:"distributionStatus"`
	TotalOrder            string `json:"totalOrder"`
	Description           string `json:"description"`
	RewardDownloadURL     string `json:"rewardDownloadURL"`
	ThumbnailImageURL     string `json:"thumbnailImageURL"`
}

type RewardRegisterResponse struct {
}

// ResponseClubs ...
type ResponseClubs struct {
	CaseTotal int64      `json:"caseTotal"`
	IsNext    bool       `json:"isNext"`
	ClubList  []ItemClub `json:"clubList"`
}

// ItemClub ...
type ItemClub struct {
	ClubID            int64        `json:"clubID"`
	ClubName          string       `json:"clubName"`
	ClubOwnerID       int64        `json:"clubOwnerID"`
	ClubOwnerName     string       `json:"clubOwnerName"`
	ClubOwnerImageURL string       `json:"clubOwnerImageUrl"`
	Description       string       `json:"description"`
	ClubImageURL      string       `json:"clubImageUrl"`
	CoverImageURL     string       `json:"coverImageUrl"`
	RewardDownloadURL string       `json:"rewardDownloadUrl"`
	ThemeColor        string       `json:"themeColor"`
	CategoryCode      int64        `json:"categoryCode"`
	CategoryName      string       `json:"categoryName"`
	ReleaseStatus     string       `json:"releaseStatus"`
	ReleaseStartDate  int64        `json:"releaseStartDate"`
	ReleaseEndDate    int64        `json:"releaseEndDate"`
	MaxMembers        int64        `json:"maxMembers"`
	ClubParticipants  int64        `json:"clubParticipants"`
	MaxMembershipNo   int64        `json:"maxMembershipNo"`
	RewardList        []ItemReward `json:"rewardList"`
	DesignList        []ItemDesign `json:"designList"`
}

// ItemReward ...
type ItemReward struct {
	RewardID                         int64  `json:"rewardID"`
	RewardName                       string `json:"rewardName"`
	Description                      string `json:"description"`
	RewardImageURL                   string `json:"rewardImageUrl"`
	ThumbnailImageURL                string `json:"thumbnailImageUrl"`
	RewardDownloadURL                string `json:"rewardDownloadUrl"`
	RewardType                       string `json:"rewardType"`
	MadeBy                           string `json:"madeBy"`
	ShipmentSource                   string `json:"shipmentSource"`
	AchieveAmount                    int64  `json:"achieveAmount"`
	DistributionStatus               string `json:"distributionStatus"`
	DistributionStartDate            int64  `json:"distributionStartDate"`
	DistributionEndDate              int64  `json:"distributionEndDate"`
	ReleaseStatus                    string `json:"releaseStatus"`
	TotalOrder                       int64  `json:"totalOrder"`
	LimitEditionFlag                 string `json:"limitEditionFlag"`
	NumberOfPeopleWhoAchievedRewards uint64 `json:"numberPeopleAchievedReward"`
}

// ItemDesign ...
type ItemDesign struct {
	DesignID              int64  `json:"designID"`
	DesignName            string `json:"designName"`
	DesignImageURL        string `json:"designImageUrl"`
	IssuanceCost          int64  `json:"issuanceCost"`
	IssuanceLimit         int64  `json:"issuanceLimit"`
	DistributionStartDate int64  `json:"distributionStartDate"`
	DistributionEndDate   int64  `json:"distributionEndDate"`
	NumberOfApplications  int64  `json:"numberOfApplications"`
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
