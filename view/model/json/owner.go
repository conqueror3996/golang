package json

// ResponseOwnerNotification ...
type ResponseOwnerNotification struct {
	ClubOwnerID   int    `json:"club_owner_id"`
	ReleaseStatus string `json:"release_status"`
}

// ClubOwnerUploadURLResponse ...
type ClubOwnerUploadURLResponse struct {
	UploadURL string `json:"upload_url"`
}
