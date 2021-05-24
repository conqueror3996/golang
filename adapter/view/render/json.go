package view

import (
	"bff-web/domain/model"
	"bff-web/view/model/json"
	"bff-web/view/render"
)

// NewJSONRender new json render
func NewJSONRender() render.JSONRender {
	return jsonRender{}
}

type jsonRender struct {
}

func (r jsonRender) ConvertErrorResult(err error, code int, messageTypes []string, messageCodes []string, messages []string, messageDetails []string, messageActions []string) *json.ErrorResult {

	cMessageCode := ""
	cMessage := ""
	cMessageDetail := ""
	cMessageAction := ""

	cMessageCode = map[bool]string{true: messageCodes[0], false: ""}[len(messageCodes) > 0]
	cMessage = map[bool]string{true: messages[0], false: ""}[len(messages) > 0]
	cMessageDetail = map[bool]string{true: messageDetails[0], false: ""}[len(messageDetails) > 0]
	cMessageAction = map[bool]string{true: messageActions[0], false: ""}[len(messageActions) > 0]

	cError := json.Error{
		Status:  code,
		Code:    cMessageCode,
		Message: cMessage,
		Detail:  cMessageDetail,
		Action:  cMessageAction,
	}
	return &json.ErrorResult{Error: cError}
}

func (r jsonRender) ConvertHealthcheck(code int) *json.Healthcheck {

	return &json.Healthcheck{}
}

func (r jsonRender) ConvertAppV1AuthLogin(code int, authLoginResponse *model.AuthLoginResponse) *json.AuthLoginResponse {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/view/render/json.go : ConvertAppV1AuthLogin : func (r jsonRender) ConvertAppV1AuthLogin(code int) *json.AuthLogin start.")))

	return &json.AuthLoginResponse{
		LoginStatus: authLoginResponse.LoginStatus,
	}
}

func (r jsonRender) ConvertAppV1AuthLoginSMS(code int, authLoginSMSResponse *model.AuthLoginSMSResponse) *json.AuthLoginSMSResponse {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/view/render/json.go : ConvertAppV1AuthLogin : func (r jsonRender) ConvertAppV1AuthLogin(code int) *json.AuthLogin start.")))

	role := &json.RoleResponse{
		CardRole:     authLoginSMSResponse.Role.CardRole,
		PartNerRole:  authLoginSMSResponse.Role.PartnerRole,
		SkinRole:     authLoginSMSResponse.Role.SkinRole,
		OperatorRole: authLoginSMSResponse.Role.OperatorRole,
		MessageRole:  authLoginSMSResponse.Role.MessageRole,
	}

	return &json.AuthLoginSMSResponse{
		LoginStatus: authLoginSMSResponse.LoginStatus,
		UserName:    authLoginSMSResponse.UserName,
		Role:        *role,
	}
}

func (r jsonRender) ConvertAppV1ErrorLogin(code int, errorResult *model.ErrorResult) *json.ErrorResult { //dang code
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/view/render/json.go : ConvertAppV1AuthLogin : func (r jsonRender) ConvertAppV1AuthLogin(code int) *json.AuthLogin start.")))

	err := &json.Error{
		Status:  errorResult.Error.Status,
		Code:    errorResult.Error.Code,
		Message: errorResult.Error.Message,
		Detail:  errorResult.Error.Detail,
		Action:  errorResult.Error.Action,
	}

	return &json.ErrorResult{
		Error: *err,
	}
}

func (r jsonRender) ConvertAppV1AuthResetRequest(code int) *json.ResetRequestResponse {
	return &json.ResetRequestResponse{}
}

func (r jsonRender) ConvertAppV1AuthReset(code int) *json.ResetResponse {
	return &json.ResetResponse{}
}

