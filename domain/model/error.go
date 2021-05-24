package model

// ErrorResult Out Entity For Error Result
type ErrorResult struct {
	Error Error `json:"error"`
}

// Error Out Entity For Error Result
type Error struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Action  string `json:"action"`
}
