package model

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

type ResultOwner struct {
	MaxPage       int                     `json:"max_page"`
	OwnerList     []*ClubOwnerListRespone `json:"owner_list"`
	PhoneNumber   string                  `json:"phone_number"`
	CompanyName   string                  `json:"compaty_name"`
	InChargeName  string                  `json:"in_charge_name"`
	MailAddress   string                  `json:"mail_address"`
	ReleaseStatus string                  `json:"release_status"`
	ClubList      []*Club                 `json:"club_list"`
}

// UploadURL
type UploadURL struct {
	Type        int `json:"type"`
	ClubOwnerID int `json:"clubOwnerID"`
}

// OwnerInfo
type OwnerInfo struct {
	ClubOwnerID      string `json:"club_owner_id"`
	ImageURL         string `json:"image_url"`
	MailAddress      string `json:"mail_address"`
	PhoneNumber      string `json:"phone_number"`
	CompanyName      string `json:"company_name"`
	InChargeName     string `json:"in_charge_name"`
	BankName         string `json:"bank_name"`
	BranchCode       string `json:"branch_code"`
	BranchName       string `json:"branch_name"`
	AccountType      string `json:"account_type"`
	AccountNo        string `json:"account_no"`
	AccountName      string `json:"account_name"`
	AccountNameKana  string `json:"account_name_kana"`
	ReleaseStatus    string `json:"release_status"`
	ReleaseStartDate string `json:"release_start_date"`
	ReleaseEndDate   string `json:"release_end_date"`
}

// ClubOwnerRegisterResponse ...
type ClubOwnerRegisterResponse struct {
	ClubOwnerID   uint64 `json:"clubOwnerID"`
	ReleaseStatus int64  `json:"releaseStatus"`
}

// ClubOwnerUploadURLResponse ...
type ClubOwnerUploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}

// OwnerRequestNotification ...
type OwnerRequestNotification struct {
	ClubOwnerName string `json:"clubOwnerName"`
	MailAddress   string `json:"mailAddress"`
	PhoneNumber   string `json:"phoneNumber"`
	CompanyName   string `json:"companyName"`
	InChargeName  string `json:"inChargeName"`
}

// ResponseOwnerNotification ...
type ResponseOwnerNotification struct {
	ClubOwnerID   int    `json:"club_owner_id"`
	ReleaseStatus string `json:"release_status"`
}
