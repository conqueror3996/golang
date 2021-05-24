package datastore

import (
	"bff-web/config"
	"database/sql"
	"fmt"

	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/savaki/dynastore"
	"github.com/srinathgs/mysqlstore"
)

// GetSessionStore セッション用Storeを返す
func GetSessionStore() (*mysqlstore.MySQLStore, error) {
	apllog.Infof(apllogdefaultutil.LogMessageV2(fmt.Sprintf("session/database/datastore.go : GetSessionStore : func GetSessionStore() (*mysqlstore.MySQLStore, error) start.")))

	var dbDsn string

	if config.GetProjectConfig().AppEnvConf.Session[0].DbDNSParam == "" {
		dbDsn = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			config.GetProjectConfig().AppEnvConf.Session[0].DbUserName,
			config.GetProjectConfig().AppEnvConf.Session[0].DbPassword,
			config.GetProjectConfig().AppEnvConf.Session[0].DbHost,
			config.GetProjectConfig().AppEnvConf.Session[0].DbName)
	} else {
		dbDsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
			config.GetProjectConfig().AppEnvConf.Session[0].DbUserName,
			config.GetProjectConfig().AppEnvConf.Session[0].DbPassword,
			config.GetProjectConfig().AppEnvConf.Session[0].DbHost,
			config.GetProjectConfig().AppEnvConf.Session[0].DbName,
			config.GetProjectConfig().AppEnvConf.Session[0].DbDNSParam)
	}

	store, err := mysqlstore.NewMySQLStore(
		dbDsn,
		config.GetProjectConfig().AppEnvConf.Session[0].Table,
		config.GetProjectConfig().AppEnvConf.Session[0].SessionCookiePath,
		config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds,
		[]byte(config.GetProjectConfig().AppEnvConf.Session[0].Secret))
	if err != nil {
		return store, err
	}
	return store, nil
}

// GetSessionStoreFromConnection セッション用Storeを返す
func GetSessionStoreFromConnection(db *sql.DB, err error) (*mysqlstore.MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(
		db,
		config.GetProjectConfig().AppEnvConf.Session[0].Table,
		config.GetProjectConfig().AppEnvConf.Session[0].SessionCookiePath,
		config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds,
		[]byte(config.GetProjectConfig().AppEnvConf.Session[0].Secret))
	if err != nil {
		return store, err
	}
	return store, nil
}

// GetSessionStoreFromConnectionDynamoDB セッション用Storeを返す
func GetSessionStoreFromConnectionDynamoDB() (*dynastore.Store, error) {
	path := config.GetProjectConfig().AppEnvConf.Session[0].Options.Path
	maxAge := config.GetProjectConfig().AppEnvConf.Session[0].Options.MaxAge
	tableName := config.GetProjectConfig().AppEnvConf.Session[0].Table
	// os.Setenv("AWS_DEFAULT_REGION", "ap-northeast-1")
	var configAWS *awssdk.Config
	if config.GetProjectConfig().AppEnvConf.Env == "local" {
		configAWS = &awssdk.Config{
			Region:   awssdk.String("ap-northeast-1"),
			Endpoint: awssdk.String("http://localhost:8000"),
		}
	} else {
		configAWS = &awssdk.Config{}
	}

	store, err := dynastore.New(dynastore.Path(path), dynastore.HTTPOnly(), dynastore.MaxAge(maxAge), dynastore.AWSConfig(configAWS), dynastore.TableName(tableName))

	if err != nil {
		return store, err
	}
	return store, nil
}
