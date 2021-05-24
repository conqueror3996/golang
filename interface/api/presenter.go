package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"bff-web/domain/model"
	"bff-web/usecase"
	"bff-web/view/render"

	echocontextw "github.com/nri4nudge/api-common/contextw/echo"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	"github.com/nri4nudge/api-common/logger/errlog"
)

// NewPresenter new presenter
func NewPresenter(render render.JSONRender, errm apperror.ErrorManager) usecase.Presenter {
	return presenter{
		render,
		errm,
	}
}

type presenter struct {
	render render.JSONRender
	errm   apperror.ErrorManager
}

// JSON render json format
func (m presenter) JSON(c echo.Context, status int, value interface{}) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : JSON : func (m presenter) JSON(c echo.Context, status int, value interface{}) start.")))

	b, err := json.Marshal(value)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}
	echocontextw.SetResponseBody(c, string(b))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("X-Content-Type-Options", "nosniff")
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set(echo.HeaderStrictTransportSecurity, "max-age=31536000, includeSubDomains; preload;")
	c.Response().Header().Set(echo.HeaderServer, "None")
	c.Response().Header().Set("Connection", "close")
	c.Response().WriteHeader(status)
	c.Response().Write(b)
}

// JSON2 render json format
func (m presenter) JSON2(c echo.Context, status int, value interface{}, headers map[string]string) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : JSON2 : JSON2(c echo.Context, status int, value interface{}, headers map[string]string) start.")))

	b, err := json.Marshal(value)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("X-Content-Type-Options", "nosniff")
	// c.Response().Header().Set(echo.HeaderStrictTransportSecurity, "max-age=15768000")
	// c.Response().Header().Set(echo.HeaderServer, "None")
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Pragma", "no-cache")
	// c.Response().Header().Set("Connection", "close")

	for rk, rv := range headers {
		c.Response().Header().Set(rk, rv)
	}

	c.Response().WriteHeader(status)
	c.Response().Write(b)
}

func (m presenter) ViewError(c echo.Context, errEntity error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewError : func (m presenter) ViewError(c echo.Context, errEntity error) start.")))

	// ***********************************************
	// AccessLog errorCode 出力内容設定
	// ***********************************************
	echocontextw.SetErrorCode(c, map[bool]string{true: m.errm.MessageCodes(errEntity)[0], false: ""}[len(m.errm.MessageCodes(errEntity)) > 0])

	// ***********************************************
	// ErrorLog 出力
	// ***********************************************
	messageLevel := logger.ValidateLogLevelErrorLog(map[bool]string{true: m.errm.MessageLevels(errEntity)[0], false: ""}[len(m.errm.MessageLevels(errEntity)) > 0], m.errm.Code(errEntity))
	logCode := "LE0100000101"
	if messageLevel == "warn" {
		logCode = "LW0100000102"
	} else if messageLevel == "info" || messageLevel == "debug" || messageLevel == "trace" {
		logCode = "LI0100000103"
	}
	errlog.LogfWLevel(messageLevel, m.errm.LogMessageV1(c.Request().Context(), c, errEntity, logger.LogTypeErrorLog, messageLevel, logCode))

	// ***********************************************
	// AccessLog ResponseBody 出力内容設定
	// ***********************************************
	responseForLog := m.render.ConvertErrorResult(errEntity, m.errm.Code(errEntity), m.errm.MessageTypes(errEntity), m.errm.MessageCodes(errEntity), m.errm.Messages(errEntity), m.errm.MessageDetails(errEntity), m.errm.MessageActions(errEntity))

	responseJSONForLog, err := json.Marshal(responseForLog)
	if err != nil {
		responseJSONForLog = []byte("{}")
	}
	echocontextw.SetResponseBody(c, string(responseJSONForLog))

	// ***********************************************
	// 応答
	// ***********************************************
	m.JSON(c, m.errm.Code(errEntity), m.render.ConvertErrorResult(errEntity, m.errm.Code(errEntity), m.errm.MessageTypes(errEntity), m.errm.MessageCodes(errEntity), m.errm.Messages(errEntity), m.errm.MessageDetails(errEntity), m.errm.MessageActions(errEntity)))
}

