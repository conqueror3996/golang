package model

import "time"

// BFF CardGetHistoryNotification ...
// CardHistoryNotification ...
type CardHistoryNotification struct {
	MemberID    string `json:"member_id"`
	SortKey     string `json:"sort_key"`
	SortOrder   string `json:"sort_order"`
	PageNum     int    `json:"page_num"`
	PageSize    int    `json:"page_size"`
	DeletedData string `json:"deleted_data"`
}

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

// NoticeUserInfo
type NoticeUserInfo struct {
	Total               int32       `json:"total"`
	IsNext              bool        `json:"isNext"`
	UserInformationList []*UserInfo `json:"userInformationList"`
}

// UserInfo ...
type UserInfo struct {
	ClubID          string `json:"clubID"`
	NoticeID        string `json:"noticeID"`
	NoticeType      string `json:"noticeType"`
	SendingDateTime string `json:"sendingDatetime"`
	Subject         string `json:"subject"`
	NoticeText      string `json:"noticeText"`
	ClubName        string `json:"clubName"`
}

// END CardGetHistoryNotification

// START CardGetMemberList
// CardMemberListResponse ...
type CardMemberListResponse struct {
	Result Result `json:"result"`
}

// UsersResponse
type UsersReponse struct {
	TotalCase int             `json:"totalCase"`
	IsNext    bool            `json:"isNext"`
	UserList  []*B1MemberInfo `json:"userList"`
}

// B1MemberInfo ...
type B1MemberInfo struct {
	UserID               string `json:"userID"`
	Surname              string `json:"surname"`
	Name                 string `json:"name"`
	FullnameKanji        string `json:"fullnameKanji"`
	KanaSurname          string `json:"kanaSurname"`
	KanaName             string `json:"kanaName"`
	FullnameKana         string `json:"fullnameKana"`
	RomanjiSurname       string `json:"romanjiSurname"`
	RomanjiName          string `json:"romanjiName"`
	FullnameRomaji       string `json:"fullnameRomaji"`
	CICResultStatus      string `json:"cicScreeningResult"`
	UsageStatus          string `json:"usageStatus"`
	YearOfBirth          string `json:"yearOfBirth"`
	MonthOfBirth         string `json:"monthOfBirth"`
	DayOfBirth           string `json:"dayOfBirth"`
	PhoneNumber          string `json:"phoneNumber"`
	PostalCode           string `json:"postalCode"`
	Prefectures          string `json:"prefectures"`
	Municipality         string `json:"municipality"`
	StreetAddress        string `json:"streetAddress"`
	SendingPostalCode    string `json:"sendingPostalCode"`
	SendingPrefectures   string `json:"sendingPrefectures"`
	SendingMunicipality  string `json:"sendingMunicipality"`
	SendingStreetAddress string `json:"sendingStreetAddress"`
	Notices              string `json:"remarks"`
}

// CardAndIssuanceHistoryCards card list
type CardAndIssuanceHistoryCards struct {
	CardList []CardNoPin `json:"cardList"`
}

