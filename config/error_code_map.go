package config

import (
	"encoding/json"
	"fmt"
)

var errorCodeMapConfig map[string]string

func init() {
	// fmt.Println("config/error_code_map_config.go : init : func init() start.")

	err := json.Unmarshal([]byte(ErrorCodesMapConfig), &errorCodeMapConfig)
	if err != nil {
		panic(err)
	}

	// fmt.Println("config/error_code_map_config.go : init : func init() end.")
}

// GetErrorCode ...
func GetErrorCode(key string) string {
	return errorCodeMapConfig[key]
}

// GetErrorCodeV1 ...
func GetErrorCodeV1(apiID string, statusCode int, errorCode string) string {
	toErrorCode := ""

	key := fmt.Sprintf("%s.%d.%s", apiID, statusCode, errorCode)
	toErrorCode = errorCodeMapConfig[key]
	if toErrorCode != "" {
		return toErrorCode
	}

	key = fmt.Sprintf("%s..%s", apiID, errorCode)
	toErrorCode = errorCodeMapConfig[key]
	if toErrorCode != "" {
		return toErrorCode
	}

	key = fmt.Sprintf("%s.%d.", apiID, statusCode)
	toErrorCode = errorCodeMapConfig[key]
	if toErrorCode != "" {
		return toErrorCode
	}

	key = fmt.Sprintf("%s..", apiID)
	toErrorCode = errorCodeMapConfig[key]
	if toErrorCode != "" {
		return toErrorCode
	}

	return "250000500900705"
}