func (m presenter) ViewErrorMCS(c echo.Context, errEntity *model.ErrorResult) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewErrorMCS : func (m presenter) ViewErrorMCS(c echo.Context, errEntity Error) start.")))

	m.JSON(c, errEntity.Error.Status, m.render.ConvertAppV1ErrorLogin(errEntity.Error.Status, errEntity))
}

func (m presenter) ViewHealthcheck(c echo.Context) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewHealthcheck : func (m presenter) ViewHealthcheck(c echo.Context) start.")))

	// ***********************************************
	// AccessLog ResponseBody 出力内容設定
	// ***********************************************
	responseForLog := m.render.ConvertHealthcheck(http.StatusOK)

	responseJSONForLog, err := json.Marshal(responseForLog)
	if err != nil {
		responseJSONForLog = []byte("{}")
	}
	echocontextw.SetResponseBody(c, string(responseJSONForLog))

	// ***********************************************
	// 応答
	// ***********************************************
	m.JSON(c, http.StatusOK, m.render.ConvertHealthcheck(http.StatusOK))
}

func (m presenter) ViewAppV1AuthLogin(c echo.Context, authLoginResponse *model.AuthLoginResponse) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1AuthLogin : func (m presenter) ViewAppV1AuthLogin(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1AuthLogin(http.StatusOK, authLoginResponse))
}

func (m presenter) ViewAppV1AuthLoginSMS(c echo.Context, authLoginSMSResponse *model.AuthLoginSMSResponse) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1AuthLogin : func (m presenter) ViewAppV1AuthLogin(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1AuthLoginSMS(http.StatusOK, authLoginSMSResponse))
}

func (m presenter) ViewAppV1AuthChangePassword(c echo.Context, authChangePasswordResponse *model.AuthChangePasswordResponse) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1AuthChangePassword : func (m presenter) ViewAppV1AuthChangePassword(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1AuthChangePassword(http.StatusOK, authChangePasswordResponse))
}

func (m presenter) ViewAppV1AuthResetRequest(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1AuthResetRequest : func (m presenter) ViewAppV1AuthResetRequest(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1AuthResetRequest(http.StatusOK))
}

func (m presenter) ViewAppV1AuthReset(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1AuthReset : func (m presenter) ViewAppV1AuthReset(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1AuthReset(http.StatusOK))
}

func (m presenter) ViewAppV1OperatorRegistration(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1OperatorRegistration : func (m presenter) ViewAppV1OperatorRegistration(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1OperatorRegistration(http.StatusOK))
}

func (m presenter) ViewAppV1OperatorUpdate(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1OperatorUpdate : func (m presenter) ViewAppV1OperatorUpdate(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1OperatorUpdate(http.StatusOK))
}

func (m presenter) ViewAppV1OperatorDelete(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1OperatorDelete : func (m presenter) ViewAppV1OperatorDelete(c echo.Context) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1OperatorDelete(http.StatusOK))
}

func (m presenter) ViewAppV1GetOperatorListPattern1(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponse) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1GetOperatorListPattern1 : func (m presenter) ViewAppV1GetOperatorListPattern1(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponse) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1GetOperatorListPattern1(http.StatusOK, operatorListSearchResponse))
}

func (m presenter) ViewAppV1GetOperatorListPattern2(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponsePattern2) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/presenter.go : ViewAppV1GetOperatorListPattern2 : func (m presenter) ViewAppV1GetOperatorListPattern2(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponsePattern2) start.")))

	m.JSON(c, http.StatusOK, m.render.ConvertAppV1GetOperatorListPattern2(http.StatusOK, operatorListSearchResponse))
}