// CardNoPin ...
type CardNoPin struct {
	CardID                   string                      `json:"cardID"`
	UserID                   string                      `json:"userID"`
	ClubID                   uint64                      `json:"clubID"`
	DesignID                 uint64                      `json:"designID"`
	BrandID                  string                      `json:"brandID"`
	GradeID                  string                      `json:"gradeID"`
	CardTypeID               string                      `json:"cardTypeID"`
	DesignCode               string                      `json:"designCode"`
	NameKana                 string                      `json:"nameKana"`
	NameKanji                string                      `json:"nameKanji"`
	NameRomaji               string                      `json:"nameRomaji"`
	PhoneNumber              string                      `json:"phoneNumber"`
	Birthday                 string                      `json:"birthday"`
	Gender                   string                      `json:"gender"`
	PostalCode               string                      `json:"postalCode"`
	AddressKanji1            string                      `json:"addressKanji1"`
	AddressKanji2            string                      `json:"addressKanji2"`
	AddressKana1             string                      `json:"addressKana1"`
	AddressKana2             string                      `json:"addressKana2"`
	CellPhoneNumber          string                      `json:"cellPhoneNumber"`
	MailAddress              string                      `json:"mailAddress"`
	PartnerCustnum           string                      `json:"partnerCustnum"`
	TisCntnum                string                      `json:"tisCntnum"`
	IssueDate                int64                       `json:"issueDate"`
	CardStatus01             string                      `json:"status01"`
	CardStatus02             string                      `json:"status02"`
	CardStatus03             string                      `json:"status03"`
	CardStatus04             string                      `json:"status04"`
	CardStatus05             string                      `json:"status05"`
	CardStatus06             string                      `json:"status06"`
	CardStatus07             string                      `json:"status07"`
	CardStatus08             string                      `json:"status08"`
	CardStatus09             string                      `json:"status09"`
	CardStatus10             string                      `json:"status10"`
	CardStatus11             string                      `json:"status11"`
	CardStatus12             string                      `json:"status12"`
	CardStatus13             string                      `json:"status13"`
	CardStatus14             string                      `json:"status14"`
	CardStatus15             string                      `json:"status15"`
	ShoppingOnce             float64                     `json:"shoppingOnce"`
	ShoppingOneDay           float64                     `json:"shoppingDay"`
	ShoppingOneMonth         float64                     `json:"shoppingMonth"`
	IntCashOne               float64                     `json:"intCashOnce"`
	IntCashDay               float64                     `json:"intCashDay"`
	IntCashMonth             float64                     `json:"intCashMonth"`
	IntShoppingOnce          float64                     `json:"intShoppingOnce"`
	IntShoppingOnceDay       float64                     `json:"intShoppingDay"`
	IntShoppingOnceMonth     float64                     `json:"intShoppingMonth"`
	CardIssuanceHistoryBasic []*CardIssuanceHistoryBasic `json:"cardIssueList"`
	CreatedAt                *time.Time                  `json:"createdAt"`
	UpdatedAt                *time.Time                  `json:"updatedAt"`
	CreatedPgmID             string                      `json:"createdPgmID"`
	UpdatedPgmID             string                      `json:"updatedPgmID"`
}

// CardIssuanceHistoryBasic ...
type CardIssuanceHistoryBasic struct {
	IssuanceApplyDate time.Time `json:"issuanceDate"`
	FirstIssueFlag    int8      `json:"firstIssueFlag"`
	IssuanceFlag      int8      `json:"reissueFlag"`
	ActivatedFlag     int8      `json:"activatedFlag"`
	ActiveDate        time.Time `json:"activationDate"`
}

// B3PhysicalCard ...
type B3PhysicalCard struct {
	DesignID       string `json:"design_id"`
	DesignImageURL string `json:"image_url"`
	DesignName     string `json:"design_name"`
	ClubID         string `json:"club_id"`
	IssuanceCost   string `json:"issuance_cost"`
	ValidFlg       string `json:"valid_flag"`
	CreatedAt      string `json:"created_at"`
	Creater        string `json:"creater"`
	UpdatedAt      string `json:"updated_at"`
	Updater        string `json:"updater"`
	CreatedPgmID   string `json:"created_pgm_id"`
	UpdatedPgmID   string `json:"updated_pgm_id"`
}

// LqdKycHistory ...
type LqdKycHistory struct {
	ApplicantID           string             `json:"applicantID"`
	ApplicantNumber       int                `json:"applicantNumber"`
	ApplicantConfirmedAt  string             `json:"applicantConfirmedAt"`
	OperationCompletedAt  string             `json:"operationCompletedAt"`
	OperatorName          string             `json:"operatorName"`
	OperatorLoginID       string             `json:"operatorLoginID"`
	FaceVerification      string             `json:"faceVerification"`
	FacePhotoVerification string             `json:"facePhotoVerification"`
	KycResultCode         string             `json:"kycResultCode"`
	DeficiencyReasonList  []DeficiencyReason `json:"deficiencyReasonList"`
	CreatedAt             *time.Time         `json:"createdAt"`
	UpdatedAt             *time.Time         `json:"updatedAt"`
	CreatedPgmID          string             `json:"createdPgmID"`
	UpdatedPgmID          string             `json:"updatedPgmID"`
}

// DeficiencyReason ...
type DeficiencyReason struct {
	SerialNumber         int    `json:"serialNumber"`
	FaultReasonCode      string `json:"faultReasonCode"`
	FaultReasonText      string `json:"faultReasonText"`
	UserNotificationCode string `json:"userNotificationCode"`
	UserMessageTitle     string `json:"userMessageTitle"`
	UserMessageContent   string `json:"userMessageContent"`
}

// Result ...
type Result struct {
	MaxPage                 int                        `json:"max_page"`
	MemberList              []*MemberInfo              `json:"member_list"`
	MemberID                string                     `json:"member_id"`
	MemberName              string                     `json:"member_name"`
	MemberKana              string                     `json:"member_kana"`
	MemberRoman             string                     `json:"member_roman"`
	CICResultStatus         string                     `json:"CIC_result_status"`
	UsageStatus             string                     `json:"usage_status"`
	Birthday                string                     `json:"birthday"`
	PhoneNumber             string                     `json:"phone_number"`
	StreetAddress           string                     `json:"street_addr"`
	DeliveryDestination     string                     `json:"delivery_destination"`
	CardList                CardInfo                   `json:"card_list"`
	Notices                 string                     `json:"notices"`
	ConfirmationHistoryList []*ConfirmationHistoryInfo `json:"confirmation_history_list"`
}

