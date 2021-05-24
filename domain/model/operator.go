package model

// AuthLogin AuthLogin's Out Entity
// type AuthLogin struct {
// 	AccessToken  string
// 	RefreshToken string
// 	ExpiresDate  int
// 	OperatorInfo OperatorInfo
// }

// CommonRequestMeta ...
type CommonRequestMeta struct {
	HeaderContentType           string
	HeaderUserAgent             string
	HeaderXNudgeInteractionID   string
	HeaderXNudgeUserID          string
	HeaderXNudgeSecurityRequest string
}

// OperatorInfo ...
type OperatorInfo struct {
	OperatorID     string           `json:"operator_id"`
	OperatorRoleID string           `json:"operator_role_id"`
	Username       string           `json:"user_name"`
	MailAddress    string           `json:"mail_address"`
	PhoneNumber    string           `json:"phone_number"`
	Belong         string           `json:"belong"`
	AccountStatus  string           `json:"account_status"`
	FirstLoginFlg  bool             `json:"first_login_flg"`
	RevokeCount    int              `json:"revoke_count"`
	CreateDate     string           `json:"create_date"`
	CreatePgmID    string           `json:"create_pgm_id"`
	UpdateDate     string           `json:"update_date"`
	UpdatePgmID    string           `json:"update_pgm_id"`
	Role           OperatorRoleInfo `json:"role"`
}

// OperatorResponeInfo ...
type OperatorResponeInfoPattern2 struct {
	Username      string           `json:"user_name"`
	MailAddress   string           `json:"mail_address"`
	PhoneNumber   string           `json:"phone_number"`
	Belong        string           `json:"belong"`
	AccountStatus string           `json:"account_status"`
	Role          OperatorRoleInfo `json:"role"`
}

// OperatorRoleInfo ...
type OperatorRoleInfo struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// TokenInfo ...
type TokenInfo struct {
	AccessToken  string
	RefreshToken string
	ExpiresDate  int
}

// AuthLoginResponse ...
type AuthLoginResponse struct {
	LoginStatus int `json:"login_status"`
}

// AuthLoginSMSResponse ...
type AuthLoginSMSResponse struct {
	LoginStatus int          `json:"login_status"`
	UserName    string       `json:"user_name"`
	Role        RoleResponse `json:"role"`
}

// AuthChangePasswordResponse ...
type AuthChangePasswordResponse struct {
	LoginStatus int                `json:"login_status"`
	UserName    string             `json:"user_name"`
	Role        RoleChangeResponse `json:"role"`
}

// RoleResponse ...
type RoleResponse struct {
	CardRole     string `json:"card_role"`
	PartnerRole  string `json:"partner_role"` // Club owner
	SkinRole     string `json:"skin_role"`    // Club
	OperatorRole string `json:"operator_role"`
	MessageRole  string `json:"message_role"`
}

// RoleChangeResponse ...
type RoleChangeResponse struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// Session ...
type Session struct {
	SessionID   string `json:"session_id"`
	ID          string `json:"id"`
	OperatorID  string `json:"operator_id"`
	AuthStatus  string `json:"auth_status"`
	UserStatus  string `json:"user_status"`
	FirstLogin  bool   `json:"first_login"`
	FirstName   string `json:"first_name"`
	FamilyName  string `json:"family_name"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Expired     int64  `json:"expires_on"`
	Flag        string `json:"login_status"`
	CreateDate  string `json:"created_at"`
	UpdatedDate string `json:"updated_at"`
	CreatePgmID string `json:"created_pgm_id"`
	UpdatePgmID string `json:"updated_pgm_id"`
}

// SessionUpdate ...
type SessionUpdate struct {
	AuthStatus    string `json:":auth_status"`
	Expired       int64  `json:":expires_on"`
	FirstLoginFlg bool   `json:":first_login"`
	UpdatedDate   string `json:":updated_at"`
	UpdatePgmID   string `json:":updated_pgm_id"`
}

// UpdateSessionCondition ...
type UpdateSessionCondition struct {
	SessionID string `json:"session_id"`
}

// ResetKey ...
type ResetKey struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Flag        string `json:"reset_status"`
	Expired     int64  `json:"expires_on"`
	CreateDate  string `json:"created_at"`
	UpdatedDate string `json:"updated_at"`
	CreatePgmID string `json:"created_pgm_id"`
	UpdatePgmID string `json:"updated_pgm_id"`
}

// ResetKeyUpdate
type ResetKeyUpdate struct {
	Flag        string `json:":reset_status"`
	UpdatedDate string `json:":updated_at"`
	UpdatePgmID string `json:":updated_pgm_id"`
}

// ResetKeyCondition ...
type ResetKeyCondition struct {
	Key string `json:"key"`
}

// OperatorListSearchResponse ...
type OperatorListSearchResponse struct {
	Result OperatorResult `json:"result"`
}

// OperatorListSearchResponse2
type OperatorListSearchResponsePattern2 struct {
	Result OperatorResponeInfoPattern2 `json:"result"`
}

// OperatorResult ...
type OperatorResult struct {
	MaxPage      int                       `json:"max_page"`
	OperatorList []*OperatorSearchResponse `json:"operator_list"`
}

// OperatorSearchResponse ...
type OperatorSearchResponse struct {
	OperatorID     string `json:"operator_id"`
	OperatorRoleID string `json:"operator_role_id"`
	Username       string `json:"user_name"`
	MailAddress    string `json:"mail_addr"`
	PhoneNumber    string `json:"phone_number"`
	Belong         string `json:"belong"`
	AccountStatus  string `json:"account_status"`
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// OperatorUpdateResponse ...
type OperatorUpdateResponse struct {
	LoginStatus int `json:"login_status"`
}

// OperatorDelete ...
type OperatorDelete struct {
	OperatorID    string `json:"operator_id"`
	AccountStatus string `json:"account_status"`
}

// OperatorDeleteResponse ...
type OperatorDeleteResponse struct {
	LoginStatus int `json:"login_status"`
}

// OperatorRegistration ...
type OperatorRegistration struct {
	UserName      string                   `json:"user_name"`
	BeLong        string                   `json:"belong"`
	MailAddress   string                   `json:"mail_address"`
	PhoneNumber   string                   `json:"phone_number"`
	Password      string                   `json:"password"`
	Role          OperatorRoleRegistration `json:"role"`
	FirstLoginFlg bool                     `json:"first_login_flg"`
}

// OperatorRoleRegistration ...
type OperatorRoleRegistration struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// ResponseOperatorRegistration ...
type ResponseOperatorRegistration struct {
	OperatorID string `json:"operatorID"`
}

// MetaRequest Input Entity
type MetaRequest struct {
	ReceptionProcessingStatus string
	HeaderContentType         string
	HeaderUserAgent           string
	HeaderXRequestedWith      string
	HeaderXNudgeInteractionID string
	HeaderXAPIKey             string
	HeaderAuthorization       string
	Body                      string
}
