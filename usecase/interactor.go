package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	//commonbff "bff-web/contextw"
	// "github.com/nri4nudge/api-common/config"
	"github.com/nri4nudge/api-common/contextw"
	// echocontextw "github.com/nri4nudge/api-common/contextw/echo"
	"github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"

	// "github.com/nri4nudge/api-common/logger"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"

	// "github.com/nri4nudge/api-common/logger/externaljournallog"
	// externaljournallogutil "github.com/nri4nudge/api-common/logger/externaljournallog/util"
	// "github.com/nri4nudge/api-common/logger/internaljournallog"
	// internaljournallogutil "github.com/nri4nudge/api-common/logger/internaljournallog/util"

	"bff-web/config"
	"bff-web/domain/model"
	"bff-web/usecase/input"

	appsesson "bff-web/session"

	commonusecase "github.com/nri4nudge/api-common/usecase"

	"bff-web/util"
)

// NewInteractor new Interactor
func NewInteractor(
	pre Presenter,
	errm apperror.ErrorManager,
	commonTransactionRepository repository.TransactionRepository,
	commonReadReplicaTransactionRepository repository.TransactionRepository,
	commoncommon commonusecase.Common,
	common Common,
	operator Operator,

) Interactor {
	return Interactor{
		pre,
		errm,
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		commoncommon,
		common,
		operator,
	}
}

// Interactor usecase interactor
type Interactor struct {
	pre                                    Presenter
	errm                                   apperror.ErrorManager
	commonTransactionRepository            repository.TransactionRepository
	commonReadReplicaTransactionRepository repository.TransactionRepository
	commoncommon                           commonusecase.Common
	common                                 Common
	operator                               Operator
}

// AppV1Healthcheck ...
func (i Interactor) AppV1Healthcheck(c echo.Context, ipt input.Healthcheck) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1Healthcheck : func (i Interactor) AppV1Healthcheck(c echo.Context, ipt input.Healthcheck) start.")

	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "API259001")

	// ***********************************************
	// 応答電文返却
	// ***********************************************
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1Healthcheck : i.pre.ViewHealthcheck call.")
	i.pre.ViewHealthcheck(c)
}

