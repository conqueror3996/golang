package userid

import (
	apconfig "bff-web/config"
	appsesson "bff-web/session"

	"github.com/nri4nudge/api-common/contextw"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type (
	// Config defines the config for UserID middleware.
	Config struct {
	}
)

var (
	// DefaultUserIDConfig is the default UserID middleware config.
	DefaultUserIDConfig = Config{}
)

// UserID returns a middleware that ...
func UserID() echo.MiddlewareFunc {
	return WithConfig(DefaultUserIDConfig)
}

// WithConfig returns a UserID middleware with config.
// See: `UserID()`.
func WithConfig(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("middle/echo/userid/user_id.go : WithConfig : func WithConfig(config Config) echo.MiddlewareFunc start.")))

			// err = next(c)
			// return err

			sess, err := session.Get(apconfig.GetProjectConfig().AppEnvConf.Session[0].SessionID, c)
			if err == nil {
				r := c.Request()
				ctx := r.Context()

				if sess.Values[appsesson.SessionVarOperatorAuthID] != nil {
					ctx = contextw.SetUserID(ctx, sess.Values[appsesson.SessionVarOperatorAuthID].(string))
				} else {
					ctx = contextw.SetUserID(ctx, "")
				}
				r = r.WithContext(ctx)
				c.SetRequest(r)
			} else {
			}

			err = next(c)
			return err

		}
	}
}
