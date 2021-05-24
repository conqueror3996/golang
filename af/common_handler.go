package af

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"

	"bff-web/interface/api"
	modeljson "bff-web/view/model/json"

	apperrorsub "github.com/nri4nudge/api-common/adapter/error"
	"github.com/nri4nudge/api-common/config"
	"github.com/nri4nudge/api-common/contextw"
	echocontextw "github.com/nri4nudge/api-common/contextw/echo"
	"github.com/nri4nudge/api-common/logger"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	"github.com/nri4nudge/api-common/logger/errlog"
)

// CustomErrorHandler ...
func CustomErrorHandler(err error, c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : CustomErrorHandler :  func CustomErrorHandler(err error, c echo.Context) start.")))

	msgCD := "250000500900101"
	code := config.GetMessageStatusCode(msgCD)
	cMessageLevel := config.GetMessageMessageLevel(msgCD)
	cMessageType := config.GetMessageMessageType(msgCD)
	cMessageOrgCode := msgCD
	cMessageCode := config.GetMessageMessageCode(msgCD)
	cMessage := config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
	cMessageDetail := config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
	cMessageAction := config.GetMessageMessageAction(msgCD)
	cMessageAddInfo := ""
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == 400 {
			// message = he.Message.(string)
			code = he.Code
			msgCD = "250000400900103"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 401 {
			code = he.Code
			msgCD = "250000401900104"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 403 {
			code = he.Code
			msgCD = "250000403900105"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 404 {
			code = he.Code
			msgCD = "250000404900106"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 405 {
			code = he.Code
			msgCD = "250000405900109"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 413 {
			code = he.Code
			msgCD = "250000413900111"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 500 {
			code = he.Code
			msgCD = "250000500900107"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else if he.Code == 503 {
			code = he.Code
			msgCD = "250000503900108"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		} else {
			code = he.Code
			msgCD = "250000000900102"
			cMessageLevel = config.GetMessageMessageLevel(msgCD)
			cMessageType = config.GetMessageMessageType(msgCD)
			cMessageOrgCode = msgCD
			cMessageCode = config.GetMessageMessageCode(msgCD)
			cMessage = config.GetMessageMessage(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageDetail = config.GetMessageMessageDetail(contextw.GetRequestLang(c.Request().Context()), msgCD)
			cMessageAction = config.GetMessageMessageAction(msgCD)
			cMessageAddInfo = ""
		}
	}

	errm := apperrorsub.NewErrorManager()

	if errm != nil {
		errw := errm.Wrap3(
			err,
			code,
			cMessageLevel,
			cMessageType,
			cMessageOrgCode,
			cMessageCode,
			cMessage,
			cMessageDetail,
			cMessageAction,
			cMessageAddInfo,
		)

		// ***********************************************
		// AccessLog errorCode 出力内容設定
		// ***********************************************
		echocontextw.SetErrorCode(c, map[bool]string{true: errm.MessageCodes(errw)[0], false: ""}[len(errm.MessageCodes(errw)) > 0])

		// ***********************************************
		// ErrorLog 出力
		// ***********************************************
		messageLevel := logger.ValidateLogLevelErrorLog(cMessageLevel, code)
		logCode := "LE0100000101"
		if messageLevel == "warn" {
			logCode = "LW0100000102"
		} else if messageLevel == "info" || messageLevel == "debug" || messageLevel == "trace" {
			logCode = "LI0100000103"
		}
		errlog.LogfWLevel(messageLevel, errm.LogMessageV1(c.Request().Context(), c, errw, logger.LogTypeErrorLog, messageLevel, logCode))
	}

	cError := modeljson.Error{
		Status:  code,
		Code:    cMessageCode,
		Message: cMessage,
		Detail:  cMessageDetail,
		Action:  cMessageAction,
	}
	errorResult := modeljson.ErrorResult{Error: cError}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("X-Content-Type-Options", "nosniff")
	// c.Response().Header().Set(echo.HeaderStrictTransportSecurity, "max-age=15768000")
	// c.Response().Header().Set(echo.HeaderServer, "None")
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Pragma", "no-cache")
	// c.Response().Header().Set("Connection", "close")

	// ***********************************************
	// AccessLog ResponseBody 出力内容設定
	// ***********************************************
	errorResultJSONString, err := json.Marshal(errorResult)
	if err != nil {
		errorResultJSONString = []byte("{}")
	}
	echocontextw.SetResponseBody(c, string(errorResultJSONString))

	// ***********************************************
	// 応答
	// ***********************************************
	c.JSON(code, errorResult)
}

// HandleAppV1Healthcheck ...
func HandleAppV1Healthcheck(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1Healthcheck : func HandleAppV1Healthcheck(api api.API) echo.HandlerFunc start.")))

		api.AppV1Healthcheck(c)

		return nil
	}
}

// HandleAppV1AuthLogin ...
func HandleAppV1AuthLogin(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1AuthLogin : func HandleAppV1AuthLogin(api api.API) echo.HandlerFunc start.")))

		api.AppV1AuthLogin(c)

		return nil
	}
}

// HandleAppV1AuthReset ...
func HandleAppV1AuthReset(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1AuthReset : func HandleAppV1AuthReset(api api.API) echo.HandlerFunc start.")))

		api.AppV1AuthReset(c)
		return nil
	}
}

// HandleAppV1AuthResetRequest ...
func HandleAppV1AuthResetRequest(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1AuthResetRequest : func HandleAppV1AuthResetRequest(api api.API) echo.HandlerFunc start.")))

		api.AppV1AuthResetRequest(c)

		return nil
	}
}

// HandleAppV1AuthChangePassword ...
func HandleAppV1AuthChangePassword(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1AuthChangePassword : func HandleAppV1AuthChangePassword(api api.API) echo.HandlerFunc start.")))

		api.AppV1AuthChangePassword(c)

		return nil
	}
}

// HandleAppV1OperatorRegistration ...
func HandleAppV1OperatorRegistration(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1RegisterOperator : func HandleAppV1RegisterOperator(api api.API) echo.HandlerFunc start.")))

		api.AppV1OperatorRegistration(c)
		return nil
	}
}

// HandleAppV1OperatorUpdate ...
func HandleAppV1OperatorUpdate(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1OperatorUpdate : func HandleAppV1OperatorUpdate(api api.API) echo.HandlerFunc start.")))

		api.AppV1OperatorUpdate(c)
		return nil
	}
}

// HandleAppV1OperatorDelete ...
func HandleAppV1OperatorDelete(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1OperatorDelete : func HandleAppV1OperatorDelete(api api.API) echo.HandlerFunc start.")))

		api.AppV1OperatorDelete(c)
		return nil
	}
}

// HandleAppV1GetOperatorList ...
func HandleAppV1GetOperatorList(api api.API) echo.HandlerFunc {
	return func(c echo.Context) error {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/common_handler.go : HandleAppV1GetOperatorList : func HandleAppV1GetOperatorList(api api.API) echo.HandlerFunc start.")))

		api.AppV1GetOperatorList(c)
		return nil
	}
}
