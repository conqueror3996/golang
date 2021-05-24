package config

import (
	"encoding/json"
)

// SevenBankErrorCode0000 正常終了
const SevenBankErrorCode0000 = "0000"

// SevenBankErrorCode6001 カード使用不可
const SevenBankErrorCode6001 = "6001"

// SevenBankErrorCode6002 カード回収
const SevenBankErrorCode6002 = "6002"

// SevenBankErrorCode6004 カード再発行
const SevenBankErrorCode6004 = "6004"

// SevenBankErrorCode6006 暗証相違（入力リトライ無し）
const SevenBankErrorCode6006 = "6006"

// SevenBankErrorCode6007 残高不足
const SevenBankErrorCode6007 = "6007"

// SevenBankErrorCode6010 指定金額不可
const SevenBankErrorCode6010 = "6010"

// SevenBankErrorCode6012 預入不可
const SevenBankErrorCode6012 = "6012"

// SevenBankErrorCode6015 その他エラー
const SevenBankErrorCode6015 = "6015"

// SevenBankErrorCode6020 取扱時間終了
const SevenBankErrorCode6020 = "6020"

// SevenBankErrorCode6031 取引限度額オーバー
const SevenBankErrorCode6031 = "6031"

// SevenBankErrorCode6033 1日利用回数オーバー
const SevenBankErrorCode6033 = "6033"

// SevenBankErrorCode6035 1回利用限度額オーバー
const SevenBankErrorCode6035 = "6035"

// SevenBankErrorCode6081 取引予約時間超過
const SevenBankErrorCode6081 = "6081"

// SevenBankErrorCode6082 QR情報照合エラー
const SevenBankErrorCode6082 = "6082"

// SevenBankErrorCode6083 取引予約無し
const SevenBankErrorCode6083 = "6083"

// SevenBankErrorCode6150 フォーマットエラー
const SevenBankErrorCode6150 = "6150"

// SevenBankErrorCode6121 元取引なし
const SevenBankErrorCode6121 = "6121"

// SevenBankErrorCode6129 取消処理不可
const SevenBankErrorCode6129 = "6129"

// SevenBankErrorCode6301 インターネット提携先認証エラー
const SevenBankErrorCode6301 = "6301"

// SevenBankErrorCode6302 インターネット提携先明細ロック中
const SevenBankErrorCode6302 = "6302"

// SevenBankErrorCode6303 インターネット提携先鍵取得エラー
const SevenBankErrorCode6303 = "6303"

// SevenBankErrorCode6304 インターネット提携先システムビジー
const SevenBankErrorCode6304 = "6304"

// SevenBankErrorCode6305 インターネット提携先システム内部処理エラー
const SevenBankErrorCode6305 = "6305"

// SevenBankErrorCodeE001 JWS署名検証エラー
const SevenBankErrorCodeE001 = "E001"

// SevenBankErrorCodeE010 JTI重複チェックエラー
const SevenBankErrorCodeE010 = "E010"

// SevenBankErrorCodeE011 JTI有効期限チェックエラー
const SevenBankErrorCodeE011 = "E011"

// SevenBankErrorCodeE012 JTIクライアントチェックエラー
const SevenBankErrorCodeE012 = "E012"

// SevenBankErrorCodeE020 アクセストークン有効期限チェックエラー
const SevenBankErrorCodeE020 = "E020"

// SevenBankErrorCodeE021 アクセストークン有効性チェックエラー
const SevenBankErrorCodeE021 = "E021"

// SevenBankErrorCodeE022 アクセストークンスコープチェックエラー
const SevenBankErrorCodeE022 = "E022"

var sevenBankErrorCodeConfig map[string]interface{}

func init() {
	// fmt.Println("config/seven_bank_error_code.go : init : func init() start.")

	err := json.Unmarshal([]byte(SevenBankErrorCodesConfig), &sevenBankErrorCodeConfig)
	if err != nil {
		panic(err)
	}

	// fmt.Println("config/seven_bank_error_code.go : init : func init() end.")
}

// GetSevenBankErrorCode ...
func GetSevenBankErrorCode(errorCode string) string {
	return (sevenBankErrorCodeConfig[errorCode].(map[string]interface{}))["errorCode"].(string)
}

// GetSevenBankStatusCode ...
func GetSevenBankStatusCode(errorCode string) int {
	return int((sevenBankErrorCodeConfig[errorCode].(map[string]interface{}))["statusCode"].(float64))
}
