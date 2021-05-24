package json

// AuthLogin AuthLogin's Out Entity
type AuthLogin struct {
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
	Role           OperatorRoleInfo
}

// OperatorRoleInfo ...
type OperatorRoleInfo struct {
	CardMemberRole string
	ClubRole       string
	ClubOwnerRole  string
	MessageRole    string
	OperatorRole   string
}

// OperatorRoleInfoResponse ...
type OperatorRoleInfoResponse struct {
	CardMemberRole string `json:"card_member_role"`
	ClubRole       string `json:"club_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	MessageRole    string `json:"message_role"`
	OperatorRole   string `json:"operator_role"`
}

// ResetRequestResponse ...
type ResetRequestResponse struct {
}

// ResetResponse ...
type ResetResponse struct {
}

// AuthChangePasswordResponse ...
type AuthChangePasswordResponse struct {
	LoginStatus int                `json:"login_status"`
	UserName    string             `json:"user_name"`
	Role        RoleChangeResponse `json:"role"`
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

// RoleResponse ...
type RoleResponse struct {
	CardRole     string `json:"card_role"`
	PartNerRole  string `json:"partner_role"`
	SkinRole     string `json:"skin_role"`
	OperatorRole string `json:"operator_role"`
	MessageRole  string `json:"message_role"`
}

// RoleResponse ...
type RoleChangeResponse struct {
	CardMemberRole string `json:"card_member_role"`
	ClubOwnerRole  string `json:"club_owner_role"`
	ClubRole       string `json:"club_role"`
	OperatorRole   string `json:"operator_role"`
	MessageRole    string `json:"message_role"`
}

// OperatorListSearchResponse ...
type OperatorListSearchResponse struct {
	Result OperatorResult `json:"result"`
}

type OperatorListSearchResponsePattern2 struct {
	Result OperatorResponeInfoPattern2 `json:"result"`
}

// OperatorResponeInfo ...
type OperatorResponeInfoPattern2 struct {
	Name          string                   `json:"name"`
	MailAddress   string                   `json:"mail_address"`
	PhoneNumber   string                   `json:"phone_number"`
	Belong        string                   `json:"belong"`
	AccountStatus string                   `json:"account_status"`
	Role          OperatorRoleInfoResponse `json:"role"`
}

// OperatorResult ...
type OperatorResult struct {
	MaxPage      int                       `json:"max_page"`
	OperatorList []*OperatorSearchResponse `json:"operator_list"`
}

// OperatorSearchResponse ...
type OperatorSearchResponse struct {
	OperatorID     string                   `json:"operator_id"`
	OperatorRoleID string                   `json:"operator_role_id"`
	Username       string                   `json:"user_name"`
	MailAddress    string                   `json:"mail_addr"`
	PhoneNumber    string                   `json:"phone_number"`
	Belong         string                   `json:"belong"`
	Role           OperatorRoleInfoResponse `json:"role"`
}

// ClubOwnerListRespone ...
type ClubOwnerListRespone struct {
	ClubOwnerID   string `json:"club_owner_id"`
	ClubOwnerName string `json:"club_owner_name"`
	MailAddress   string `json:"mail_address"`
	ImageURL      string `json:"image_url"`
	ReleaseStatus string `json:"release_status"`
}

// ClubOwnerListResult ...
type ClubOwnerListResult struct {
	MaxPage   int                     `json:"max_page"`
	OwnerList []*ClubOwnerListRespone `json:"owner_list"`
}

// ClubOwnerRespone ...
type ClubOwnerRespone struct {
	PhoneNumber   string  `json:"phone_number"`
	CompanyName   string  `json:"compaty_name"`
	InChargeName  string  `json:"in_charge_name"`
	MailAddress   string  `json:"mail_address"`
	ReleaseStatus string  `json:"release_status"`
	ClubList      []*Club `json:"club_list"`
}

// Club ...
type Club struct {
	ClubID       string `json:"club_id"`
	ClubName     string `json:"club_name"`
	ClubImageURL string `json:"club_image_url"`
}

type OwnerListResult struct {
	MaxPage       int                     `json:"max_page"`
	OwnerList     []*ClubOwnerListRespone `json:"owner_list"`
	PhoneNumber   string                  `json:"phone_number"`
	CompanyName   string                  `json:"compaty_name"`
	InChargeName  string                  `json:"in_charge_name"`
	MailAddress   string                  `json:"mail_address"`
	ReleaseStatus string                  `json:"release_status"`
	ClubList      []*Club                 `json:"club_list"`
}

type OwnerListRespone struct {
	Result *OwnerListResult `json:"result"`
}

// OperatorUpdateResponse ...
type OperatorUpdateResponse struct {
}

// OperatorDeleteResponse ...
type OperatorDeleteResponse struct {
}

// RegisterResponse ...
type RegisterResponse struct {
}