// AppV1AuthLogin ...
func (i Interactor) AppV1AuthLogin(c echo.Context, ipt input.AuthLogin, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : func (i Interactor) AppV1AuthLogin(c echo.Context, ipt input.AuthLogin, iptMeta input.MetaRequest) start.")))
	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010101")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.ReceptionProcessingStatus=%s", iptMeta.ReceptionProcessingStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : iptMeta.HeaderAuthorization=%s", iptMeta.HeaderAuthorization)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : ipt.ID=%s", ipt.ID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : ipt.Password=%s", ipt.Password)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : ipt.SmsCode=%s", ipt.SmsCode)))

	// 	チェック仕様No1-1～1-8を実施する（要求電文内で要確認の項目があれば、更新する）
	//     ※1-9～1-11は属性チェックのみ。項目自体存在しないパラメータはスキップする

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.operator.ValidateInputMeta(ctx, iptMetaRequest) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, false)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.operator.ValidateAppV1AuthLogin(ctx, ipt) call.")))
	err = i.operator.ValidateAppV1AuthLogin(ctx, ipt)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 3.3 サービス利用可否チェック
	// チェック仕様No.2-1～2-2を実施する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.SessionGet(ctx, c) call.")))
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	actionLogin := "login"
	//リクエスト種別判断
	authLoginResponse := new(model.AuthLoginResponse)
	authLoginSMSResponse := new(model.AuthLoginSMSResponse)
	// パスワード認証を実行する
	if (ipt.ID != "" && ipt.Password != "") && ipt.SmsCode == "" {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.SetSessionAuthInfo call.")))
		err = i.common.SetSessionAuthInfo(ctx, sess, "", "", appsesson.SessionValOperatorAuthStatus0, "", false, "", "", "", "", "")
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		// 3.5.1 「オペレータ認証API」呼び出し
		// /v1/auth/login
		url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/auth/login"
		jsonBodyLogin, err := json.Marshal(ipt)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyLogin))
		// Set header request
		iptRequestMeta := &model.CommonRequestMeta{
			HeaderContentType:           iptMeta.HeaderContentType,
			HeaderUserAgent:             iptMeta.HeaderUserAgent,
			HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
			HeaderXNudgeUserID:          "1",
			HeaderXNudgeSecurityRequest: "1",
		}
		err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		// 応答チェック
		if res.StatusCode == 200 { // A）正常の場合
			data := new(model.OperatorInfo)
			err = json.Unmarshal(body, &data)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}

			//3.5.2 「sendOTPAPI」呼び出し
			url = config.GetProjectConfig().AppEnvConf.EndPoint.MCSUser + "/v1/user/send_otp"
			var jsonBodySendOTP = []byte(`{"userID": "` + data.OperatorID + `", "phoneNumber": "` + data.PhoneNumber + `"}`)
			reqOTP, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodySendOTP))
			reqOTP.Header.Set("Content-Type", iptMeta.HeaderContentType)
			reqOTP.Header.Set("User-Agent", fmt.Sprintf("%s/%s", config.GetProjectConfig().AppEnvConf.AppID, config.GetProjectConfig().AppEnvConf.Version))
			reqOTP.Header.Set("X-Nudge-Interaction-Id", iptMeta.HeaderXNudgeInteractionID)
			reqOTP.Header.Set("X-Nudge-User-Id", "operator_id")
			client = &http.Client{}
			resOTP, err := client.Do(reqOTP)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
			defer resOTP.Body.Close()
			body, _ = ioutil.ReadAll(resOTP.Body)
			if resOTP.StatusCode == 200 { // A）　正常の場合
				// 3.5.4 セッション初期化
				jsonRoleStr, err := json.Marshal(data.Role)
				if err != nil {
					http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
				}
				name := strings.Split(data.Username, " ")
				familyName := ""
				firstName := ""
				if len(name) > 0 {
					firstName = name[0]
					if len(name) > 1 {
						for i := 1; i < len(name); i++ {
							familyName += name[i] + " "
						}
					}
					familyName = strings.Trim(familyName, " ")
				}

				apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.SetSessionAuthInfo call.")))
				err = i.common.SetSessionAuthInfo(
					ctx,
					sess,
					ipt.ID,
					data.OperatorID,                         // request.id
					appsesson.SessionValOperatorAuthStatus2, // 2
					data.AccountStatus,                      // operatorAuthApi.response.user_status
					data.FirstLoginFlg,                      // operatorAuthApi.response.first_login
					firstName,                               // firstName
					familyName,                              // familyName
					data.OperatorRoleID,
					string(jsonRoleStr), // operatorAuthApi.response.role
					data.PhoneNumber,    // operatorAuthApi.response.phone_number
				)
				if err != nil {
					i.pre.ViewError(c, err)
					return
				}

				apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : sess.Values[appsesson.SessionVarOperatorAuthStatus]=" + sess.Values[appsesson.SessionVarOperatorAuthStatus].(string))))

				apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.SetCtxUserInfo call.")))
				ctx, err = i.commoncommon.SetCtxUserInfo(ctx, c, ipt.ID)
				if err != nil {
					i.pre.ViewError(c, err)
					return
				}
			} else { // B）　異常の場合
				errorResult := new(model.ErrorResult)
				err = json.Unmarshal(body, &errorResult)
				if err != nil {
					i.pre.ViewError(c, err)
					return
				}
				i.pre.ViewErrorMCS(c, errorResult)
				return
			}

		} else { // B）　異常の場合
			errorResult := new(model.ErrorResult)
			err = json.Unmarshal(body, &errorResult)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
			i.pre.ViewErrorMCS(c, errorResult)
			return
		}

		sess, err = i.common.SessionGet(ctx, c)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}

		//3.5.7 応答電文返却
		// 応答パターン1:個別設定事項
		// login_status:1を設定する。
		authLoginResponse = &model.AuthLoginResponse{
			LoginStatus: 1,
		}
	} else if ipt.ID == "" && ipt.Password == "" && ipt.SmsCode != "" { //3.6 SMS認証
		// 3.4.2 リクエスト項目が「sms_code」のみの場合
		// SMS認証を実行する

		// 3.6.2 認証ステータスチェック
		// インスタンス 5.1: SESSION.auth_statusが未設定の場合エラー
		// インスタンス 5.2: SESSION.auth_statusが3以外の場合エラー
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus2, 'BFWE01010106', 'BFWE01010107') call.")))
		err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus2, "BFWE01010106", "BFWE01010107")
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		// セッションの"id"
		sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)
		// sessionPhoneNumber := sess.Values[appsesson.SessionValOperatorPhoneNumber].(string)
		sessionOperatorID := sess.Values[appsesson.SessionVarOperatorID].(string)

		// 3.6.3 「verifyOTPAPI」呼び出し
		url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSUser + "/v1/user/verify_otp"
		// var jsonBodyStr = []byte(`{"phone_number": "` + sessionPhoneNumber + `","sms_code": "` + ipt.SmsCode + `"}`)
		var jsonBodyStr = []byte(`{"userID": "` + sessionOperatorID + `","code": "` + ipt.SmsCode + `"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyStr))
		req.Header.Set("Content-Type", iptMeta.HeaderContentType)
		req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", config.GetProjectConfig().AppEnvConf.AppID, config.GetProjectConfig().AppEnvConf.Version))
		req.Header.Set("X-Nudge-Interaction-Id", iptMeta.HeaderXNudgeInteractionID)
		req.Header.Set("X-Nudge-User-Id", "operator_id")
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode == 200 { // A）　正常の場合
			// START 3.6.4 セッション更新
			// セッション更新
			if sess.Values[appsesson.SessionValFirstLoginFlg].(bool) == true {
				sess.Values[appsesson.SessionVarOperatorAuthStatus] = appsesson.SessionValOperatorAuthStatus3
			} else {
				sess.Values[appsesson.SessionVarOperatorAuthStatus] = appsesson.SessionValOperatorAuthStatus1
			}
			// 3.6.6 応答電文返却
			// 	A) セッション中の"auth_status"が"3:初回ログイン"の場合
			loginStatus := 0
			if sess.Values[appsesson.SessionVarOperatorAuthStatus] == appsesson.SessionValOperatorAuthStatus3 {
				// 応答パターン3:個別設定事項
				// 	login_status:2を設定する。
				loginStatus = 2
			} else if sess.Values[appsesson.SessionVarOperatorAuthStatus] == appsesson.SessionValOperatorAuthStatus3 {
				//  B) セッション中の"auth_status"が"1:ログイン"の場合
				//       応答パターン:2:個別設定事項
				//       login_status:0を設定する。
				loginStatus = 0
			}
			var operatorRole model.OperatorRoleInfo
			strRole := sess.Values[appsesson.SessionValOperatorRoleDetail].(string)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
			json.Unmarshal([]byte(strRole), &operatorRole)
			roleResponse := &model.RoleResponse{
				CardRole:     operatorRole.CardMemberRole,
				PartnerRole:  operatorRole.ClubOwnerRole,
				SkinRole:     operatorRole.ClubRole,
				OperatorRole: operatorRole.OperatorRole,
				MessageRole:  operatorRole.MessageRole,
			}
			userName := sess.Values[appsesson.SessionValOperatorFamilyName].(string) + " " + sess.Values[appsesson.SessionValOperatorFirstName].(string)
			authLoginSMSResponse = &model.AuthLoginSMSResponse{
				LoginStatus: loginStatus,
				UserName:    userName,
				Role:        *roleResponse,
			}
			apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.SetSessionAuthInfo call.")))
			err = i.common.SetSessionAuthInfo(
				ctx,
				sess,
				sessionID,
				sess.Values[appsesson.SessionVarOperatorID].(string),
				sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
				sess.Values[appsesson.SessionValAccountStatus].(string),
				sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
				sess.Values[appsesson.SessionValOperatorFirstName].(string),
				sess.Values[appsesson.SessionValOperatorFamilyName].(string),
				sess.Values[appsesson.SessionValOperatorRole].(string),
				sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
				sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
			)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}

			actionLogin = "sms"
			// END 3.6.4 セッション更新
		} else { // B）　異常の場合
			errorResult := new(model.ErrorResult)
			err = json.Unmarshal(body, &errorResult)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
			i.pre.ViewErrorMCS(c, errorResult)
			return
		}
	}
	// セッション保存※DynamoDBに登録
	expiredOn := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))
	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expiredOn
	sess.Save(c.Request(), c.Response())

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	if actionLogin == "login" {
		// // 3.6.1 パスワード初期化用設定
		session := &model.Session{
			SessionID:   sess.ID,
			ID:          ipt.ID,
			OperatorID:  sess.Values[appsesson.SessionVarOperatorID].(string),
			AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			UserStatus:  sess.Values[appsesson.SessionValAccountStatus].(string),
			FirstLogin:  sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
			FirstName:   sess.Values[appsesson.SessionValOperatorFirstName].(string),
			FamilyName:  sess.Values[appsesson.SessionValOperatorFamilyName].(string),
			Role:        sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
			PhoneNumber: sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
			Expired:     expiredOn.Unix(),
			Flag:        "0",
			CreateDate:  util.TimeFormatMillisecondV1(util.NowTimeV1()),
			UpdatedDate: "",
			CreatePgmID: contextw.GetAPIID(ctx),
			UpdatePgmID: "",
		}

		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.InsertKey(ctx, awsSess, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
		err = i.common.InsertKey(ctx, awsSess, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.pre.ViewAppV1AuthLogin(c) call.")))
		i.pre.ViewAppV1AuthLogin(c, authLoginResponse)
	} else {
		session := &model.SessionUpdate{
			AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			Expired:     expiredOn.Unix(),
			UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
			UpdatePgmID: contextw.GetAPIID(ctx),
		}
		keyUpdate := &model.UpdateSessionCondition{
			SessionID: sess.ID,
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
		err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthLogin : i.pre.ViewAppV1AuthLoginSMS(c) call.")))
		i.pre.ViewAppV1AuthLoginSMS(c, authLoginSMSResponse)
	}

}

// AppV1AuthReset ...
func (i Interactor) AppV1AuthReset(c echo.Context, ipt input.AuthReset, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : func (i Interactor) AppV1AuthReset(c echo.Context, ipt input.AuthLogin, iptMeta input.MetaRequest) start.")))
	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010103")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.ReceptionProcessingStatus=%s", iptMeta.ReceptionProcessingStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : ipt.Key=%s", ipt.Key)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : ipt.Password=%s", ipt.Password)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, false)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 	チェック仕様No1-1～1-15を実施する（要求電文内で要確認の項目があれば、更新する）
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.operator.ValidateAppV1AuthReset(ctx, ipt, iptMeta) call.")))
	err = i.operator.ValidateAppV1AuthReset(ctx, ipt)
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 3.3 サービス利用可否チェック
	// チェック仕様No.2-1～2-2を実施する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 3.4 key検索
	// リクエスト項目「key」を元にDynamoDBからデータ取得する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.GetResetKey(ctx, awsSess, ipt.Key) call.")))
	resetKey, err := i.common.GetResetKey(ctx, awsSess, ipt.Key)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// A) 結果が取得でき、かつ現在時刻が取得した項目の「expired※unixtime表記」の時刻以内の場合
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.ValidateExpiredAuthReset(ctx, resetKey) call.")))
	err = i.operator.ValidateExpiredAuthReset(ctx, resetKey)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// /v1/auth/update
	url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/auth/update"
	// セッションの"id"
	sessionID := resetKey.ID
	iptToSVC := &input.AuthChangePasswordRequestToSVC{
		MailAddress: sessionID,    // セッションの"id"
		Password:    ipt.Password, // request.password
	}
	jsonBodyPassword, err := json.Marshal(iptToSVC)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyPassword))
	// Set header request
	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          "1",
		HeaderXNudgeSecurityRequest: "1",
	}
	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// A）正常の場合
	if res.StatusCode == 200 {

		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.SetCtxUserInfo call.")))
		ctx, err = i.commoncommon.SetCtxUserInfo(ctx, c, sessionID)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}

		resetKey := &model.ResetKeyCondition{
			Key: resetKey.Key,
		}

		resetKeyUpdate := &model.ResetKeyUpdate{
			Flag:        "1",
			UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
			UpdatePgmID: contextw.GetAPIID(ctx),
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.common.UpdateKey(ctx, awsSess, resetKey, resetKeyUpdate, config.GetProjectConfig().AppEnvConf.DynamoDB.TableResetPassword) call.")))
		err := i.common.UpdateKey(ctx, awsSess, resetKey, resetKeyUpdate, config.GetProjectConfig().AppEnvConf.DynamoDB.TableResetPassword)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
	} else {
		// B）　異常の場合
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// セッション保存※DynamoDBに登録

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthReset : i.pre.ViewAppV1OperatorsOperatorIDPassword(c) call.")))
	i.pre.ViewAppV1AuthReset(c)
}

// AppV1AuthChangePassword ...
func (i Interactor) AppV1AuthChangePassword(c echo.Context, ipt input.AuthChangePassword, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : func (i Interactor) AppV1AuthChangePassword(c echo.Context, ipt input.AuthChangePassword, iptMeta input.MetaRequest) start.")))

	ctx := c.Request().Context()

	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010104")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.ReceptionProcessingStatus=%s", iptMeta.ReceptionProcessingStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : iptMeta.HeaderAuthorization=%s", iptMeta.HeaderAuthorization)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : ipt.NewPassword=%s", ipt.Password)))

	// 3.2. 要求電文チェック
	//  チェック仕様No1-1～1-13を実施する（要求電文内で要確認の項目があれば、更新する）

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, true)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.operator.ValidateAppV1AuthChangePassword(ctx, ipt) call.")))
	err = i.operator.ValidateAppV1AuthChangePassword(ctx, ipt)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 3.3 サービス利用可否チェック
	// チェック仕様No.2-1～2-2を実施する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.common.SessionGet(ctx, c) call.")))
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 3.5 ログインステータスチェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus3, 'BFWE01010405', 'BFWE01010406') call.")))
	// インスタンス 3.1: SESSION.auth_statusが未設定の場合エラー
	// インスタンス 3.2: SESSION.auth_statusが3以外の場合エラー
	err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus3, "BFWE01010405", "BFWE01010406")
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}
	authChangePasswordResponse := new(model.AuthChangePasswordResponse)
	//「パスワード変更API」を呼び出す
	// /v1/auth/update
	url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/auth/update"
	// セッションの"id"
	sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)
	iptToSVC := &input.AuthChangePasswordRequestToSVC{
		MailAddress: sessionID,    // セッションの"id"
		Password:    ipt.Password, // request.password
	}
	jsonBodyStr, err := json.Marshal(iptToSVC)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyStr))

	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          sessionID,
		HeaderXNudgeSecurityRequest: "1",
	}
	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// A）　正常の場合
	if res.StatusCode == 200 {
		// data := new(model.AuthLogin)
		// err = json.Unmarshal(body, &data)
		// if err != nil {
		// 	i.pre.ViewError(c, err)
		// 	return
		// }

		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.common.SetCtxUserInfo call.")))
		ctx, err = i.commoncommon.SetCtxUserInfo(ctx, c, sessionID)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		// 3.6 SESSION更新
		// ・セッション中の"auth_status"を"1:ログイン"に設定する。
		sess.Values[appsesson.SessionVarOperatorAuthStatus] = appsesson.SessionValOperatorAuthStatus1
		// ・セッション中の"first_login"を"0:ログイン済み"に設定する。
		sess.Values[appsesson.SessionValFirstLoginFlg] = false
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.common.SetSessionAuthInfo call.")))
		err = i.common.SetSessionAuthInfo(
			ctx,
			sess,
			sessionID,
			sess.Values[appsesson.SessionVarOperatorID].(string),
			sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			sess.Values[appsesson.SessionValAccountStatus].(string),
			sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
			sess.Values[appsesson.SessionValOperatorFirstName].(string),
			sess.Values[appsesson.SessionValOperatorFamilyName].(string),
			sess.Values[appsesson.SessionValOperatorRole].(string),
			sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
			sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
		)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : sess.Values[appsesson.SessionVarOperatorAuthStatus]=" + sess.Values[appsesson.SessionVarOperatorAuthStatus].(string))))
		// 3.8 応答電文返却
		// 応答パターン4:個別設定事項
		var operatorRole model.OperatorRoleInfo
		strRole := sess.Values[appsesson.SessionValOperatorRoleDetail].(string)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		json.Unmarshal([]byte(strRole), &operatorRole)
		RoleChangeResponse := &model.RoleChangeResponse{
			CardMemberRole: operatorRole.CardMemberRole,
			ClubOwnerRole:  operatorRole.ClubOwnerRole,
			ClubRole:       operatorRole.ClubRole,
			OperatorRole:   operatorRole.OperatorRole,
			MessageRole:    operatorRole.MessageRole,
		}
		authChangePasswordResponse = &model.AuthChangePasswordResponse{
			LoginStatus: 0,
			UserName:    sess.Values[appsesson.SessionValOperatorFirstName].(string),
			Role:        *RoleChangeResponse,
		}
	} else {
		// B）　異常の場合
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// セッション保存※DynamoDBに登録4
	expired := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))

	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expired
	sess.Save(c.Request(), c.Response())

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	session := &model.SessionUpdate{
		AuthStatus:    sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
		Expired:       expired.Unix(),
		FirstLoginFlg: sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
		UpdatedDate:   util.TimeFormatMillisecondV1(util.NowTimeV1()),
		UpdatePgmID:   contextw.GetAPIID(ctx),
	}
	keyUpdate := &model.UpdateSessionCondition{
		SessionID: sess.ID,
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.common.InsertKey(ctx, awsSess, resetKey, 'nudge-dynamodb-bff-web-reset-password') call.")))
	err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthChangePassword : i.pre.ViewAppV1OperatorsOperatorIDPassword(c) call.")))
	i.pre.ViewAppV1AuthChangePassword(c, authChangePasswordResponse)
}

// AppV1AuthResetRequest ...
func (i Interactor) AppV1AuthResetRequest(c echo.Context, ipt input.ResetRequest, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : func (i Interactor) AppV1AuthResetRequest(c echo.Context, ipt input.AuthLogin, iptMeta input.MetaRequest) start.")))
	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010102")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : ipt.mail_addr=%s", ipt.MailAddr)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, false)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 要求電文チェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.operator.ValidateAppV1AuthResetRequest(ctx, ipt) call.")))
	err = i.operator.ValidateAppV1AuthResetRequest(ctx, ipt)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// サービス利用可否チェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	// 3.6.1 パスワード初期化用設定
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	// パスワード再設定URL発行
	uuid := util.UUIDV2()
	url := config.GetProjectConfig().AppEnvConf.EndPoint.FrontEnd + "/password-reset?key=" + uuid

	// 管理者パスワード初期化APIリクエスト
	apiURL := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/auth/init"
	// jsonBodyStr := []byte(`{"mail_addr": "` + ipt.MailAddr + `","url": "` + url + `"}`)
	iptToSVC := &input.ResetRequestToSVC{
		MailAddress: ipt.MailAddr,
		URL:         url,
	}
	jsonBodyStr, err := json.Marshal(iptToSVC)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBodyStr))
	// Set header request
	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          "1",
		HeaderXNudgeSecurityRequest: "1",
	}
	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// パスワード初期化用設定保存
	if res.StatusCode == 200 {
		expired := time.Now().Add(time.Hour * 1).Unix()
		resetKey := &model.ResetKey{
			ID:          ipt.MailAddr,
			Key:         uuid,
			Flag:        "0",
			Expired:     expired,
			CreateDate:  util.TimeFormatMillisecondV1(util.NowTimeV1()),
			UpdatedDate: "",
			CreatePgmID: contextw.GetAPIID(ctx),
			UpdatePgmID: "",
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.common.InsertKey(ctx, awsSess, resetKey, config.GetProjectConfig().AppEnvConf.DynamoDB.TableResetPassword) call.")))
		err := i.common.InsertKey(ctx, awsSess, resetKey, config.GetProjectConfig().AppEnvConf.DynamoDB.TableResetPassword)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
	} else {
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// 応答電文を生成し、クライアントへ返却する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1AuthResetRequest : i.pre.ViewAppV1AuthResetRequest(c) call.")))
	i.pre.ViewAppV1AuthResetRequest(c)
}

// AppV1OperatorRegistration ...
func (i Interactor) AppV1OperatorRegistration(c echo.Context, ipt input.OperatorRegistration, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : func (i Interactor) AppV1OperatorRegistration(c echo.Context, ipt input.OperatorRegistration, iptMeta input.OperatorRegistrationMeta) start.")))
	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010501")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.FirstName=%s", ipt.FirstName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.LastName=%s", ipt.LastName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.BeLong=%s", ipt.BeLong)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.MailAddress=%s", ipt.MailAddress)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.PhoneNumber=%s", ipt.PhoneNumber)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.Role=%v", ipt.Role)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : ipt.FirstLoginFlg=%b", ipt.FirstLoginFlg)))

	// 3.2. 要求電文チェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, false)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.operator.ValidateAppV1OperatorRegistration(ctx, ipt) call.")))
	err = i.operator.ValidateAppV1OperatorRegistration(ctx, ipt)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	//3.4 SESSION取得
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.SessionGet(ctx, c) call.")))
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	// SESSION.auth_statusが未設定の場合エラー
	// SESSION.auth_statusが1以外の場合エラー
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, 'BFWE01010106', 'BFWE01010107') call.")))
	err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, "BFWE01050121", "BFWE01050122")
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	errorResult, err := i.common.ValidateAccountStatus(
		ctx,
		sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
		sess.Values[appsesson.SessionVarOperatorID].(string),
		iptMeta,
		"detail",
		"BFWE01050124",
	)
	if errorResult != nil {
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)

	operatorRoleRegistration := &model.OperatorRoleRegistration{
		CardMemberRole: ipt.Role.CardMemberRole,
		ClubRole:       ipt.Role.ClubRole,
		ClubOwnerRole:  ipt.Role.ClubOwnerRole,
		MessageRole:    ipt.Role.MessageRole,
		OperatorRole:   ipt.Role.OperatorRole,
	}

	requestRegisterOperator := &model.OperatorRegistration{
		UserName:      ipt.FirstName + " " + ipt.LastName,
		BeLong:        ipt.BeLong,
		MailAddress:   ipt.MailAddress,
		PhoneNumber:   ipt.PhoneNumber,
		Password:      ipt.Password,
		Role:          *operatorRoleRegistration,
		FirstLoginFlg: *ipt.FirstLoginFlg,
	}
	// 3.6 管理者登録
	// /v1/operator/register
	url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/operator/register"

	jsonBody, err := json.Marshal(requestRegisterOperator)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	// Set header request
	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          sessionID,
		HeaderXNudgeSecurityRequest: "1",
	}
	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// A）正常の場合
	if res.StatusCode == 200 {
		data := new(model.ResponseOperatorRegistration)
		err = json.Unmarshal(body, &data)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.SetSessionAuthInfo call.")))
		err = i.common.SetSessionAuthInfo(
			ctx,
			sess,
			sessionID,
			sess.Values[appsesson.SessionVarOperatorID].(string),
			sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			sess.Values[appsesson.SessionValAccountStatus].(string),
			sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
			sess.Values[appsesson.SessionValOperatorFirstName].(string),
			sess.Values[appsesson.SessionValOperatorFamilyName].(string),
			sess.Values[appsesson.SessionValOperatorRole].(string),
			sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
			sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
		)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}

	} else {
		// B）　異常の場合
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// 3.7 SESSION保存※ttl更新
	expiresOn := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))
	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expiresOn
	sess.Save(c.Request(), c.Response())

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	session := &model.SessionUpdate{
		AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
		Expired:     expiresOn.Unix(),
		UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
		UpdatePgmID: contextw.GetAPIID(ctx),
	}
	keyUpdate := &model.UpdateSessionCondition{
		SessionID: sess.ID,
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
	err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	// 3.8 応答電文返却
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.pre.ViewAppV1OperatorRegistration(c) call.")))
	i.pre.ViewAppV1OperatorRegistration(c)
}

// AppV1OperatorUpdate ...
func (i Interactor) AppV1OperatorUpdate(c echo.Context, ipt input.OperatorUpdate, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : func (i Interactor) AppV1OperatorUpdate(c echo.Context, ipt input.OperatorUpdate, iptMeta input.MetaRequest) start.")))

	ctx := c.Request().Context()

	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010502")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.ReceptionProcessingStatus=%s", iptMeta.ReceptionProcessingStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.Role=%s", ipt.OperatorID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.FirstName=%s", ipt.FirstName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.LastName=%s", ipt.LastName)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.Belong=%s", ipt.Belong)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.Password=%s", ipt.Password)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : ipt.Role=%v", ipt.Role)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, true)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.operator.ValidateAppV1OperatorUpdate(ctx, ipt) call.")))
	// 3.3. パターンチェック
	err = i.operator.ValidateAppV1OperatorUpdate(ctx, ipt)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.SessionGet(ctx, c) call.")))
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 3.5 ログインステータスチェック
	// チェック仕様 4.1: SESSION.auth_statusが未設定の場合エラー
	// チェック仕様 4.2: SESSION.auth_statusが1以外の場合エラー
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, 'BFWE01050217', 'BFWE01050218') call.")))
	err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, "BFWE01050217", "BFWE01050218")
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	errorResult, err := i.common.ValidateAccountStatus(
		ctx,
		sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
		sess.Values[appsesson.SessionVarOperatorID].(string),
		iptMeta,
		"detail",
		"BFWE01050219",
	)
	if errorResult != nil {
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	//

	//「パスワード変更API」を呼び出す
	// /v1/auth/update
	url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/operator/update"
	// セッションの"id"
	sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)

	operatorUpdateRequestToSVC := &input.OperatorUpdateRequestToSVC{
		OperatorID: ipt.OperatorID,
		UserName:   ipt.FirstName + " " + ipt.LastName,
		Belong:     ipt.Belong,
		Password:   ipt.Password,
		Role:       ipt.Role,
	}
	if ipt.Password != "" {
		operatorUpdateRequestToSVC.AccountStatus = "01"
	}

	jsonBodyStr, err := json.Marshal(operatorUpdateRequestToSVC)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyStr))
	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          "1",
		HeaderXNudgeSecurityRequest: "1",
	}

	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// A）　正常の場合
	if res.StatusCode == 200 {

		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.SetCtxUserInfo(ctx, c, sessionID) call.")))
		ctx, err = i.commoncommon.SetCtxUserInfo(ctx, c, sessionID)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		// 3.6 SESSION更新
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.SetSessionAuthInfo call.")))
		err = i.common.SetSessionAuthInfo(
			ctx,
			sess,
			sessionID,
			sess.Values[appsesson.SessionVarOperatorID].(string),
			sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			sess.Values[appsesson.SessionValAccountStatus].(string),
			sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
			sess.Values[appsesson.SessionValOperatorFirstName].(string),
			sess.Values[appsesson.SessionValOperatorFamilyName].(string),
			sess.Values[appsesson.SessionValOperatorRole].(string),
			sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
			sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
		)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : sess.Values[appsesson.SessionVarOperatorAuthStatus]=" + sess.Values[appsesson.SessionVarOperatorAuthStatus].(string))))
		// 3.8 応答電文返却
		// 応答パターン4:個別設定事項
	} else {
		// B）　異常の場合
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// セッション保存※DynamoDBに登録
	expired := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))

	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expired
	sess.Save(c.Request(), c.Response())
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	session := &model.SessionUpdate{
		AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
		Expired:     expired.Unix(),
		UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
		UpdatePgmID: contextw.GetAPIID(ctx),
	}
	keyUpdate := &model.UpdateSessionCondition{
		SessionID: sess.ID,
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
	err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorUpdate : i.pre.ViewAppV1OperatorUpdate(c, operatorUpdateResponse) call.")))
	i.pre.ViewAppV1OperatorUpdate(c)
}

// AppV1OperatorDelete ...
func (i Interactor) AppV1OperatorDelete(c echo.Context, ipt input.OperatorDelete, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : func (i Interactor) AppV1OperatorDelete(c echo.Context, ipt input.OperatorDelete, iptMeta input.MetaRequest) start.")))

	ctx := c.Request().Context()

	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010503")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.ReceptionProcessingStatus=%s", iptMeta.ReceptionProcessingStatus)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : ipt.OperatorID=%s", ipt.OperatorID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : ipt.AccountStatus=%s", ipt.AccountStatus)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, true)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.operator.ValidateAppV1OperatorDelete(ctx, ipt) call.")))
	// 3.2. 要求電文チェック
	// チェック仕様No1-1～1-15を実施する（要求電文内で要確認の項目があれば、更新する）
	err = i.operator.ValidateAppV1OperatorDelete(ctx, ipt)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 3.3 サービス利用可否チェック
	// チェック仕様No.2-1～2-2を実施する
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.SessionGet(ctx, c) call.")))
	// 3.4 SESSION取得
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 3.5 ログインステータスチェック
	// インスタンス 3.1: SESSION.auth_statusが未設定の場合エラー
	// インスタンス 3.2: SESSION.auth_statusが1以外の場合エラー
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, 'BFWE01050307', 'BFWE01050308') call.")))
	err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, "BFWE01050307", "BFWE01050308")
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 3.5.2 アカウントステータス取得
	errorResult, err := i.common.ValidateAccountStatus(
		ctx,
		sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
		sess.Values[appsesson.SessionVarOperatorID].(string),
		iptMeta,
		"detail",
		"BFWE01050309",
	)
	if errorResult != nil {
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	//「パスワード変更API」を呼び出す
	// /v1/auth/update
	url := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/operator/update"
	// セッションの"id"
	sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)

	operatorDeleteInput := &model.OperatorDelete{
		OperatorID:    ipt.OperatorID,
		AccountStatus: ipt.AccountStatus,
	}
	jsonBodyStr, err := json.Marshal(operatorDeleteInput)
	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusInternalServerError)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBodyStr))

	iptRequestMeta := &model.CommonRequestMeta{
		HeaderContentType:           iptMeta.HeaderContentType,
		HeaderUserAgent:             iptMeta.HeaderUserAgent,
		HeaderXNudgeInteractionID:   iptMeta.HeaderXNudgeInteractionID,
		HeaderXNudgeUserID:          sessionID,
		HeaderXNudgeSecurityRequest: "1",
	}
	err = i.common.SetHeaderRequest(ctx, req, iptRequestMeta)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// A）　正常の場合
	if res.StatusCode == 200 {

		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.SetCtxUserInfo call.")))
		ctx, err = i.commoncommon.SetCtxUserInfo(ctx, c, sessionID)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		// 3.6 SESSION更新
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.SetSessionAuthInfo call.")))
		err = i.common.SetSessionAuthInfo(
			ctx,
			sess,
			sessionID,
			sess.Values[appsesson.SessionVarOperatorID].(string),
			sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
			sess.Values[appsesson.SessionValAccountStatus].(string),
			sess.Values[appsesson.SessionValFirstLoginFlg].(bool),
			sess.Values[appsesson.SessionValOperatorFirstName].(string),
			sess.Values[appsesson.SessionValOperatorFamilyName].(string),
			sess.Values[appsesson.SessionValOperatorRole].(string),
			sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
			sess.Values[appsesson.SessionValOperatorPhoneNumber].(string),
		)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : sess.Values[appsesson.SessionVarOperatorAuthStatus]=" + sess.Values[appsesson.SessionVarOperatorAuthStatus].(string))))
	} else {
		// B）　異常の場合
		errorResult := new(model.ErrorResult)
		err = json.Unmarshal(body, &errorResult)
		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}

	// セッション保存※DynamoDBに登録
	expired := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))

	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expired
	sess.Save(c.Request(), c.Response())
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	session := &model.SessionUpdate{
		AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
		Expired:     expired.Unix(),
		UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
		UpdatePgmID: contextw.GetAPIID(ctx),
	}
	keyUpdate := &model.UpdateSessionCondition{
		SessionID: sess.ID,
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
	err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorDelete : i.pre.ViewAppV1OperatorDelete(c, operatorDeleteResponse) call.")))
	i.pre.ViewAppV1OperatorDelete(c)
}

// AppV1GetOperatorList ...
func (i Interactor) AppV1GetOperatorList(c echo.Context, ipt input.OperatorList, iptMeta input.MetaRequest) {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : func (i Interactor) AppV1GetOperatorList(c echo.Context, ipt input.OperatorList, iptMeta input.MetaRequest) start.")))
	// ***********************************************
	// CTX取得
	// ***********************************************
	ctx := c.Request().Context()

	// ***********************************************
	// API情報設定
	// ***********************************************
	ctx, _ = i.commoncommon.SetCtxAPIInfo(ctx, c, "BFW010504")

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : contextw.GetRequestLang(ctx)=%s", contextw.GetRequestLang(ctx))))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : iptMeta.HeaderContentType=%s", iptMeta.HeaderContentType)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : iptMeta.HeaderUserAgent=%s", iptMeta.HeaderUserAgent)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : iptMeta.HeaderXRequestedWith=%s", iptMeta.HeaderXRequestedWith)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : iptMeta.HeaderXNudgeInteractionID=%s", iptMeta.HeaderXNudgeInteractionID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : iptMeta.HeaderXAPIKey=%s", iptMeta.HeaderXAPIKey)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.SearchText=%s", ipt.SearchText)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.SortKey=%s", ipt.SortKey)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.SortOrder=%s", ipt.SortOrder)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.PageNum=%s", ipt.PageNum)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.PageSize=%s", ipt.PageSize)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : ipt.OperatorID=%s", ipt.OperatorID)))

	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.operator.ValidateInputMeta(ctx, iptMeta) call.")))
	err := i.common.ValidateInputMeta(ctx, iptMeta, false)
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 要求電文チェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.operator.ValidateAppV1GetOperatorList(ctx, ipt, iptMeta) call.")))
	err = i.operator.ValidateAppV1GetOperatorList(ctx, ipt)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// サービス利用可否チェック
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.commoncommon.CheckServiceAvailability(ctx) call.")))
	err = i.commoncommon.CheckServiceAvailability(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	//3.4 SESSION取得
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.SessionGet(ctx, c) call.")))
	sess, err := i.common.SessionGet(ctx, c)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}
	// 3.5 ステータスチェック
	// 3.5.1 ログインステータスチェック
	// SESSION.auth_statusが未設定の場合エラー
	// SESSION.auth_statusが1以外の場合エラー
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1OperatorRegistration : i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, 'BFWE01050411', 'BFWE01050412') call.")))
	err = i.common.ValidateAuthLoginStatus(ctx, c, appsesson.SessionValOperatorAuthStatus1, "BFWE01050411", "BFWE01050412")
	// 異常の場合
	if err != nil {
		// 応答電文を生成し、クライアントに返却する
		i.pre.ViewError(c, err)
		return
	}

	// 3.5.2 アカウントステータス取得
	errorResult, err := i.common.ValidateAccountStatus(
		ctx,
		sess.Values[appsesson.SessionValOperatorRoleDetail].(string),
		sess.Values[appsesson.SessionVarOperatorID].(string),
		iptMeta,
		"detail",
		"BFWE01050415",
	)
	if errorResult != nil {
		i.pre.ViewErrorMCS(c, errorResult)
		return
	}
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// sessionID := sess.Values[appsesson.SessionVarOperatorAuthID].(string)
	var OperatorListRequestToSVCPattern1 = new(input.OperatorListRequestToSVCPattern1)
	var OperatorListRequestToSVCPattern2 = new(input.RequestToSVCPattern2)
	var operatorRole model.OperatorRoleInfo
	var jsonBodyStrCheck []byte
	dataResultPattern2 := new(model.OperatorListSearchResponsePattern2)
	dataResultPattern1 := new(model.OperatorListSearchResponse)
	apiURLCheck := config.GetProjectConfig().AppEnvConf.EndPoint.MCSOperator + "/v1/operator/list"
	json.Unmarshal([]byte(sess.Values[appsesson.SessionValOperatorRoleDetail].(string)), &operatorRole)
	if ipt.OperatorID != "" {
		Pattern1 := &input.OperatorListRequestToSVCPattern1{
			OperatorID: sess.Values[appsesson.SessionVarOperatorID].(string), // セッションから取得したIDを設定
			Role:       input.Role(operatorRole),
		}
		OperatorListRequestToSVCPattern1 = Pattern1
		jsonBodyStrCheck, err = json.Marshal(OperatorListRequestToSVCPattern1)

		if err != nil {
			i.pre.ViewError(c, err)
			return
		}
	} else {
		Pattern2 := &input.RequestToSVCPattern2{
			SearchText:   ipt.SearchText,
			SortKey:      ipt.SortKey,
			SortOrder:    ipt.SortOrder,
			PageNum:      ipt.PageNum,
			PageSize:     ipt.PageSize,
			OperatorRole: operatorRole.OperatorRole,
		}
		OperatorListRequestToSVCPattern2 = Pattern2
		jsonBodyStrCheck, err = json.Marshal(OperatorListRequestToSVCPattern2)

		if err != nil {
			i.pre.ViewError(c, err)
			return
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
	client := &http.Client{}
	err = i.common.SetHeaderRequest(ctx, reqCheck, iptRequestMeta)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	resCheck, err := client.Do(reqCheck)
	if err != nil {
		panic(err)
	}
	defer resCheck.Body.Close()
	bodyResult, _ := ioutil.ReadAll(resCheck.Body)

	if resCheck.StatusCode == 200 {
		if ipt.OperatorID != "" {
			err = json.Unmarshal(bodyResult, &dataResultPattern2)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
		} else {
			err = json.Unmarshal(bodyResult, &dataResultPattern1)
			if err != nil {
				i.pre.ViewError(c, err)
				return
			}
		}
	}
	// セッション保存※DynamoDBに登録
	expiresOn := time.Now().Add(time.Second * time.Duration(config.GetProjectConfig().AppEnvConf.Session[0].ExpiresInSeconds))
	sess.Values[appsesson.SessionVarMysqlstoreExpiresOn] = expiresOn
	sess.Save(c.Request(), c.Response())
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.common.InitNewSDKSession(ctx) call.")))
	awsSess, err := i.common.InitNewSDKSession(ctx)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	session := &model.SessionUpdate{
		AuthStatus:  sess.Values[appsesson.SessionVarOperatorAuthStatus].(string),
		Expired:     expiresOn.Unix(),
		UpdatedDate: util.TimeFormatMillisecondV1(util.NowTimeV1()),
		UpdatePgmID: contextw.GetAPIID(ctx),
	}
	keyUpdate := &model.UpdateSessionCondition{
		SessionID: sess.ID,
	}
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession) call.")))
	err = i.common.UpdateKey(ctx, awsSess, keyUpdate, session, config.GetProjectConfig().AppEnvConf.DynamoDB.TableSession)
	if err != nil {
		i.pre.ViewError(c, err)
		return
	}

	// 応答電文を生成し、クライアントに返却する
	if ipt.OperatorID != "" {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.pre.ViewAppV1GetOperatorList(c) call.")))
		i.pre.ViewAppV1GetOperatorListPattern2(c, dataResultPattern2)
	} else {
		apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/interactor.go : AppV1GetOperatorList : i.pre.ViewAppV1GetOperatorList(c) call.")))
		i.pre.ViewAppV1GetOperatorListPattern1(c, dataResultPattern1)
	}
}