func (r jsonRender) ConvertAppV1AuthChangePassword(code int, authChangePasswordResponse *model.AuthChangePasswordResponse) *json.AuthChangePasswordResponse {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/view/render/json.go : ConvertAppV1EmployeePassword : func (r jsonRender) ConvertAppV1EmployeePassword(code int) *json.EmployeePassword start.")))

	role := &json.RoleChangeResponse{
		CardMemberRole: authChangePasswordResponse.Role.CardMemberRole,
		ClubOwnerRole:  authChangePasswordResponse.Role.ClubOwnerRole,
		ClubRole:       authChangePasswordResponse.Role.ClubRole,
		OperatorRole:   authChangePasswordResponse.Role.OperatorRole,
		MessageRole:    authChangePasswordResponse.Role.MessageRole,
	}

	return &json.AuthChangePasswordResponse{
		LoginStatus: authChangePasswordResponse.LoginStatus,
		UserName:    authChangePasswordResponse.UserName,
		Role:        *role,
	}
}

func (r jsonRender) ConvertAppV1OperatorUpdate(code int) *json.OperatorUpdateResponse {

	return &json.OperatorUpdateResponse{}
}

func (r jsonRender) ConvertAppV1OperatorDelete(code int) *json.OperatorDeleteResponse {

	return &json.OperatorDeleteResponse{}
}

func (r jsonRender) ConvertAppV1OperatorRegistration(code int) *json.RegisterResponse {
	return &json.RegisterResponse{}
}

func (r jsonRender) ConvertAppV1GetOperatorListPattern1(code int, operatorListSearchResponse *model.OperatorListSearchResponse) *json.OperatorListSearchResponse {

	vOperatorSearchResponse := []*json.OperatorSearchResponse{}
	for _, v := range operatorListSearchResponse.Result.OperatorList {
		role := &json.OperatorRoleInfoResponse{
			CardMemberRole: v.CardMemberRole,
			ClubRole:       v.ClubRole,
			ClubOwnerRole:  v.ClubOwnerRole,
			MessageRole:    v.MessageRole,
			OperatorRole:   v.OperatorRole,
		}
		rj := &json.OperatorSearchResponse{
			OperatorID:     v.OperatorID,
			OperatorRoleID: v.OperatorRoleID,
			Username:       v.Username,
			MailAddress:    v.MailAddress,
			PhoneNumber:    v.PhoneNumber,
			Belong:         v.Belong,
			Role:           *role,
		}
		vOperatorSearchResponse = append(vOperatorSearchResponse, rj)
	}
	operatorResult := &json.OperatorResult{
		MaxPage:      operatorListSearchResponse.Result.MaxPage,
		OperatorList: vOperatorSearchResponse,
	}
	return &json.OperatorListSearchResponse{Result: *operatorResult}
}

func (r jsonRender) ConvertAppV1GetOperatorListPattern2(code int, operatorListSearchResponse *model.OperatorListSearchResponsePattern2) *json.OperatorListSearchResponsePattern2 {

	vOperatorSearchResponse := &json.OperatorListSearchResponsePattern2{}

	vOperatorSearchResponse.Result.Name = operatorListSearchResponse.Result.Username
	vOperatorSearchResponse.Result.MailAddress = operatorListSearchResponse.Result.MailAddress
	vOperatorSearchResponse.Result.PhoneNumber = operatorListSearchResponse.Result.PhoneNumber
	vOperatorSearchResponse.Result.Belong = operatorListSearchResponse.Result.Belong
	vOperatorSearchResponse.Result.AccountStatus = operatorListSearchResponse.Result.AccountStatus

	vOperatorSearchResponse.Result.Role.CardMemberRole = operatorListSearchResponse.Result.Role.CardMemberRole
	vOperatorSearchResponse.Result.Role.ClubRole = operatorListSearchResponse.Result.Role.ClubRole
	vOperatorSearchResponse.Result.Role.ClubOwnerRole = operatorListSearchResponse.Result.Role.ClubOwnerRole
	vOperatorSearchResponse.Result.Role.MessageRole = operatorListSearchResponse.Result.Role.MessageRole
	vOperatorSearchResponse.Result.Role.OperatorRole = operatorListSearchResponse.Result.Role.OperatorRole

	operatorResult := &json.OperatorListSearchResponsePattern2{
		Result: vOperatorSearchResponse.Result,
	}
	return operatorResult
}
