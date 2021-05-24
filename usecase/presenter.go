package usecase

import (
	"github.com/labstack/echo/v4"

	"bff-web/domain/model"
)

// Presenter presenter interface
type Presenter interface {
	JSON(c echo.Context, code int, value interface{})
	JSON2(c echo.Context, status int, value interface{}, headers map[string]string)
	ViewError(c echo.Context, err error)
	ViewErrorMCS(c echo.Context, errorResult *model.ErrorResult)
	ViewHealthcheck(c echo.Context)
	ViewAppV1AuthLogin(c echo.Context, authLoginResponse *model.AuthLoginResponse)
	ViewAppV1AuthLoginSMS(c echo.Context, authLoginSMSResponse *model.AuthLoginSMSResponse)
	ViewAppV1AuthResetRequest(c echo.Context)
	ViewAppV1AuthReset(c echo.Context)
	ViewAppV1AuthChangePassword(c echo.Context, authChangePasswordResponse *model.AuthChangePasswordResponse)
	ViewAppV1OperatorUpdate(c echo.Context)
	ViewAppV1OperatorDelete(c echo.Context)
	ViewAppV1OperatorRegistration(c echo.Context)
	ViewAppV1GetOperatorListPattern1(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponse)
	ViewAppV1GetOperatorListPattern2(c echo.Context, operatorListSearchResponse *model.OperatorListSearchResponsePattern2)
}
