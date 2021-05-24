package json

import "time"

type NotificationListResponse struct {
	Result NotificationListResult `json:"result"`
}

// NotificationListResult ...
type NotificationListResult struct {
	MaxPage     int                    `json:"max_page"`
	MessageList []*MessageInfoResponse `json:"message_list"`
}

// MessageInfoResponse ...
type MessageInfoResponse struct {
	NoticeID        string    `json:"notice_id"`
	ClubName        string    `json:"club_name"`
	UserID          string    `json:"User_id"`
	NoticeType      string    `json:"notice_type"`
	Subject         string    `json:"subject"`
	NoticeText      string    `json:"notice_text"`
	SendingDatetime time.Time `json:"sending_datetime"`
}
