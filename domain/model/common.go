package model

// ConfigurationA ...
type ConfigurationA struct {
	ConfigType int
	BasePath   string
	Host       string
	Headers    map[string]string
	UserAgent  string
}
