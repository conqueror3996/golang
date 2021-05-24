package model

import "time"

// NotificationList ...
type NotificationList struct {
	NoticeType string `json:"noticeType"`
	UserID     string `json:"userID"`
	ClubName   string `json:"clubName"`
	Subject    string `json:"subject"`
	NoticeText string `json:"noticeText"`
	SortKey    string `json:"sortKey"`
	Order      string `json:"order"`
	Page       string `json:"page"`
	Size       string `json:"size"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// NotificationListResult ...
type NotificationListResult struct {
	MaxPage     int                    `json:"max_page"`
	MessageList []*MessageInfoResponse `json:"message_list"`
}

// ResponseNotificationList ...
type MessageInfo struct {
	NoticeID        string    `json:"notice_id"`
	ClubID          string    `json:"club_id"`
	UserID          string    `json:"User_id"`
	NoticeType      string    `json:"notice_type"`
	Subject         string    `json:"subject"`
	NoticeText      string    `json:"notice_text"`
	SendingDatetime time.Time `json:"sending_datetime"`
}

type MessageInfoResponse struct {
	NoticeID        string    `json:"notice_id"`
	ClubName        string    `json:"club_name"`
	UserID          string    `json:"User_id"`
	NoticeType      string    `json:"notice_type"`
	Subject         string    `json:"subject"`
	NoticeText      string    `json:"notice_text"`
	SendingDatetime time.Time `json:"sending_datetime"`
}

// END NotificationList

// NotificationRegistration ...
type NotificationRegistration struct {
	NoticeType string `json:"noticeType"`
	UserID     string `json:"userId"`
	ClubID     string `json:"clubId"`
	Subject    string `json:"subject"`
	NoticeText string `json:"noticeText"`
}

// ResponseNotificationRegistration ...
type ResponseNotificationRegistration struct {
	NoticeID        string `json:"noticeID"`
	NoticeType      string `json:"noticeType"`
	Subject         string `json:"subject"`
	NoticeText      string `json:"noticeText"`
	SendingDatetime string `json:"sendingDatetime"`
	UserID          string `json:"userID"`
	ClubID          string `json:"clubID"`
	DeleteFlag      string `json:"deleteFlag"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedPgmID    string `json:"createdPgmID"`
	UpdatedPgmID    string `json:"updatedPgmID"`
}
