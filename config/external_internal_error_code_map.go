package config

import (
	"encoding/json"
	"fmt"
)

var externalInternalErrorCodesMapConfig map[string]string

func init() {
	// fmt.Println("config/external_internal_error_code_map.go : init : func init() start.")

	err := json.Unmarshal([]byte(ExternalInternalErrorCodesMapConfig), &externalInternalErrorCodesMapConfig)
	if err != nil {
		panic(err)
	}

	// fmt.Println("config/external_internal_error_code_map.go : init : func init() end.")
}

// GetGetExternal2InternalErrorCode ...
func GetGetExternal2InternalErrorCode(key string) string {
	return externalInternalErrorCodesMapConfig[key]
}

// GetExternal2InternalErrorCodeV1 ...
func GetExternal2InternalErrorCodeV1(externalType int, apiID string, externalAPIID string, externalStatusCode int, externalErrorCode string) string {
	internalErrorCode := ""

	key := fmt.Sprintf("%s.%d.%s", externalAPIID, externalStatusCode, externalErrorCode)
	internalErrorCode = externalInternalErrorCodesMapConfig[key]
	if internalErrorCode != "" {
		if internalErrorCode == ExternalInternalErrorCodesMapConfigOrigin {
			if externalType == 0 {
				return externalErrorCode
			}
			return key
		}
		return internalErrorCode
	}

	key = fmt.Sprintf("%s..%s", externalAPIID, externalErrorCode)
	internalErrorCode = externalInternalErrorCodesMapConfig[key]
	if internalErrorCode != "" {
		if internalErrorCode == ExternalInternalErrorCodesMapConfigOrigin {
			if externalType == 0 {
				return externalErrorCode
			}
			return key
		}
		return internalErrorCode
	}

	key = fmt.Sprintf("%s.%d.", externalAPIID, externalStatusCode)
	internalErrorCode = externalInternalErrorCodesMapConfig[key]
	if internalErrorCode != "" {
		if internalErrorCode == ExternalInternalErrorCodesMapConfigOrigin {
			if externalType == 0 {
				if externalErrorCode != "" {
					return externalErrorCode
				}
				return key
			}
			return key
		}
		return internalErrorCode
	}

	key = fmt.Sprintf("%s..", externalAPIID)
	internalErrorCode = externalInternalErrorCodesMapConfig[key]
	if internalErrorCode != "" {
		if internalErrorCode == ExternalInternalErrorCodesMapConfigOrigin {
			if externalType == 0 {
				if externalErrorCode != "" {
					return externalErrorCode
				}
				return key
			}
			return key
		}
		return internalErrorCode
	}

	return "250000500900705"
}
