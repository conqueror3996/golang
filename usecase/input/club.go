package input

type ClubUpdate struct {
	ClubID           string `json:"club_id"`
	ContentType      string `json:"content_type"`
	ClubImageURL     string `json:"club_image_url"`
	CoverImageURL    string `json:"cover_image_url"`
	ClubName         string `json:"club_name"`
	ClubOwnerID      string `json:"club_owner_id"`
	Description      string `json:"description"`
	ThemeColor       string `json:"theme_color"`
	CategoryCode     string `json:"category_code"`
	ReleaseStatus    string `json:"release_status"`
	ReleaseStartDate string `json:"release_start_date"`
	ReleaseEndDate   string `json:"release_end_date"`
	MaxMembers       string `json:"max_members"`
	ClubParticipants string `json:"club_participants"`
	MaxMembershipNo  string `json:"max_membersip_no"`
}

// RewardRegistration ...
type RewardRegistration struct {
	ClubID             string `json:"club_id"`
	RewardName         string `json:"reward_name"`
	Section            string `json:"section"`
	AchievementAmount  string `json:"achievement_amount"`
	Created            string `json:"created"`
	ShipmentSource     string `json:"shipment_source"`
	DistributionPeriod string `json:"distribution_period"`
	DistributionStatus string `json:"distribution_status"`
	TotalOrderQuantity string `json:"total_order_quantity"`
	RewardContent      string `json:"reward_content"`
	RewardImage        string `json:"reward_image"`
	RewardThumbnail    string `json:"reward_thumbnail"`
}

type RewardUpdate struct {
	RewardID           string `json:"reward_id"`
	ContentType        string `json:"content_type"`
	ClubID             string `json:"club_id"`
	RewardName         string `json:"reward_name"`
	DistributionPeriod string `json:"distribution_period"`
	DistributionStatus string `json:"distribution_status"`
	TotalOrderQuantity string `json:"total_order_quantity"`
	DownloadURL        string `json:"downloadURL"`
	RewardContent      string `json:"reward_content"`
	RewardImage        string `json:"reward_image"`
	RewardThumbnail    string `json:"reward_thumbnail"`
}

type RewardInquiry struct {
	ClubID string `json:"club_id"`
}
