package config

import (
	"encoding/json"
	"strconv"

	"github.com/kelseyhightower/envconfig"

	devconfig "bff-web/config/dev"
	localconfig "bff-web/config/local"
	prodconfig "bff-web/config/prod"
	stgconfig "bff-web/config/stg"
)

// OSEnvConfig ...
type OSEnvConfig struct {
	ActiveEnv               string `envconfig:"ACTIVE_ENV"`
	AppID                   string `envconfig:"APP_ID"`
	AppPort                 string `envconfig:"APP_PORT"`
	ApllogLevel             string `envconfig:"APLLOG_LEVEL"`
	ErrlogLevel             string `envconfig:"ERRLOG_LEVEL"`
	AccesslogLevel          string `envconfig:"ACCESSLOG_LEVEL"`
	ExternaljournalLogLevel string `envconfig:"EXTERNALJOURNALLOG_LEVEL"`
	InternaljournalLogLevel string `envconfig:"INTERNALJOURNALLOG_LEVEL"`
	NotificationLogLevel    string `envconfig:"NOTIFICATIONLOG_LEVEL"`
	DbHost                  string `envconfig:"DB_HOST"`
	DbHostRo                string `envconfig:"DB_HOST_RO"`
	DbName                  string `envconfig:"DB_NAME"`
	DbUser                  string `envconfig:"DB_USER"`
	DbPassword              string `envconfig:"DB_PASSWORD"`
	DbMaxIdleConns          string `envconfig:"DB_MAX_IDLE_CONNS"`
	DbMaxOpenConns          string `envconfig:"DB_MAX_OPEN_CONNS"`
	DbConnMaxLifetime       string `envconfig:"DB_CONN_MAX_LIFE_TIME"`
	DbMaxIdleConnsRo        string `envconfig:"DB_MAX_IDLE_CONNS_RO"`
	DbMaxOpenConnsRo        string `envconfig:"DB_MAX_OPEN_CONNS_RO"`
	DbConnMaxLifetimeRo     string `envconfig:"DB_CONN_MAX_LIFE_TIME_RO"`
	SmSecretID              string `envconfig:"SM_SECRET_ID"`
	SvcRepaymentBasePath    string `envconfig:"SVC_REPAYMENT_BASE_PATH"`
	SvcPaymentBasePath      string `envconfig:"SVC_PAYMENT_BASE_PATH"`
}

// ProjectConfig ...
type ProjectConfig struct {
	AppComConf AppComConfig
	AppEnvConf AppEnvConfig
	OSEnvConf  OSEnvConfig
}

// AppComConfig ...
type AppComConfig struct {
	ActiveEnv string `json:"activeEnv"`
}

// AppEnvConfig ...
type AppEnvConfig struct {
	Env             string                    `json:"env"`
	AppID           string                    `json:"appID"`
	Version         string                    `json:"version"`
	SitePrefix      string                    `json:"sitePrefix"`
	Port            string                    `json:"port"`
	Session         []EnvSessionConfig        `json:"session"`
	Log             EnvLogConfig              `json:"log"`
	Db              []EnvDBConfig             `json:"db"`
	SecretsManager  []EnvSecretsManagerConfig `json:"secretsManager"`
	InternalService EnvInternalServiceConfig  `json:"internalService"`
	ExternalService EnvExternalServiceConfig  `json:"externalService"`
	App             AppConfig                 `json:"app"`
	DynamoDB        AppDynamoDBConfig         `json:"dynamoDB"`
	EndPoint        AppEndpointConfig         `json:"endpoint"`
}

// EnvSessionConfig ...
type EnvSessionConfig struct {
	DbHost            string                  `json:"dbHost"`
	DbName            string                  `json:"dbName"`
	DbUserName        string                  `json:"dbUserName"`
	DbPassword        string                  `json:"dbPassword"`
	DbDNSParam        string                  `json:"dbDnsParam"`
	DbDsn             string                  `json:"dbDns"`
	MaxIdleConns      int                     `json:"maxIdleConns"`
	MaxOpenConns      int                     `json:"maxOpenConns"`
	ConnMaxLifetime   int                     `json:"connMaxLifetime"`
	LogMode           bool                    `json:"logMode"`
	SessionID         string                  `json:"sessionID"`
	SessionResetKEY   string                  `json:"sessionResetKEY"`
	Table             string                  `json:"table"`
	SessionCookiePath string                  `json:"sessionCookiePath"`
	ExpiresInSeconds  int                     `json:"expiresInSeconds"`
	Secret            string                  `json:"secret"`
	Options           EnvSessionOptionsConfig `json:"options"`
}

