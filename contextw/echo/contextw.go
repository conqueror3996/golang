package echocontextw

import (
	"github.com/labstack/echo/v4"
)

// CompanyIDKey ...
const CompanyIDKey = "xCompanyID"

// SetCompanyID ...
func SetCompanyID(c echo.Context, companyID string) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("contextw/echo/contextw.go : SetCompanyID : func SetCompanyID(c echo.Context, companyID string) start.")))

	if c == nil {
		return
	}

	c.Set(CompanyIDKey, companyID)
}
