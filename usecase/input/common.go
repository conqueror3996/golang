package input

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
