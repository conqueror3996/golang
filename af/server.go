package af

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	view "bff-web/adapter/view/render"
	"bff-web/config"
	"bff-web/interface/api"
	"bff-web/middle/echo/userid"
	"bff-web/session/datastore"
	"bff-web/usecase"

	apperrorsub "github.com/nri4nudge/api-common/adapter/error"
	commongormdb "github.com/nri4nudge/api-common/adapter/gormdb"
	commongormdbreadreplica "github.com/nri4nudge/api-common/adapter/gormdbreadreplica"
	"github.com/nri4nudge/api-common/logger"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	"github.com/nri4nudge/api-common/logger/errlog"
	"github.com/nri4nudge/api-common/logger/middleware/echo/accesslogger"
	"github.com/nri4nudge/api-common/middleware/realip"
	"github.com/nri4nudge/api-common/middleware/requestinfo"
	commonusecase "github.com/nri4nudge/api-common/usecase"
)

//Run ...
func Run() {
	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("af/server.go : Run : func Run() start.")))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/server.go : Run : config.GetProjectConfig().AppEnvConf.Env=" + config.GetProjectConfig().AppEnvConf.Env)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("af/server.go : Run : config.GetProjectConfig().AppEnvConf.SitePrefix=" + config.GetProjectConfig().AppEnvConf.SitePrefix)))

	// ***********************************************
	// echo new
	// ***********************************************
	e := echo.New()

	e.HTTPErrorHandler = CustomErrorHandler
	e.HideBanner = true
	e.HidePort = true
	e.Debug = false

	// ***********************************************
	// middleware RecoverWithConfig 設定
	// ***********************************************
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))

	// ***********************************************
	// middleware requestinfo 設定
	// ***********************************************
	e.Use(echo.WrapMiddleware(requestinfo.RequestInfo))

	// ***********************************************
	// middleware realip 設定
	// ***********************************************
	e.Use(echo.WrapMiddleware(realip.RealIP))

	// store, err := datastore.GetSessionStoreFromConnection((*commongormdatabase.GetDBM()).GetSQLDB())
	// if err != nil {
	// 	errm := apperrorsub.NewErrorManager()
	// 	msgCD := "0100005000107"
	// 	errw := errm.Wrap1(
	// 		err,
	// 		"ja",
	// 		msgCD,
	// 		"",
	// 	)
	// 	errlog.Errorf(errm.LogMessageV2(errw, logger.LogTypeErrorLog, logger.LogLevel.M[logger.ErrorLevel], "LE0100000301"))
	// 	os.Exit(1)
	// }

	// e.Use(session.Middleware(store))
	// defer store.Close()

	store, err := datastore.GetSessionStoreFromConnectionDynamoDB()
	if err != nil {
		errm := apperrorsub.NewErrorManager()
		msgCD := "0100005000107"
		errw := errm.Wrap1(
			err,
			"ja",
			msgCD,
			"",
		)
		errlog.Errorf(errm.LogMessageV2(errw, logger.LogTypeErrorLog, logger.LogLevel.M[logger.ErrorLevel], "LE0100000301"))
		os.Exit(1)
	}

	e.Use(session.Middleware(store))
	//defer store.Close()

	e.Use(userid.UserID())
	// ***********************************************
	// middleware accesslogger 設定
	// ***********************************************
	// e.Use(accesslogger.Logger())
	e.Use(accesslogger.LoggerWithConfig(accesslogger.LoggerConfig{
		Level:      config.GetProjectConfig().AppEnvConf.Log.AccessLog.Level,
		OutputType: config.GetProjectConfig().AppEnvConf.Log.AccessLog.OutputType,
	}))

	if config.GetProjectConfig().AppEnvConf.App.Cors.CorsSetting {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     config.GetProjectConfig().AppEnvConf.App.Cors.HeaderAccessControlAllowOrigin,
			AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowCredentials: config.GetProjectConfig().AppEnvConf.App.Cors.HeaderAccessControlAllowCredentials,
		}))
	}
	// ***********************************************
	// middleware BodyLimit 設定
	// ***********************************************
	e.Use(middleware.BodyLimit(config.GetProjectConfig().AppEnvConf.App.Request.ContentLengthLimit))

	// ***********************************************
	// 各レイヤのオブジェクト生成
	// ***********************************************
	jsonRender := view.NewJSONRender()

	errorManager := apperrorsub.NewErrorManager()

	presenter := api.NewPresenter(jsonRender, errorManager)

	commonTransactionRepository := commongormdb.NewTransactionRepository()
	commonReadReplicaTransactionRepository := commongormdbreadreplica.NewTransactionRepository()

	commoncommon := commonusecase.NewCommon(errorManager, commonTransactionRepository, commonReadReplicaTransactionRepository)

	operator := usecase.NewOperator(commonTransactionRepository, commonReadReplicaTransactionRepository, errorManager)

	common := usecase.NewCommon(
		presenter,
		errorManager,
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		commoncommon,
		operator,
	)

	interactor := usecase.NewInteractor(
		presenter,
		errorManager,
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		commoncommon,
		common,
		operator,
	)

	api := api.NewAPI(interactor)

	// ***********************************************
	// ルーティング 設定
	// ***********************************************
	// API259001_ヘルスチェック
	e.GET("/healthcheck", HandleAppV1Healthcheck(api))

	// BFW010101_認証API
	e.POST(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/auth/login", HandleAppV1AuthLogin(api))

	// BFW010102_パスワード再設定リクエストAPI
	e.POST(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/auth/reset_request", HandleAppV1AuthResetRequest(api))

	// BFW010103_パスワード再設定API
	e.PUT(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/auth/reset", HandleAppV1AuthReset(api))

	// BFW010104_パスワード変更API
	e.PUT(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/auth/change", HandleAppV1AuthChangePassword(api))

	// BFW010501_オペレーター作成API
	e.POST(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/operator/register", HandleAppV1OperatorRegistration(api))

	// BFW010502_オペレーター更新API
	e.POST(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/operator/update", HandleAppV1OperatorUpdate(api))

	// BFW010503_オペレーター削除API
	e.PUT(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/operator/delete", HandleAppV1OperatorDelete(api))

	// BFW010504_オペレータ検索API
	e.POST(config.GetProjectConfig().AppEnvConf.SitePrefix+"/v1/operator/list", HandleAppV1GetOperatorList(api))

	// ***********************************************
	// server start
	// ***********************************************
	go func() {
		apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("af/server.go : Run : e.Start")))
		apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("%s", e.Start(config.GetProjectConfig().AppEnvConf.Port))))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("af/server.go : Run : before <-quit")))
	<-quit
	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("af/server.go : Run : after <-quit")))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("af/server.go : Run : e.Shutdown(ctx)")))
	if err := e.Shutdown(ctx); err != nil {
		apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("%s", err)))
	}
}
