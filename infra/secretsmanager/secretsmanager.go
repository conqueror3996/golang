package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"

	apperrorsub "github.com/nri4nudge/api-common/adapter/error"
	"github.com/nri4nudge/api-common/contextw"
	apperror "github.com/nri4nudge/api-common/infra/error"
)

var errm apperror.ErrorManager

func init() {
	errm = apperrorsub.NewErrorManager()
}

func newSecretsmanagerSession() (*session.Session, error) {
	return (*sm).NewSecretsmanagerSession()
}

// GetSecret ...
func GetSecret(ctx context.Context, sec string) (map[string]interface{}, error) {

	smSession, err := newSecretsmanagerSession()
	if err != nil {
		return nil, errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			"250000500900402",
			"",
		)
	}

	svc := secretsmanager.New(smSession,
		aws.NewConfig().WithRegion("ap-northeast-1"))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(sec),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	secretString := aws.StringValue(result.SecretString)
	// secretString := "{\"clientSecret\":\"012345678901234567890\",\"dummykey1\":\"dummyvalue1\"}"

	resmap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(secretString), &resmap); err != nil {
		return nil, errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			"250000500900212",
			"",
		)
	}

	return resmap, nil
}