// MemberInfo ...
type MemberInfo struct {
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

// CardInfo ...
type CardInfo struct {
	Card           string   `json:"card"`
	IssueDate      int64    `json:"issue_date"`
	CardDesign     string   `json:"card_design"`
	ContractNumber string   `json:"contract_number"`
	CardStatus     []string `json:"card_status"`
}

// END BFF CardGetMemberList

// ClubRewardInfo ...
type ClubRewardInfo struct {
	ClubID          uint64             `json:"clubID"`
	ClubName        string             `json:"clubName"`
	ClubOwnerID     uint64             `json:"clubOwnerID"`
	ClubOwnerName   string             `json:"clubOwnerName"`
	AchievementList []*AchievementInfo `json:"achievementList"`
}

// START: CardGetHistoryBenefit ...
// CardHistoryBenefitResult ...
type CardHistoryBenefitResult struct {
	ClubName                 string                      `json:"club_name"`
	ClubOwnerName            string                      `json:"club_owner_name"`
	ClubUsageAmount          int                         `json:"club_usage_amount"`
	ClubUsagePeriodList      []string                    `json:"club_usage_period_list"`
	RewardAchievementHistory []*RewardAchievementHistory `json:"reward_achievement_history_list"`
}

// RewardAchievementHistory ...
type RewardAchievementHistory struct {
	AchievementRewardName string `json:"achievement_reward_name"`
	AchievementDate       string `json:"achievement_date"`
}

// CardHistoryBenefit ...
type CardHistoryBenefit struct {
	ClubRewardList []*ClubRewardInfo `json:"clubRewardList"`
}

// AchievementInfo ...
type AchievementInfo struct {
	UserID            string `json:"userId"`
	CardID            string `json:"cardId"`
	ClubID            uint64 `json:"clubId"`
	RewardName        string `json:"rewardName"`
	RewardID          uint64 `json:"rewardID"`
	AchievementDate   uint64 `json:"achievementDate"`
	RewardImageURL    string `json:"rewardImageUrl"`
	ThumbnailImageURL string `json:"thumbnailImageUrl"`
	RewardDownloadURL string `json:"rewardDownloadUrl"`
	RewardType        string `json:"rewardType"`
	AchieveAmount     uint64 `json:"achieveAmount"`
}

// ClubAmount ...
type ClubAmount struct {
	AmountList []*Amount `json:"amountList"`
}

type Amount struct {
	CardID             string `json:"cardID"`
	UserID             string `json:"userID"`
	ClubID             uint64 `json:"clubID"`
	ClubName           string `json:"clubName"`
	CurrentUsageAmount int    `json:"currentUsageAmount"`
	NoticeID           string `json:"noticeID"`
	NoticeType         string `json:"noticeType"`
	SendingDateTime    string `json:"sendingDatetime"`
	Subject            string `json:"subject"`
	NoticeText         string `json:"noticeText"`
}

// ClubInfo ...
type ClubInfoHistoryNotification struct {
	ClubID           uint64                    `json:"clubID"`
	ClubName         string                    `json:"clubName"`
	ClubOwnerID      uint64                    `json:"clubOwnerID"`
	ClubOwnerName    string                    `json:"clubOwnerName"`
	Description      string                    `json:"description"`
	ClubImageURL     string                    `json:"clubImageUrl"`
	CoverImageURL    string                    `json:"coverImageUrl"`
	ThemeColor       string                    `json:"themeColor"`
	CategoryCode     uint64                    `json:"categoryCode"`
	ReleaseStatus    string                    `json:"releaseStatus"`
	ReleaseStartDate int64                     `json:"releaseStartDate"`
	ReleaseEndDate   int64                     `json:"releaseEndDate"`
	MaxMembers       uint64                    `json:"maxMembers"`
	ClubParticipants uint64                    `json:"clubParticipants"`
	MaxMembershipNo  uint64                    `json:"maxMembershipNo"`
	Rewards          []RewardResponseInquiryDB `json:"rewards"`
	CreatedAt        string                    `json:"createdAt"`
	UpdatedAt        string                    `json:"updatedAt"`
	CreatedPgmID     string                    `json:"createdPgmID"`
	UpdatedPgmID     string                    `json:"updatedPgmID"`
}

// RewardResponseInquiryDB ...
type RewardResponseInquiryDB struct {
	RewardID                   uint64 `json:"rewardID"`
	RewardName                 string `json:"rewardName"`
	ClubID                     uint64 `json:"clubID"`
	Description                string `json:"description"`
	RewardImageURL             string `json:"rewardImageUrl"`
	ThumbnailImageURL          string `json:"thumbnailImageUrl"`
	RewardDownloadURL          string `json:"rewardDownloadUrl"`
	RewardType                 string `json:"rewardType"`
	MadeBy                     string `json:"madeBy"`
	ShipmentSource             string `json:"shipmentSource"`
	AchieveAmount              int64  `json:"achieveAmount"`
	DistributionStatus         string `json:"distributionStatus"`
	DistributionStartDate      int64  `json:"distributionStartDate"`
	DistributionEndDate        int64  `json:"distributionEndDate"`
	ReleaseStatus              string `json:"releaseStatus"`
	TotalOrder                 int64  `json:"totalOrder"`
	LimitEditionFlag           string `json:"limitEditionFlag"`
	NumberPeopleAchievedReward int64  `json:"numberPeopleAchievedReward"`
	CreatedAt                  string `json:"createdAt"`
	UpdatedAt                  string `json:"updatedAt"`
	CreatedPgmID               string `json:"createdPgmID"`
	UpdatedPgmID               string `json:"updatedPgmID"`
}

// ClubCard ...
type ClubCard struct {
	ClubList []*ClubInfoHistoryBenefit `json:"clubList"`
}

type ClubInfoHistoryBenefit struct {
	CardID             string `json:"cardID"`
	UserID             string `json:"userID"`
	ClubID             uint64 `json:"clubID"`
	ClubName           string `json:"clubName"`
	ClubOwnerID        int    `json:"clubOwnerID"`
	ClubOwnerName      string `json:"clubOwnerName"`
	ClubOwnerImageUrl  string `json:"clubOwnerImageUrl"`
	ClubImageUrl       string `json:"clubImageUrl"`
	CoverImageUrl      string `json:"coverImageUrl"`
	ThemeColor         string `json:"themeColor"`
	MemberNo           int    `json:"memberNo"`
	ClubEnrollmentDate uint64 `json:"clubEnrollmentDate"`
	ClubWithdrawalDate uint64 `json:"clubWithdrawalDate"`
}

// ClubUsage ...
type ClubUsage struct {
	ClubUsageAmount     int      `json:"club_usage_amount"`
	ClubUsagePeriodList []string `json:"club_usage_period_list"`
}

// END CardGetHistoryBenefit

// START CardGetHistoryRepayment

// CardHistoryRepaymentResponse ...
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

// END CardGetHistoryRepayment

// START CardGetHistoryUsage
// CardHistoryUsage ...
type CardHistoryUsage struct {
	CardID              string `json:"card_id"`
	UserID              string `json:"user_id"`
	TransactionsOfMonth string `json:"tran_month"`
}

// PaymentFixTradeResponse ...
type PaymentFixTradeResponse struct {
	Month        string       `json:"month"`
	FixTradeList []*TradeInfo `json:"fixTradeList"`
}

// TradeInfo ...
type TradeInfo struct {
	UserID                  string  `json:"userID"`
	TransactionInfo         string  `json:"transactionInfo"`
	TisCntNum               string  `json:"tisCntNum"`
	UseDate                 string  `json:"useDate"`
	UseMonth                string  `json:"useMonth"`
	AuthorizationDate       string  `json:"authorizationDate"`
	FixDate                 string  `json:"fixDate"`
	FixMonth                string  `json:"fixMonth"`
	FormatType              string  `json:"formatType"`
	UsageType               string  `json:"usageType"`
	UsageTypeName           string  `json:"usageTypeName"`
	MerchantCategoryCode    string  `json:"merchantCategoryCode"`
	MerchantCategoryName    string  `json:"merchantCategoryName"`
	MerchantName            string  `json:"merchantName"`
	AuthorizationAmountSign int     `json:"authorizationAmountSign"`
	AuthorizationAmount     float32 `json:"authorizationAmount"`
	FixedAmountSign         int     `json:"fixedAmountSign"`
	FixedAmount             float32 `json:"fixedAmount"`
	ApprovalNumber          string  `json:"approvalNumber"`
	CalcFlg                 bool    `json:"calcFlg"`
	FileDateTime            string  `json:"fileDateTime"`
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

// END CardGetHistoryUsage
