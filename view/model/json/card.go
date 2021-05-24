package json

// CardMemberListResponse ...
type CardMemberListResponse struct {
	Result Result `json:"result"`
}

// Result ...
type Result struct {
	MaxPage                 int                        `json:"max_page"`
	MemberList              []*MemberList              `json:"member_list"`
	MemberID                string                     `json:"member_id"`
	MemberName              string                     `json:"member_name"`
	MemberKana              string                     `json:"member_kana"`
	MemberRoman             string                     `json:"member_roman"`
	CICResultStatus         string                     `json:"CIC_result_status"`
	UsageStatus             string                     `json:"usage_status"`
	Card                    string                     `json:"card"`
	Birthday                string                     `json:"birthday"`
	PhoneNumber             string                     `json:"phone_number"`
	StreetAddress           string                     `json:"street_addr"`
	DeliveryDestination     string                     `json:"delivery_destination"`
	CardDesign              string                     `json:"card_design"`
	ContractNumber          string                     `json:"contract_number"`
	CardStatus              CardStatus                 `json:"card_status"`
	Notices                 string                     `json:" notices"`
	IssueDate               string                     `json:" issue_date"`
	ConfirmationHistoryList []*ConfirmationHistoryInfo `json:" confirmation_history_list"`
}

// MemberList ...
type MemberList struct {
	MemberName     string `json:"member_name"`
	MemberKana     string `json:"member_kana"`
	MemberRoman    string `json:"member_roman"`
	ContractNumber string `json:"contract_number"`
	MemberID       string `json:"member_id"`
	PhoneNumber    string `json:"phone_number"`
}

// CardStatus ...
type CardStatus struct {
	CardStatus01 string `json:"card_status01"`
	CardStatus02 string `json:"card_status02"`
	CardStatus03 string `json:"card_status03"`
	CardStatus04 string `json:"card_status04"`
	CardStatus05 string `json:"card_status05"`
	CardStatus06 string `json:"card_status06"`
	CardStatus07 string `json:"card_status07"`
	CardStatus08 string `json:"card_status08"`
	CardStatus09 string `json:"card_status09"`
	CardStatus10 string `json:"card_status10"`
	CardStatus11 string `json:"card_status11"`
	CardStatus12 string `json:"card_status12"`
	CardStatus13 string `json:"card_status13"`
	CardStatus14 string `json:"card_status14"`
	CardStatus15 string `json:"card_status15"`
}

// ConfirmationHistoryList ...
type ConfirmationHistoryInfo struct {
	CoordinationID string   `json:"coordination_id"`
	DateTime       string   `json:"date_time"`
	Result         string   `json:"result"`
	NgReason       []string `json:"ng_reason"`
}

// CardHistoryUsageResponse ...
type CardHistoryUsageResponse struct {
	Result HistoryUsageResult `json:"result"`
}

// HistoryUsageResult ...
type HistoryUsageResult struct {
	TotalUsageAmount   int                 `json:"total_usage_amount"`
	TotalConfirmAmount int                 `json:"total_confirm_amount"`
	UsageHistoryList   []*UsageHistoryInfo `json:"usage_history_list"`
}

// UsageHistoryInfo ...
type UsageHistoryInfo struct {
	UsageDay           string  `json:"usage_day"`
	Category           string  `json:"category"`
	Destination        string  `json:"destination"`
	UsageAmount        float32 `json:"usage_amount"`
	UsageConfirmAmount float32 `json:"usage_confirm_amount"`
	ConfirmAmountDate  string  `json:"confirm_amount_date"`
	UsageSection       string  `json:"usage_section"`
	UsageSectionName   string  `json:"usage_section_name"`
}

// CardHistoryPaymentResponse ...
type CardHistoryRepaymentResponse struct {
	RepayMonth           string             `json:"repay_month"`
	RepaymentDetailsList []*RepaymentDetail `json:"repayment_details_list"`
	InterestDetailList   []*InterestDetail  `json:"interest_detail_list"`
}

// RepaymentDetail ...
type RepaymentDetail struct {
	RepaymentID              string `json:"repayament_id"`
	RepaymentAmount          string `json:"repayment_amount"`
	TotalRepaymentBalancePre string `json:"total_repayment_balance_pre"`
	TotalRepaymentBalanceFix string `json:"total_repayment_balance_fix"`
	TotalInterest            string `json:"total_interest"`
	InterestAppropriation    string `json:"interest_appropriation"`
	PrincipalAppropriation   string `json:"principal_appropriation"`
	RepaymentDateTime        string `json:"repayment_datetime"`
}

// InterestDetail ...
type InterestDetail struct {
	CalDate       string `json:"cal_date"`
	Principal     string `json:"principal"`
	Interest      string `json:"interest"`
	TotalInterest string `json:"total_interest"`
}

// BFF CardGetHistoryNotification ...
// CardHistoryNotificationResponse ...
type CardHistoryNotificationResponse struct {
	Result CardHistoryNotificationResult `json:"result"`
}

// CardHistoryNotificationResult ...
type CardHistoryNotificationResult struct {
	MaxPage           int                  `json:"max_page"`
	MemberMessageList []*MemberMessageInfo `json:"member_message_list"`
}

// MemberMessageInfo ...
type MemberMessageInfo struct {
	MessageID       string `json:"message_id"`
	MessageISection string `json:"message_section"`
	ClubName        string `json:"club_name"`
	DateTime        string `json:"date_time"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

type CardHistoryBenefitResponse struct {
	Result []*CardHistoryBenefitResult `json:"result"`
}
type CardHistoryBenefitResult struct {
	ClubName                 string                      `json:"club_name"`
	ClubOwnerName            string                      `json:"club_owner_name"`
	ClubUsageAmount          int                         `json:"club_usage_amount"`
	ClubUsageList            []string                    `json:"club_usage_list"`
	RewardAchievementHistory []*RewardAchievementHistory `json:"reward_achievement_history_list"`
}

type RewardAchievementHistory struct {
	AchievementRewardName string `json:"achievement_reward_name"`
	AchievementDate       string `json:"achievement_date"`
}
