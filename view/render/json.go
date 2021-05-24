package render

import (
	"bff-web/domain/model"
	"bff-web/view/model/json"
)

// JSONRender api render interface
type JSONRender interface {
	ConvertErrorResult(err error, code int, messageTypes []string, messageCodes []string, messages []string, messageDetails []string, messageActions []string) *json.ErrorResult
	ConvertHealthcheck(code int) *json.Healthcheck
	ConvertAppV1AuthLogin(code int, authLoginResponse *model.AuthLoginResponse) *json.AuthLoginResponse
	ConvertAppV1AuthLoginSMS(code int, authLoginSMSResponse *model.AuthLoginSMSResponse) *json.AuthLoginSMSResponse
	ConvertAppV1AuthResetRequest(code int) *json.ResetRequestResponse
	ConvertAppV1AuthReset(code int) *json.ResetResponse
	ConvertAppV1AuthChangePassword(code int, authChangePasswordResponse *model.AuthChangePasswordResponse) *json.AuthChangePasswordResponse
	ConvertAppV1OperatorUpdate(code int) *json.OperatorUpdateResponse
	ConvertAppV1OperatorDelete(code int) *json.OperatorDeleteResponse
	ConvertAppV1OperatorRegistration(code int) *json.RegisterResponse
	ConvertAppV1ErrorLogin(code int, errorResult *model.ErrorResult) *json.ErrorResult
	ConvertAppV1GetOperatorListPattern1(code int, operatorListSearchResponse *model.OperatorListSearchResponse) *json.OperatorListSearchResponse
	ConvertAppV1GetOperatorListPattern2(code int, operatorListSearchResponse *model.OperatorListSearchResponsePattern2) *json.OperatorListSearchResponsePattern2
}
