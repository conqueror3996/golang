package usecase

import (
	"bff-web/domain/model"
	appsesson "bff-web/session"
	"bff-web/usecase/input"
	"bff-web/util"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gorilla/sessions"

	"bff-web/config"
	//contextwbff "bff-web/contextw"
	//echocontextw "bff-web/contextw/echo"
	//"bff-web/domain/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	commonusecase "github.com/nri4nudge/api-common/usecase"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Common common related interface
type Common interface {
	InitNewSDKSession(ctx context.Context) (*awsSession.Session, error)
	CreateTable(ctx context.Context, sess *awsSession.Session, tableName string, partitionKey string) error
	InsertKey(ctx context.Context, sess *awsSession.Session, key interface{}, tableName string) error
	UpdateKey(ctx context.Context, sess *awsSession.Session, keyUpdate interface{}, dataUpdate interface{}, tableName string) error
	GetResetKey(ctx context.Context, sess *awsSession.Session, iptKey string) (*model.ResetKey, error)
	DeleteResetKey(ctx context.Context, sess *awsSession.Session, iptKey string) error
	SetSessionAuthInfo(ctx context.Context, sess *sessions.Session, id string, operator_id string, authStatus string, userStatus string, firstLogin bool, firstName string, familyName string, role string, roleDetail string, phoneNumber string) error
	SetSessionResetKey(ctx context.Context, sess *sessions.Session, id string, key string, expired int64) error
	SessionGet(ctx context.Context, c echo.Context) (*sessions.Session, error)
	SessionResetKeyGet(ctx context.Context, c echo.Context) (*sessions.Session, error)
	ValidateAuthLoginStatus(ctx context.Context, c echo.Context, authStatus string, msgCD01 string, msgCD02 string) error
	SetHeaderRequest(ctx context.Context, header *http.Request, iptMeta *model.CommonRequestMeta) error
	ValidateInputMeta(ctx context.Context, iptMeta input.MetaRequest, chkXAPIKey bool) error
	ValidateAccountStatus(ctx context.Context, strRole string, operatorID string, iptMeta input.MetaRequest, status string, errorCode string) (*model.ErrorResult, error)
}

// NewCommon generate common instance
func NewCommon(
	pre Presenter,
	errm apperror.ErrorManager,
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	commoncommon commonusecase.Common,
	operator Operator,
) Common {
	return common{
		pre,
		errm,
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		commoncommon,
		operator,
	}
}

type common struct {
	pre                                    Presenter
	errm                                   apperror.ErrorManager
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	commoncommon                           commonusecase.Common
	operator                               Operator
}

func (u common) ValidateInputMeta(ctx context.Context, iptMeta input.MetaRequest, chkXAPIKey bool) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : func (u common) ValidateInputMeta(ctx context.Context, iptMeta *model.MetaRequest) error start.")))

	if iptMeta.ReceptionProcessingStatus != "0" {
		msgCD := "BFWE00000001"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if iptMeta.HeaderContentType == "" {
		msgCD := "BFWE00000002"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !strings.Contains(iptMeta.HeaderContentType, "application/json") {
		msgCD := "BFWE00000003"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(iptMeta.HeaderContentType) > 16 {
		msgCD := "BFWE00000003"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if iptMeta.HeaderUserAgent == "" {
		msgCD := "BFWE00000004"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.ValidateUserAgent(iptMeta.HeaderUserAgent) {
		msgCD := "BFWE00000005"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if iptMeta.HeaderXRequestedWith == "" {
		msgCD := "BFWE00000006"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if iptMeta.HeaderXRequestedWith != "XMLHttpRequest" {
		msgCD := "BFWE00000007"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if iptMeta.HeaderXNudgeInteractionID == "" {
		msgCD := "BFWE00000008"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(iptMeta.HeaderXNudgeInteractionID) {
		msgCD := "BFWE00000009"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(iptMeta.HeaderXNudgeInteractionID) > 59 {
		msgCD := "BFWE00000009"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if chkXAPIKey == true {
		if iptMeta.HeaderXAPIKey == "" {
			msgCD := "BFWE00000010"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchAlphanumericSymbolStringV1(iptMeta.HeaderXAPIKey) {
			msgCD := "BFWE00000011"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if iptMeta.HeaderXAPIKey != config.GetProjectConfig().AppEnvConf.Session[0].Secret {
			msgCD := "BFWE00000011"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	return nil
}

func (u common) InitNewSDKSession(ctx context.Context) (*awsSession.Session, error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : func (u common) InitNewSDKSession(ctx context.Context) (*awsSession.Session, error) start.")))

	sess, err := awsSession.NewSession(&aws.Config{
		// Region:   aws.String("ap-northeast-1"),
		// Endpoint: aws.String("http://localhost:8000"),
		// Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})

	if err != nil {
		msgCD := "0100005001002"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return sess, nil
}

func (u common) CreateTable(ctx context.Context, sess *awsSession.Session, tableName string, partitionKey string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : CreateTable : func (u common) CreateTable(ctx context.Context, sess *awsSession.Session, tableName string) error start.")))

	svc := dynamodb.New(sess)

	inputCre := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(partitionKey),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(partitionKey),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(inputCre)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				return nil
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (u common) InsertKey(ctx context.Context, sess *awsSession.Session, key interface{}, tableName string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : InsertKey : func (u common) InsertKey(ctx context.Context, sess *awsSession.Session, key interface{}, tableName string) error start.")))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : InsertKey : tableName=%s" , tableName)))

	svc := dynamodb.New(sess)

	item, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		msgCD := "0100005001003"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		msgCD := "0100005001004"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u common) UpdateKey(ctx context.Context, sess *awsSession.Session, keyUpdate interface{}, dataUpdate interface{}, tableName string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : UpdateKey : func (u common) UpdateKey(ctx context.Context, sess *awsSession.Session, key interface{}, tableName string) error start.")))

	svc := dynamodb.New(sess)

	key, err := dynamodbattribute.MarshalMap(keyUpdate)
	if err != nil {
		msgCD := "0100005001003"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	update, err := dynamodbattribute.MarshalMap(dataUpdate)
	if err != nil {
		msgCD := "0100005001003"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	updateString := util.PrintFields(dataUpdate)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: update,
		TableName:                 aws.String(tableName),
		Key:                       key,
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          aws.String(updateString),
	}
	_, err = svc.UpdateItem(input)
	if err != nil {
		msgCD := "0100005001004"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u common) GetResetKey(ctx context.Context, sess *awsSession.Session, iptKey string) (*model.ResetKey, error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : GetResetKey : func (u common) GetResetKey(ctx context.Context, sess *awsSession.Session, iptKey string) (*model.ResetKey, error) start.")))

	svc := dynamodb.New(sess)
	table := config.GetProjectConfig().AppEnvConf.DynamoDB.TableResetPassword

	resetKey := &model.ResetKeyCondition{
		Key: iptKey,
	}
	key, err := dynamodbattribute.MarshalMap(resetKey)
	if err != nil {
		msgCD := "0100005001003"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	input2 := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(table),
	}
	result2, err := svc.GetItem(input2)
	if err != nil {
		msgCD := "0100005001005"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if result2.Item == nil {
		return nil, nil
	}

	resetKeyOutput := new(model.ResetKey)
	err = dynamodbattribute.UnmarshalMap(result2.Item, &resetKeyOutput)
	if err != nil {
		msgCD := "0100005001006"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return resetKeyOutput, nil
}

func (u common) DeleteResetKey(ctx context.Context, sess *awsSession.Session, iptKey string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : DeleteResetKey : func (u common) InsertResetKey(ctx context.Context, sess *awsSession.Session, resetKey *model.ResetKey) error start.")))

	svc := dynamodb.New(sess)
	// table := config.GetProjectConfig().AppEnvConf.Session[0].Table

	resetKey := &model.ResetKeyCondition{
		Key: iptKey,
	}

	key, err := dynamodbattribute.MarshalMap(resetKey)
	if err != nil {
		msgCD := "0100005001003"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	input4 := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String("reset_key_management"),
	}

	_, err = svc.DeleteItem(input4)
	if err != nil {
		msgCD := "0100005001007"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u common) SetHeaderRequest(ctx context.Context, req *http.Request, iptMeta *model.CommonRequestMeta) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : func (u common) SetHeaderRequest(ctx context.Context, header *http.Request, iptMeta *model.CommonRequestMeta) error start.")))

	req.Header.Set("Content-Type", iptMeta.HeaderContentType)
	req.Header.Set("User-Agent", iptMeta.HeaderUserAgent)
	req.Header.Set("X-Nudge-Interaction-Id", iptMeta.HeaderXNudgeInteractionID)
	req.Header.Set("X-Nudge-User-Id", iptMeta.HeaderXNudgeUserID)
	req.Header.Set("X-Nudge-Security-Request", iptMeta.HeaderXNudgeSecurityRequest)

	return nil
}

func (u common) SetSessionAuthInfo(ctx context.Context, sess *sessions.Session, id string, operator_id string, authStatus string, userStatus string, firstLogin bool, firstName string, familyName string, role string, roleDetail string, phoneNumber string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : func (u common) SetSessionAuthInfo(ctx context.Context, sess *sessions.Session, id string, authStatus string, userStatus string, firstLogin bool, firstName string, familyName string, role string, roleDetail string, phoneNumber string) error start.")))

	sess.Values[appsesson.SessionVarOperatorAuthID] = id
	sess.Values[appsesson.SessionVarOperatorID] = operator_id
	sess.Values[appsesson.SessionVarOperatorAuthStatus] = authStatus
	sess.Values[appsesson.SessionValAccountStatus] = userStatus
	sess.Values[appsesson.SessionValFirstLoginFlg] = firstLogin
	sess.Values[appsesson.SessionValOperatorFirstName] = firstName
	sess.Values[appsesson.SessionValOperatorFamilyName] = familyName
	sess.Values[appsesson.SessionValOperatorRole] = role
	sess.Values[appsesson.SessionValOperatorRoleDetail] = roleDetail
	sess.Values[appsesson.SessionValOperatorPhoneNumber] = phoneNumber

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionVarOperatorAuthID]=%s.", id)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionVarOperatorID]=%s.", operator_id)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionVarOperatorAuthStatus]=%s.", authStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValAccountStatus]=%s.", userStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValFirstLoginFlg]=%t.", firstLogin)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValOperatorFirstName]=%s.", firstName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValOperatorFamilyName]=%s.", familyName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValOperatorRole]=%s.", role)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValOperatorRoleDetail]=%s.", roleDetail)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionAuthInfo : sess.Values[appsesson.SessionValOperatorPhoneNumber]=%s.", phoneNumber)))
	return nil
}

func (u common) SetSessionResetKey(ctx context.Context, sess *sessions.Session, id string, key string, expired int64) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SetSessionResetKey : func (u common) SetSessionResetKey(ctx context.Context, sess *sessions.Session, id string, key string, expired int64) error start.")))

	sess.Values[appsesson.SessionVarOperatorAuthID] = id
	sess.Values[appsesson.SessionValKey] = key
	sess.Values[appsesson.SessionValExpired] = expired

	return nil
}

func (u common) SessionGet(ctx context.Context, c echo.Context) (*sessions.Session, error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SessionGet : func (u common) SessionGet(ctx context.Context, c echo.Context) (*sessions.Session, error) start.")))

	sess, err := session.Get(config.GetProjectConfig().AppEnvConf.Session[0].SessionID, c)
	if err != nil {
		msgCD := "010100015001001"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	sess.Options = &sessions.Options{
		Path:     config.GetProjectConfig().AppEnvConf.Session[0].Options.Path,
		Domain:   config.GetProjectConfig().AppEnvConf.Session[0].Options.Domain,
		MaxAge:   config.GetProjectConfig().AppEnvConf.Session[0].Options.MaxAge,
		Secure:   config.GetProjectConfig().AppEnvConf.Session[0].Options.Secure,
		HttpOnly: config.GetProjectConfig().AppEnvConf.Session[0].Options.HTTPOnly,
	}
	if config.GetProjectConfig().AppEnvConf.Session[0].Options.SameSiteMode == "None" {
		sess.Options.SameSite = http.SameSiteNoneMode
	}

	return sess, nil
}

func (u common) SessionResetKeyGet(ctx context.Context, c echo.Context) (*sessions.Session, error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : SessionResetKeyGet : func (u common) SessionResetKeyGet(ctx context.Context, c echo.Context) (*sessions.Session, error) start.")))

	sess, err := session.Get(config.GetProjectConfig().AppEnvConf.Session[0].SessionResetKEY, c)
	if err != nil {
		msgCD := "010100015001001"
		return nil, u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	sess.Options = &sessions.Options{
		Path:     config.GetProjectConfig().AppEnvConf.Session[0].Options.Path,
		Domain:   config.GetProjectConfig().AppEnvConf.Session[0].Options.Domain,
		MaxAge:   config.GetProjectConfig().AppEnvConf.Session[0].Options.MaxAge,
		Secure:   config.GetProjectConfig().AppEnvConf.Session[0].Options.Secure,
		HttpOnly: config.GetProjectConfig().AppEnvConf.Session[0].Options.HTTPOnly,
	}
	if config.GetProjectConfig().AppEnvConf.Session[0].Options.SameSiteMode == "None" {
		sess.Options.SameSite = http.SameSiteNoneMode
	}

	return sess, nil
}

func (u common) ValidateAuthLoginStatus(ctx context.Context, c echo.Context, authStatus string, msgCD01 string, msgCD02 string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateAuthLoginStatus : func (u common) ValidateAuthLoginStatus(ctx context.Context, c echo.Context) error start.")))

	sess, err := session.Get(config.GetProjectConfig().AppEnvConf.Session[0].SessionID, c)
	if err != nil {
		msgCD := "0100005001001"
		return u.errm.Wrap1(
			err,
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	if sess.Values[appsesson.SessionVarOperatorAuthStatus] == nil {
		//msgCD := "BFWE01010405"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD01,
			"",
		)
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateAuthLoginStatus : sess.Values[appsesson.SessionVarOperatorAuthStatus]=%s", sess.Values[appsesson.SessionVarOperatorAuthStatus].(string))))

	if sess.Values[appsesson.SessionVarOperatorAuthStatus].(string) != authStatus {
		//msgCD := "BFWE01010406"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD02,
			"",
		)
	}

	return nil
}

func (u common) ValidateAccountStatus(ctx context.Context, strRole string, operatorID string, iptMeta input.MetaRequest, status string, errorCode string) (*model.ErrorResult, error) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateAccountStatus : func (u common) ValidateAccountStatus(ctx context.Context, strRole string, operatorID string, iptMeta input.MetaRequest, status string, errorCode string) (*model.ErrorResult, error) start.")))

	client := &http.Client{}
	// 3.5.2 アカウントステータスチェック
	// 3.5.2.1 アカウントステータス取得
	var operatorRole model.OperatorRoleInfo
	json.Unmarshal([]byte(strRole), &operatorRole)
	apiURLCheck := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/operator/list"
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateAccountStatus : urlOperatorList=%s", apiURLCheck)))
	var jsonBodyStrCheck []byte
	var err error
	if status == "search" {
		operatorRole.OperatorRole = "04" // セッションから取得したロール内の'operator_role'を'04'にして設定
		OperatorListRequestPattern2 := &input.OperatorListRequestToSVCPattern2{
			SearchText: operatorID, // セッションから取得したIDを設定
			SortKey:    "04",       // '04:オペレータID'を固定で設定
			SortOrder:  "asc",      // 'asc'を固定で設定
			PageNum:    "1",        // '1'を固定で設定
			PageSize:   "1",        // '1'を固定で設定
			Role:       input.Role(operatorRole),
		}
		jsonBodyStrCheck, err = json.Marshal(OperatorListRequestPattern2)

		if err != nil {
			return nil, err
		}
	}

	if status == "detail" {
		//operatorRole.OperatorRole = "04" // セッションから取得したロール内の'operator_role'を'04'にして設定
		OperatorListRequestToSVCPattern1 := &input.OperatorListRequestToSVCPattern1{
			OperatorID: operatorID, // セッションから取得したIDを設定
			Role:       input.Role(operatorRole),
		}
		jsonBodyStrCheck, err = json.Marshal(OperatorListRequestToSVCPattern1)

		if err != nil {
			return nil, err
		}
	}
	reqCheck, err := http.NewRequest("POST", apiURLCheck, bytes.NewBuffer(jsonBodyStrCheck))
	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          "1",
		HeaderXNudgeSecurityRequest: "1",
	}
	err = u.SetHeaderRequest(ctx, reqCheck, iptRequestMeta)
	if err != nil {
		return nil, err
	}

	resCheck, err := client.Do(reqCheck)
	if err != nil {
		panic(err)
	}
	defer resCheck.Body.Close()
	bodyCheck, _ := ioutil.ReadAll(resCheck.Body)
	dataCheck := new(model.OperatorListSearchResponse)

	if resCheck.StatusCode == 200 {
		err = json.Unmarshal(bodyCheck, &dataCheck)
		if err != nil {
			return nil, err
		}
		// 3.5.2.2 アカウントステータスチェック
		// 3.5.2.1で取得したデータからaccount_statusを取得し、チェック仕様No.3-3を実施する
		if dataCheck.Result.OperatorList[0].AccountStatus != "01" {
			msgCD := errorCode
			err = u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)

			return nil, err
		}
	} else {
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(bodyCheck, &errorResult)
		if err != nil {
			return nil, err
		}
		return errorResult, nil
	}

	return nil, nil
}
