package input

// BFF CardGetMemberList
// CardMemberList
type CardMemberList struct {
	SearchText string `json:"search_text"`
	SortKey    string `json:"sort_key"`
	SortOrder  string `json:"sort_order"`
	PageNum    string `json:"page_num"`
	PageSize   string `json:"page_size"`
	MemberID   string `json:"member_id"`
}

// OperatorListRequestToSVCPattern1 ...
type CardMemberListRequestToSVCPattern1 struct {
	SearchType  string `json:"searchType"`
	SearchValue string `json:"searchValue"`
	SortKey     string `json:"sortKey"`
	Order       string `json:"order"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}

// OperatorListRequestToSVCPattern2 ...
type CardMemberListRequestToSVCPattern2 struct {
	MemberID string `json:"member_id"`
}

// CardHistoryUsage ...
/* NguyenDinhDien - Date : 20210331
type CardHistoryUsage struct {
	CardID              string `json:"card_id"`
	UserID              string `json:"user_id"`
	TransactionsOfMonth string `json:"tran_month"`
}*/
// CardHistoryUsage ...
type CardHistoryUsage struct {
	MemberID  string `json:"member_id"`
	YearMonth string `json:"year_month"`
}

// CardHistoryRepayment ...
type CardHistoryRepayment struct {
	UserID     string `json:"user_id"`
	RepayMonth string `json:"repay_month"`
}

// BFF CardGetHistoryNotification ...
// CardHistoryNotification ...
type CardHistoryNotification struct {
	MemberID    string `json:"member_id"`
	SortKey     string `json:"sort_key"`
	SortOrder   string `json:"sort_order"`
	PageNum     string `json:"page_num"`
	PageSize    string `json:"page_size"`
	DeletedData string `json:"deleted_data"`
}

// CardHistoryNotification ...
type CardHistoryNotificationRequest struct {
	MemberID    string `json:"userId"`
	SortKey     string `json:"sortKey"`
	SortOrder   string `json:"order"`
	PageNum     string `json:"page"`
	PageSize    string `json:"size"`
	DeletedData string `json:"deletedData"`
}

// CardHistoryBenefit
type CardHistoryBenefit struct {
	MemberID string `json:"member_id"`
}
