package secretsmanager

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sm *SecretsManager

// init initialize
func init() {
	var sma SecretsManager
	sma = &asm{}
	sm = &sma
}

// GetSecretsManager ...
func GetSecretsManager() *SecretsManager {
	return sm
}

// SecretsManager ...
type SecretsManager interface {
	NewSecretsmanagerSession() (*session.Session, error)
}

type asm struct {
}

func (m *asm) NewSecretsmanagerSession() (*session.Session, error) {
	var err error
	var smSession *session.Session

	smSession, err = session.NewSession(
		&aws.Config{
			Retryer: CustomRetryer{
				DefaultRetryer: client.DefaultRetryer{
					// 最大再試行回数を設定します
					// NumMaxRetries: client.DefaultRetryerMaxNumRetries, // DefaultRetryerMaxNumRetries = 3
					NumMaxRetries: 3,
					// 最小再試行遅延を設定します
					// MinRetryDelay: client.DefaultRetryerMinRetryDelay, // DefaultRetryerMinRetryDelay = 30 * time.Millisecond
					MinRetryDelay: 30 * time.Millisecond,
					// スロットル時の最小遅延を設定します
					// MinThrottleDelay: client.DefaultRetryerMinThrottleDelay, // DefaultRetryerMinThrottleDelay = 500 * time.Millisecond
					MinThrottleDelay: 500 * time.Millisecond,
					// 最大再試行遅延を設定します
					// MaxRetryDelay: client.DefaultRetryerMaxRetryDelay, // DefaultRetryerMaxRetryDelay = 300 * time.Second
					MaxRetryDelay: 20 * time.Second,
					// スロットル時の最大遅延を設定します
					// MaxThrottleDelay: client.DefaultRetryerMaxThrottleDelay, // DefaultRetryerMaxThrottleDelay = 300 * time.Second
					MaxThrottleDelay: 20 * time.Second,
				}},
		})
	if err != nil {
		return nil, err
	}

	return smSession, nil
}

// CustomRetryer ...
type CustomRetryer struct {
	client.DefaultRetryer
}
