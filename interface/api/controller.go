package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"

	"bff-web/usecase"
	"bff-web/usecase/input"

	echocontextw "github.com/nri4nudge/api-common/contextw/echo"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

// NewAPI Get API instance
func NewAPI(it usecase.Interactor) API {
	return controller{it}
}

// API api instance
type API interface {
	AppV1Healthcheck(c echo.Context)
	AppV1AuthLogin(c echo.Context)
	AppV1AuthReset(c echo.Context)
	AppV1AuthResetRequest(c echo.Context)
	AppV1AuthChangePassword(c echo.Context)
	AppV1OperatorRegistration(c echo.Context)
	AppV1OperatorUpdate(c echo.Context)
	AppV1OperatorDelete(c echo.Context)
	AppV1GetOperatorList(c echo.Context)
}

type controller struct {
	it usecase.Interactor
}

// AppV1Healthcheck ...
func (ctl controller) AppV1Healthcheck(c echo.Context) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("%s", "interface/api/controller.go : "+"AppV1Healthcheck : "+"func (ctl controller) AppV1Healthcheck(c echo.Context) start.")))

	ipt := &input.Healthcheck{}

	echocontextw.SetResponseBody(c, "")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1Healthcheck : ctl.it.AppV1Healthcheck call.")))
	ctl.it.AppV1Healthcheck(c, *ipt)
}

// AppV1AuthLogin ...
func (ctl controller) AppV1AuthLogin(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthLogin : func (ctl controller) AppV1AuthLogin(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderAuthorization:       c.Request().Header.Get("Authorization"),
	}

	ipt := new(input.AuthLogin)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthLogin : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1AuthLogin(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthLogin : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1AuthLogin(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthLogin : ctl.it.AppV1AuthLogin(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1AuthLogin(c, *ipt, *iptMeta)
}

// AppV1AuthReset ...
func (ctl controller) AppV1AuthReset(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthReset : func (ctl controller) AppV1AuthReset(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
	}

	ipt := new(input.AuthReset)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthReset : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1AuthReset(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthReset : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1AuthReset(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthReset : ctl.it.AppV1AuthReset(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1AuthReset(c, *ipt, *iptMeta)
}

// AppV1AuthChangePassword ...
func (ctl controller) AppV1AuthChangePassword(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthChangePassword : func (ctl controller) AppV1AuthChangePassword(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
		HeaderAuthorization:       c.Request().Header.Get("Authorization"),
	}

	ipt := new(input.AuthChangePassword)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthChangePassword : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1AuthChangePassword(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthChangePassword : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1AuthChangePassword(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthChangePassword : ctl.it.AppV1AuthChangePassword(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1AuthChangePassword(c, *ipt, *iptMeta)
}

// AppV1AuthResetRequest ...
func (ctl controller) AppV1AuthResetRequest(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthResetRequest : func (ctl controller) AppV1AuthResetRequest(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
	}

	ipt := new(input.ResetRequest)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthResetRequest : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1AuthResetRequest(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthResetRequest : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1AuthResetRequest(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1AuthResetRequest : ctl.it.AppV1AuthResetRequest(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1AuthResetRequest(c, *ipt, *iptMeta)
}

// AppV1OperatorRegistration ...
func (ctl controller) AppV1OperatorRegistration(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorRegistration : func (ctl controller) AppV1OperatorRegistration(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
	}

	ipt := new(input.OperatorRegistration)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorRegistration : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1OperatorRegistration(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorRegistration : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1OperatorRegistration(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorRegistration : ctl.it.OperatorRegistration(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1OperatorRegistration(c, *ipt, *iptMeta)
}

// AppV1OperatorUpdate ...
func (ctl controller) AppV1OperatorUpdate(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorUpdate : func (ctl controller) AppV1OperatorUpdate(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
	}

	ipt := new(input.OperatorUpdate)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorUpdate : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1OperatorUpdate(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorUpdate : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1OperatorUpdate(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorUpdate : ctl.it.AppV1OperatorUpdate(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1OperatorUpdate(c, *ipt, *iptMeta)
}

// AppV1OperatorDelete ...
func (ctl controller) AppV1OperatorDelete(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorDelete : func (ctl controller) AppV1OperatorDelete(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
	}

	ipt := new(input.OperatorDelete)

	jsonBlob, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorDelete : ioutil.ReadAll(c.Request().Body) - err != nil")))
		ctl.it.AppV1OperatorDelete(c, *ipt, *iptMeta)
		return
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(jsonBlob)) // Reset

	err = json.Unmarshal(jsonBlob, &ipt)
	if err != nil {
		iptMeta.ReceptionProcessingStatus = "1"
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorDelete : json.Unmarshal(jsonBlob, &ipt) - err != nil")))
		ctl.it.AppV1OperatorDelete(c, *ipt, *iptMeta)
		return
	}

	iptMeta.Body = string(jsonBlob)

	// ***********************************************
	// AccessLog RequestBody 出力内容設定
	// ***********************************************
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// TODO AccessLog request.body 出力内容精査
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
	echocontextw.SetRequestBody(c, iptMeta.Body)

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1OperatorDelete : ctl.it.AppV1OperatorDelete(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1OperatorDelete(c, *ipt, *iptMeta)
}

// AppV1GetOperatorList ...
func (ctl controller) AppV1GetOperatorList(c echo.Context) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1GetOperatorList : func (ctl controller) AppV1GetOperatorList(c echo.Context) start.")))

	iptMeta := &input.MetaRequest{
		ReceptionProcessingStatus: "0",
		HeaderContentType:         c.Request().Header.Get("Content-Type"),
		HeaderUserAgent:           c.Request().Header.Get("User-Agent"),
		HeaderXRequestedWith:      c.Request().Header.Get("X-Requested-With"),
		HeaderXNudgeInteractionID: c.Request().Header.Get("X-Nudge-Interaction-Id"),
		HeaderXAPIKey:             c.Request().Header.Get("X-API-KEY"),
	}

	ipt := &input.OperatorList{
		SearchText: c.QueryParam("search_text"),
		SortKey:    c.QueryParam("sort_key"),
		SortOrder:  c.QueryParam("sort_order"),
		PageNum:    c.QueryParam("page_num"),
		PageSize:   c.QueryParam("page_size"),
		OperatorID: c.QueryParam("operator_id"),
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("interface/api/controller.go : AppV1GetOperatorList : ctl.it.AppV1GetOperatorList(c, *ipt, *iptMeta) call.")))
	ctl.it.AppV1GetOperatorList(c, *ipt, *iptMeta)
}
