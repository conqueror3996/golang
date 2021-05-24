package config

import (
	"encoding/json"
	"fmt"
)

var internalSevenBankErrorCodeMapConfig map[string]string

func init() {
	// fmt.Println("config/internal_seven_bank_error_code_map.go : init : func init() start.")

	err := json.Unmarshal([]byte(InternalSevenBankErrorCodesMapConfig), &internalSevenBankErrorCodeMapConfig)
	if err != nil {
		panic(err)
	}

	// fmt.Println("config/internal_seven_bank_error_code_map.go : init : func init() end.")
}

// GetInternal2SevenBankErrorCode ...
func GetInternal2SevenBankErrorCode(key string) string {
	return internalSevenBankErrorCodeMapConfig[key]
}

// GetInternal2SevenBankErrorCodeV1 ...
func GetInternal2SevenBankErrorCodeV1(apiID string, internalStatusCode int, internalErrorCode string) string {
	sevenBankErrorCode := ""

	key := fmt.Sprintf("%s.%d.%s", apiID, internalStatusCode, internalErrorCode)
	sevenBankErrorCode = internalSevenBankErrorCodeMapConfig[key]
	if sevenBankErrorCode != "" {
		return sevenBankErrorCode
	}

	key = fmt.Sprintf("%s..%s", apiID, internalErrorCode)
	sevenBankErrorCode = internalSevenBankErrorCodeMapConfig[key]
	if sevenBankErrorCode != "" {
		return sevenBankErrorCode
	}

	key = fmt.Sprintf("%s.%d.", apiID, internalStatusCode)
	sevenBankErrorCode = internalSevenBankErrorCodeMapConfig[key]
	if sevenBankErrorCode != "" {
		return sevenBankErrorCode
	}

	key = fmt.Sprintf("%s..", apiID)
	sevenBankErrorCode = internalSevenBankErrorCodeMapConfig[key]
	if sevenBankErrorCode != "" {
		return sevenBankErrorCode
	}

	return SevenBankErrorCode6305
}

// GetInternal2SevenBankErrorV1 ...
func GetInternal2SevenBankErrorV1(apiID string, internalStatusCode int, internalErrorCode string) (int, string) {
	configSevenBankErrorCode := GetInternal2SevenBankErrorCodeV1(apiID, internalStatusCode, internalErrorCode)

	sevenBankStatusCode := GetSevenBankStatusCode(configSevenBankErrorCode)
	sevenBankErrorCode := GetSevenBankErrorCode(configSevenBankErrorCode)

	return sevenBankStatusCode, sevenBankErrorCode
}