// EnvSessionOptionsConfig ...
type EnvSessionOptionsConfig struct {
	Path         string `json:"path"`
	Domain       string `json:"domain"`
	MaxAge       int    `json:"maxAge"`
	Secure       bool   `json:"secure"`
	HTTPOnly     bool   `json:"httpOnly"`
	SameSiteMode string `json:"sameSiteMode"`
}

// EnvSessionOptions2Config ...
type EnvSessionOptions2Config struct {
	ExpiresOnexpiresIn int `json:"expiresOnexpiresIn"`
}

// EnvLogConfig ...
type EnvLogConfig struct {
	AplLog             EnvAplLogConfig    `json:"apllog"`
	ErrLog             EnvAplLogConfig    `json:"errlog"`
	AccessLog          EnvAccessLogConfig `json:"accesslog"`
	ExternalJournalLog EnvAplLogConfig    `json:"externaljournallog"`
	InternalJournalLog EnvAplLogConfig    `json:"internaljournallog"`
	NotificationLog    EnvAplLogConfig    `json:"notificationlog"`
}

// EnvAplLogConfig ...
type EnvAplLogConfig struct {
	Level                int    `json:"level"`
	OutputType           string `json:"outputType"`
	OutputFilePath       string `json:"outputFilePath"`
	OutputFileMaxSize    int    `json:"outputFileMaxSize"`
	OutputFileMaxBackups int    `json:"outputFileMaxBackups"`
	OutputFileMaxAge     int    `json:"outputFileMaxAge"`
	OutputFileLocalTime  bool   `json:"outputFileLocalTime"`
	OutputFileCompress   bool   `json:"outputFileCompress"`
}

// EnvAccessLogConfig ...
type EnvAccessLogConfig struct {
	Level                int    `json:"level"`
	OutputType           string `json:"outputType"`
	OutputFilePath       string `json:"outputFilePath"`
	OutputFileMaxSize    int    `json:"outputFileMaxSize"`
	OutputFileMaxBackups int    `json:"outputFileMaxBackups"`
	OutputFileMaxAge     int    `json:"outputFileMaxAge"`
	OutputFileLocalTime  bool   `json:"outputFileLocalTime"`
	OutputFileCompress   bool   `json:"outputFileCompress"`
}

