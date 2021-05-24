package input

// OwnerRegistration ...
type OwnerRegistration struct {
	ClubOwnerName string `json:"club_owner_name"`
	MailAddress   string `json:"mail_address"`
	PhoneNumber   string `json:"phone_number"`
	CompanyName   string `json:"company_name"`
	InChargeName  string `json:"in_charge_name"`
}

// OwnerList ...
type OwnerList struct {
	Type        string `json:"type"`
	SortKey     string `json:"sort_key"`
	SortOrder   string `json:"sort_order"`
	PageNum     string `json:"page_num"`
	PageSize    string `json:"page_size"`
	SearchText  string `json:"search_text"`
	ClubOwnerID string `json:"club_owner_id"`
}

// OwnerListRequestPattern1 ...
type OwnerListRequestPattern1 struct {
	SearchType  string `json:"searchType"`
	Size        string `json:"size"`
	Page        string `json:"page"`
	SortKey     string `json:"sortKey"`
	Order       string `json:"order"`
	SearchValue string `json:"searchValue"`
}

// OwnerUpdate
type OwnerUpdate struct {
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
