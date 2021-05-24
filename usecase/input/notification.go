package input

// NotificationList ...
type NotificationList struct {
	MessageSection string `json:"message_section"`
	SortKey        string `json:"sort_key"`
	SortOrder      string `json:"sort_order"`
	PageNum        string `json:"page_num"`
	PageSize       string `json:"page_size"`
	MemberID       string `json:"member_id"`
	ClubID         string `json:"club_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}

// NotificationRegistration ...
type NotificationRegistration struct {
	MessageSection string `json:"message_section"`
	MemberID       string `json:"member_id"`
	ClubID         string `json:"club_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
}
