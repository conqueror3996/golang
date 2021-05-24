package input

// AuthLogin Input Entity
type AuthLogin struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	SmsCode  string `json:"sms_code"`
}

// AuthLoginMeta Input Entity
type AuthLoginMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderAuthorization       string
	HeaderXRequestedWith      string
	Body                      string
}

// AuthReset Input Entity
type AuthReset struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}

// AuthChangePasswordRequestToSVC ...
type AuthChangePasswordRequestToSVC struct {
	MailAddress string `json:"id"`
	Password    string `json:"password"`
}

// AuthResetMeta Input Entity
type AuthResetMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderXRequestedWith      string
	Body                      string
}

// AuthChangePassword ...
type AuthChangePassword struct {
	ID       string `json:"mail_address"`
	Password string `json:"password"`
}

// AuthChangePasswordMeta ...
type AuthChangePasswordMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderAuthorization       string
	HeaderXRequestedWith      string
	Body                      string
}

// ResetRequest ...
type ResetRequest struct {
	MailAddr string `json:"mail_addr"`
}

// ResetRequestToSVC ...
type ResetRequestToSVC struct {
	MailAddress string `json:"mail_addr"`
	URL         string `json:"url"`
}

// ResetRequestMeta ...
type ResetRequestMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXRequestedWith      string
	HeaderXNudgeInteractionID string
	Body                      string
}

// OperatorList ...
type OperatorList struct {
	SearchText string `json:"search_text"`
	SortKey    string `json:"sort_key"`
	SortOrder  string `json:"sort_order"`
	PageNum    string `json:"page_num"`
	PageSize   string `json:"page_size"`
	OperatorID string `json:"operator_id"`
}

// OperatorListMeta ...
type OperatorListMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXRequestedWith      string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderNWSessionID         string
	Body                      string
}

// OperatorListRequestToSVCPattern1 ...
type OperatorListRequestToSVCPattern1 struct {
	OperatorID string `json:"operator_id"`
	Role       Role   `json:"role"`
	// Role       string `json:"role"`
}

// OperatorListRequestToSVCPattern2 ...
type OperatorListRequestToSVCPattern2 struct {
	SearchText string `json:"search_text"`
	SortKey    string `json:"sort_key"`
	SortOrder  string `json:"sort_order"`
	PageNum    string `json:"page_num"`
	PageSize   string `json:"page_size"`
	Role       Role   `json:"role"`
}

// RequestToSVCPattern2 ...
type RequestToSVCPattern2 struct {
	SearchText   string `json:"search_text"`
	SortKey      string `json:"sort_key"`
	SortOrder    string `json:"sort_order"`
	PageNum      string `json:"page_num"`
	PageSize     string `json:"page_size"`
	OperatorRole string `json:"operator_role"`
}

// OperatorUpdate ...
type OperatorUpdate struct {
	OperatorID    string `json:"operator_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Belong        string `json:"belong"`
	MailAddress   string `json:"mail"`
	PhoneNumber   string `json:"tel"`
	Password      string `json:"password"`
	Role          Role   `json:"role"`
	FirstLoginFlg *bool  `json:"first_login_flg"`
}

// OperatorUpdateRequestToSVC ...
type OperatorUpdateRequestToSVC struct {
	OperatorID    string `json:"operator_id"`
	AccountStatus string `json:"account_status"`
	UserName      string `json:"user_name"`
	Belong        string `json:"belong"`
	MailAddress   string `json:"mail_address"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `json:"password"`
	Role          Role   `json:"role"`
	FirstLoginFlg bool   `json:"first_login_flg"`
}

// Role ...
type Role struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// OperatorUpdateMeta ...
type OperatorUpdateMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderXRequestedWith      string
	Body                      string
}

// OperatorDelete ...
type OperatorDelete struct {
	OperatorID    string `json:"operator_id"`
	AccountStatus string `json:"operator_status"`
}

// OperatorDeleteMeta ...
type OperatorDeleteMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderXRequestedWith      string
	Body                      string
}

// Operator Input Entity
type Operator struct {
	OperatorID     string
	OperatorRoleID string
	Username       string
	Password       string
	MailAddress    string
	PhoneNumber    string
	Belong         string
	AccountStatus  string
	FirstLoginFlg  bool
	RevokeCount    int
	CreateDate     string
	CreatePgmID    string
	UpdateDate     string
	UpdatePgmID    string
}

// GetOperator Input Entity
type GetOperator struct {
	OperatorID            string
	OperatorRoleID        string
	Username              string
	MailAddress           string
	PhoneNumber           string
	Belong                string
	AccountStatus         string
	FirstLoginFlg         bool
	RevokeCount           int
	CreateDate            string
	CreatePgmID           string
	UpdateDate            string
	UpdatePgmID           string
	Page                  string
	MaximumRecordsPerPage string
	KeyWords              string
}

// OperatorInquiry Input Entity
type OperatorInquiry struct {
	OperatorID            string
	Page                  string
	MaximumRecordsPerPage string
	KeyWords              string
}

// OperatorInquiryMeta Input Entity
type OperatorInquiryMeta struct {
	ReceptionProcessingStatus   string
	HeaderContentType           string
	HeaderHost                  string
	HeaderUserAgent             string
	HeaderXNudgeInteractionID   string
	HeaderXNudgeUserID          string
	HeaderXNudgeSecurityRequest string
	HeaderAuthorization         string
	Body                        string
}

// OperatorOperatorIDInquiry Input Entity
type OperatorOperatorIDInquiry struct {
	OperatorID     string
	OperatorRoleID string
	Username       string
	MailAddress    string
	PhoneNumber    string
	Belong         string
	AccountStatus  string
	FirstLoginFlg  bool
	RevokeCount    int
	CreateDate     string
	CreatePgmID    string
	UpdateDate     string
	UpdatePgmID    string
}

// OperatorOperatorIDInquiryMeta Input Entity
type OperatorOperatorIDInquiryMeta struct {
	ReceptionProcessingStatus   string
	HeaderContentType           string
	HeaderHost                  string
	HeaderUserAgent             string
	HeaderXNudgeInteractionID   string
	HeaderXNudgeUserID          string
	HeaderXNudgeSecurityRequest string
	HeaderAuthorization         string
	Body                        string
	PathOperatorID              string
}

// OperatorRegistration Input Entity
type OperatorRegistration struct {
	LastName      string                   `json:"last_name"`
	FirstName     string                   `json:"first_name"`
	BeLong        string                   `json:"belong"`
	MailAddress   string                   `json:"mail"`
	PhoneNumber   string                   `json:"phone_number"`
	Password      string                   `json:"password"`
	Role          OperatorRoleRegistration `json:"role"`
	FirstLoginFlg *bool                    `json:"first_login_flg"`
}

// OperatorRoleRegistration ...
type OperatorRoleRegistration struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// OperatorRegistrationMeta Input Entity
type OperatorRegistrationMeta struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderAuthorization       string
	HeaderXRequestedWith      string
	Body                      string
}

// // MetaRequest Input Entity
// type MetaRequest struct {
// 	ReceptionProcessingStatus string
// 	HeaderContentType         string
// 	HeaderUserAgent           string
// 	HeaderXRequestedWith      string
// 	HeaderXNudgeInteractionID string
// 	HeaderXAPIKey             string
// 	HeaderAuthorization       string
// 	Body                      string
// }