// EnvDBConfig ...
type EnvDBConfig struct {
	DbHost          string `json:"dbHost"`
	DbName          string `json:"dbName"`
	DbUserName      string `json:"dbUserName"`
	DbPassword      string `json:"dbPassword"`
	DbDNSParam      string `json:"dbDnsParam"`
	DbDsn           string `json:"dbDns"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
	LogLevel        int    `json:"logLevel"`
}

// EnvSecretsManagerConfig ...
type EnvSecretsManagerConfig struct {
	SecretID string `json:"secretId"`
}

// EnvInternalServiceConfig ...
type EnvInternalServiceConfig struct {
	SvcRepayment EnvInternalServiceSvcConfig `json:"svcRepayment"`
	SvcPayment   EnvInternalServiceSvcConfig `json:"svcPayment"`
}

// EnvInternalServiceSvcConfig ...
type EnvInternalServiceSvcConfig struct {
	BasePath string `json:"basePath"`
}

// EnvExternalServiceConfig ...
type EnvExternalServiceConfig struct {
	XxxXxxxxxx EnvExternalServiceXxxXxxxxxxConfig `json:"xxxXxxxxxx"`
}

// EnvExternalServiceXxxXxxxxxxConfig ...
type EnvExternalServiceXxxXxxxxxxConfig struct {
	BasePath string `json:"basePath"`
}

// AppConfig ...
type AppConfig struct {
	Cors                AppCorsConfig                `json:"cors"`
	Request             AppRequestConfig             `json:"request"`
	SevenBank           AppSevenBankConfig           `json:"sevenBank"`
	CommonKey           AppCommonKeyConfig           `json:"commonKey"`
	RepaymentErrMessage AppRepaymentErrMessageConfig `json:"repaymentErrMessage"`
	Operator            AppOperatorConfig            `json:"operator"`
}

// AppCorsConfig is ...
type AppCorsConfig struct {
	CorsSetting                         bool     `json:"corsSetting"`
	HeaderAccessControlAllowOrigin      []string `json:"headerAccessControlAllowOrigin"`
	HeaderAccessControlAllowCredentials bool     `json:"headerAccessControlAllowCredentials"`
}

// AppRequestConfig ...
type AppRequestConfig struct {
	ContentLengthLimit string `json:"contentLengthLimit"`
}

// AppSevenBankConfig ...
type AppSevenBankConfig struct {
	ClientSecret string `json:"clientSecret"`
}

// AppCommonKeyConfig ...
type AppCommonKeyConfig struct {
	CommonKeyCryptIV  string `json:"commonKeyCryptIV"`
	CommonKeyCryptKey string `json:"commonKeyCryptKey"`
}

// AppOperatorConfig ...
type AppOperatorConfig struct {
	PwCryptIV                               string `json:"pwCryptIV"`
	PwCryptKey                              string `json:"pwCryptKey"`
	PaximumNumberLoginHistory               int    `json:"maximumNumberLoginHistory"`
	PaximumNumberLoginPasswordChangeHistory int    `json:"maximumNumberLoginPasswordChangeHistory"`
	MaximumRecordsPerPage01                 int    `json:"maximumRecordsPerPage01"`
	MaximumNumberLoginFailNumForRevoke      int    `json:"maximumNumberLoginFailNumForRevoke"`
}

// AppRepaymentErrMessageConfig ...
type AppRepaymentErrMessageConfig struct {
	Cache AppRepaymentErrMessageCacheConfig `json:"cache"`
}

// AppRepaymentErrMessageCacheConfig ...
type AppRepaymentErrMessageCacheConfig struct {
	Interval int `json:"interval"`
}

// AppEndpointConfig ...
type AppEndpointConfig struct {
	MCSOperator     string `json:"svc-operator"`
	MCSUser         string `json:"svc-user"`
	MCSCard         string `json:"svc-card"`
	MCSPayment      string `json:"svc-payment"`
	MCSClub         string `json:"svc-club"`
	MCSNotification string `json:"svc-notification"`
	MCSApplication  string `json:"svc-application"`
	FrontEnd        string `json:"frontend"`
}

// AppDynamoDBConfig ...
type AppDynamoDBConfig struct {
	TableSession       string `json:"session"`
	TableResetPassword string `json:"reset-password"`
}

var conf *ProjectConfig

// GetProjectConfig ...
func GetProjectConfig() *ProjectConfig {
	return conf
}

func init() {
	// fmt.Println("config/config.go : init : func init() start.")

	conf = &ProjectConfig{}
	var OSEnvConf OSEnvConfig
	var Appcomconf AppComConfig
	var Appenvconf AppEnvConfig

	envconfig.Process("", &OSEnvConf)
	conf.OSEnvConf = OSEnvConf

	err := json.Unmarshal([]byte(AppConfigCommon), &Appcomconf)
	if err != nil {
		panic(err)
	}
	conf.AppComConf = Appcomconf

	if OSEnvConf.ActiveEnv != "" {
		conf.AppComConf.ActiveEnv = OSEnvConf.ActiveEnv
	}

	if conf.AppComConf.ActiveEnv == "local" {
		err = json.Unmarshal([]byte(localconfig.AppConfig), &Appenvconf)
	} else if conf.AppComConf.ActiveEnv == "dev" {
		err = json.Unmarshal([]byte(devconfig.AppConfig), &Appenvconf)
	} else if conf.AppComConf.ActiveEnv == "stg" {
		err = json.Unmarshal([]byte(stgconfig.AppConfig), &Appenvconf)
	} else if conf.AppComConf.ActiveEnv == "prod" {
		err = json.Unmarshal([]byte(prodconfig.AppConfig), &Appenvconf)
	}
	if err != nil {
		panic(err)
	}
	conf.AppEnvConf = Appenvconf

	if OSEnvConf.AppID != "" {
		conf.AppEnvConf.AppID = OSEnvConf.AppID
	}

	if OSEnvConf.AppPort != "" {
		conf.AppEnvConf.Port = OSEnvConf.AppPort
	}

	if OSEnvConf.ApllogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.ApllogLevel)
		if err != nil {
			conf.AppEnvConf.Log.AplLog.Level = level
		}
	}
	if OSEnvConf.ErrlogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.ErrlogLevel)
		if err != nil {
			conf.AppEnvConf.Log.ErrLog.Level = level
		}
	}
	if OSEnvConf.AccesslogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.AccesslogLevel)
		if err != nil {
			conf.AppEnvConf.Log.AccessLog.Level = level
		}
	}
	if OSEnvConf.ExternaljournalLogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.ExternaljournalLogLevel)
		if err != nil {
			conf.AppEnvConf.Log.ExternalJournalLog.Level = level
		}
	}
	if OSEnvConf.InternaljournalLogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.InternaljournalLogLevel)
		if err != nil {
			conf.AppEnvConf.Log.InternalJournalLog.Level = level
		}
	}
	if OSEnvConf.NotificationLogLevel != "" {
		level, err := strconv.Atoi(OSEnvConf.NotificationLogLevel)
		if err != nil {
			conf.AppEnvConf.Log.NotificationLog.Level = level
		}
	}

	if OSEnvConf.DbHost != "" {
		conf.AppEnvConf.Db[0].DbHost = OSEnvConf.DbHost
	}
	if OSEnvConf.DbHostRo != "" {
		conf.AppEnvConf.Db[1].DbHost = OSEnvConf.DbHostRo
	}
	if OSEnvConf.DbName != "" {
		conf.AppEnvConf.Db[0].DbName = OSEnvConf.DbName
		conf.AppEnvConf.Db[1].DbName = OSEnvConf.DbName
	}
	if OSEnvConf.DbUser != "" {
		conf.AppEnvConf.Db[0].DbUserName = OSEnvConf.DbUser
		conf.AppEnvConf.Db[1].DbUserName = OSEnvConf.DbUser
	}
	if OSEnvConf.DbPassword != "" {
		conf.AppEnvConf.Db[0].DbPassword = OSEnvConf.DbPassword
		conf.AppEnvConf.Db[1].DbPassword = OSEnvConf.DbPassword
	}

	if OSEnvConf.DbMaxIdleConns != "" {
		dbMaxIdleConns, err := strconv.Atoi(OSEnvConf.DbMaxIdleConns)
		if err != nil {
			conf.AppEnvConf.Db[0].MaxIdleConns = dbMaxIdleConns
		}
	}
	if OSEnvConf.DbMaxOpenConns != "" {
		dbMaxOpenConns, err := strconv.Atoi(OSEnvConf.DbMaxOpenConns)
		if err != nil {
			conf.AppEnvConf.Db[0].MaxOpenConns = dbMaxOpenConns
		}
	}
	if OSEnvConf.DbConnMaxLifetime != "" {
		dbConnMaxLifetime, err := strconv.Atoi(OSEnvConf.DbConnMaxLifetime)
		if err != nil {
			conf.AppEnvConf.Db[0].ConnMaxLifetime = dbConnMaxLifetime
		}
	}

	if OSEnvConf.DbMaxIdleConnsRo != "" {
		dbMaxIdleConnsRo, err := strconv.Atoi(OSEnvConf.DbMaxIdleConnsRo)
		if err != nil {
			conf.AppEnvConf.Db[1].MaxIdleConns = dbMaxIdleConnsRo
		}
	}
	if OSEnvConf.DbMaxOpenConnsRo != "" {
		dbMaxOpenConnsRo, err := strconv.Atoi(OSEnvConf.DbMaxOpenConnsRo)
		if err != nil {
			conf.AppEnvConf.Db[1].MaxOpenConns = dbMaxOpenConnsRo
		}
	}
	if OSEnvConf.DbConnMaxLifetimeRo != "" {
		dbConnMaxLifetimeRo, err := strconv.Atoi(OSEnvConf.DbConnMaxLifetimeRo)
		if err != nil {
			conf.AppEnvConf.Db[1].ConnMaxLifetime = dbConnMaxLifetimeRo
		}
	}

	if OSEnvConf.SmSecretID != "" {
		conf.AppEnvConf.SecretsManager[0].SecretID = OSEnvConf.SmSecretID
	}

	if OSEnvConf.SvcRepaymentBasePath != "" {
		conf.AppEnvConf.InternalService.SvcRepayment.BasePath = OSEnvConf.SvcRepaymentBasePath
	}
	if OSEnvConf.SvcPaymentBasePath != "" {
		conf.AppEnvConf.InternalService.SvcPayment.BasePath = OSEnvConf.SvcPaymentBasePath
	}

	// fmt.Println("config/config.go : init : func init() end.")
}
