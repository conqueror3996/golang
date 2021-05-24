package main

import (
	// "log"
	// "runtime"
	// "net/http"

	// "github.com/pkg/profile"

	// net/http/pprof コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
	// _ "net/http/pprof"

	"fmt"
	"time"

	server "bff-web/af"
	"bff-web/config"

	commonconfig "github.com/nri4nudge/api-common/config"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	"github.com/nri4nudge/api-common/logger/errlog"
	"github.com/nri4nudge/api-common/logger/externaljournallog"
	"github.com/nri4nudge/api-common/logger/internaljournallog"
)

func main() {
	// ***********************************************
	// MessagesConfig 共通設定に業務設定をマージ
	// ***********************************************
	commonconfig.UpdateMessageConfig(config.MessagesConfig)

	// ***********************************************
	// LogMessagesConfig 共通設定に業務設定をマージ
	// ***********************************************
	commonconfig.UpdateLogMessageConfig(config.LogMessagesConfig)

	// ***********************************************
	// AppLog 設定
	// ***********************************************
	apllogLoggerConfig := apllog.LoggerConfig{
		Level:      config.GetProjectConfig().AppEnvConf.Log.AplLog.Level,
		OutputType: config.GetProjectConfig().AppEnvConf.Log.AplLog.OutputType,
	}
	apllog.SetConfig(apllogLoggerConfig)

	// ***********************************************
	// ErrorLog 設定
	// ***********************************************
	errlogLoggerConfig := errlog.LoggerConfig{
		Level:      config.GetProjectConfig().AppEnvConf.Log.ErrLog.Level,
		OutputType: config.GetProjectConfig().AppEnvConf.Log.ErrLog.OutputType,
	}
	errlog.SetConfig(errlogLoggerConfig)

	// ***********************************************
	// ExternalJournalLog 設定
	// ***********************************************
	externaljournallogLoggerConfig := externaljournallog.LoggerConfig{
		Level:      config.GetProjectConfig().AppEnvConf.Log.ExternalJournalLog.Level,
		OutputType: config.GetProjectConfig().AppEnvConf.Log.ExternalJournalLog.OutputType,
	}
	externaljournallog.SetConfig(externaljournallogLoggerConfig)

	// ***********************************************
	// InternalJournalLog 設定
	// ***********************************************
	internaljournallogLoggerConfig := internaljournallog.LoggerConfig{
		Level:      config.GetProjectConfig().AppEnvConf.Log.InternalJournalLog.Level,
		OutputType: config.GetProjectConfig().AppEnvConf.Log.InternalJournalLog.OutputType,
	}
	internaljournallog.SetConfig(internaljournallogLoggerConfig)

	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("main.go : main : func main() start.")))

	// ***********************************************
	// Profile
	// ***********************************************
	// runtime/pprof
	// デフォルトではCPU Profileが有効になっている
	// defer profile.Start(profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	// net/http/pprof
	// runtime.SetBlockProfileRate(1)
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// ***********************************************
	// server run
	// ***********************************************
	server.Run()

	// ***********************************************
	// sleep
	//  ログ出力待ち
	// ***********************************************
	time.Sleep(time.Second * 30)

	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("main.go : main : func main() end.")))
}
