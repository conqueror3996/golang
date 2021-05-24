package usecase

import (
	"bff-web/config"
	"bff-web/domain/model"
	"bff-web/usecase/input"
	"bff-web/util"
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

// Operator ...
type Operator interface {
	ValidateAppV1AuthLogin(ctx context.Context, ipt input.AuthLogin) error
	ValidateAppV1AuthResetRequest(ctx context.Context, ipt input.ResetRequest) error
	ValidateAppV1AuthChangePassword(ctx context.Context, ipt input.AuthChangePassword) error
	ValidateAppV1AuthReset(ctx context.Context, ipt input.AuthReset) error
	ValidateExpiredAuthReset(ctx context.Context, resetKey *model.ResetKey) error
	ValidateAppV1GetOperatorList(ctx context.Context, ipt input.OperatorList) error
	ValidateAppV1OperatorUpdate(ctx context.Context, ipt input.OperatorUpdate) error
	ValidateAppV1OperatorDelete(ctx context.Context, ipt input.OperatorDelete) error
	ValidateAppV1OperatorRegistration(ctx context.Context, ipt input.OperatorRegistration) error
}

// NewOperator ...
func NewOperator(
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	errm apperror.ErrorManager,
) Operator {
	return operator{
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		errm,
	}
}

type operator struct {
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	errm                                   apperror.ErrorManager
}

func (u operator) ValidateAppV1AuthLogin(ctx context.Context, ipt input.AuthLogin) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1AuthLogin : func (u operator) ValidateAppV1AuthLogin(ctx context.Context, ipt input.AuthLogin) error start.")))

	if ipt.ID != "" {
		if !util.MatchAlphanumericSymbolStringV1(ipt.ID) {
			msgCD := "BFWE01010101"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		// emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[\\.a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		emailRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_\.]{0,61}@[a-zA-Z0-9-]{2,}(\.[a-zA-Z0-9]{2,4}){1,2}$`)

		iptEmailAdress := strings.Trim(ipt.ID, " ")
		if (!emailRegex.MatchString(iptEmailAdress)) || (utf8.RuneCountInString(ipt.ID) > 255) {
			msgCD := "BFWE01010101"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.Password != "" {
		// if !util.MatchAlphanumericSymbolStringV1(ipt.Password) {
		// 	msgCD := "BFWE01010102"
		// 	return u.errm.Wrap1(
		// 		errors.New(""),
		// 		contextw.GetRequestLang(ctx),
		// 		msgCD,
		// 		"",
		// 	)
		// }

		// 英数字と記号　!#%&'/=`*+?{}^$-|.　を許容
		// New          !~<>,;:_=?*+#."&§%°()|[]-$^@/
		// 英字、数字、記号それぞれ1文字以上含むこと
		// 8文字以上0文字以下
		// 旧パワードと同じ場合エラー1つ前のパスワードはエラー）

		if !util.MatchStringV1("^[a-zA-Z0-9!~<>,;:_=?*+#.\"&§%°()|[\\]\\-$^@\\/]+$", ipt.Password) {
			msgCD := "BFWE01010102"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
		newPasswordSlice := strings.Split(ipt.Password, "")

		conditions := 0
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice { // !~<>,;:_=?*+#.\"&§%°()|[]-$^@/
			if util.StringArrayContainsV1([]string{"!", "#", "%", "&", "/", "=",
				"~", "\"", "°", "(", ")", "[", "]", "@", "§",
				"*", "+", "?", "^", "$", "-", "|",
				".", "<", ">", ",", ";", ":", "_", "\\"}, newPassword) {
				conditions++
				break
			}
		}
		if conditions != 4 {
			msgCD := "BFWE01010102"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if utf8.RuneCountInString(ipt.Password) < 8 || utf8.RuneCountInString(ipt.Password) > 12 {
			msgCD := "BFWE01010102"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.SmsCode != "" {
		if !util.MatchAlphanumericSymbolStringV1(ipt.SmsCode) {
			msgCD := "BFWE01010103"
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

func (u operator) ValidateAppV1AuthChangePassword(ctx context.Context, ipt input.AuthChangePassword) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/employee.go : ValidateAppV1AuthChangePassword : func (u employee) ValidateAppV1AuthChangePassword(ctx context.Context, ipt input.AuthChangePassword) error start.")))

	if ipt.Password == "" {
		msgCD := "BFWE01010401"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// if !util.MatchAlphanumericSymbolStringV1(ipt.Password) {
	// 	msgCD := "BFWE01010402"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	// 英数字と記号　!#%&'/=~`*+?{}^$-|.　を許容
	// New          !~<>,;:_=?*+#."&§%°()|[]-$^@/
	// 英字、数字、記号それぞれ1文字以上含むこと
	// 8文字以上20文字以下
	// 旧パスワードと同じ場合エラー（1つ前のパスワードはエラー）

	if !util.MatchStringV1("^[a-zA-Z0-9!~<>,;:_=?*+#.\"&§%°()|[\\]\\-$^@\\/]+$", ipt.Password) {
		msgCD := "BFWE01010402"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	newPasswordSlice := strings.Split(ipt.Password, "")

	conditions := 0
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice { // !~<>,;:_=?*+#.\"&§%°()|[]-$^@/
		if util.StringArrayContainsV1([]string{"!", "#", "%", "&", "/", "=",
			"~", "\"", "°", "(", ")", "[", "]", "@", "§",
			"*", "+", "?", "^", "$", "-", "|",
			".", "<", ">", ",", ";", ":", "_", "\\"}, newPassword) {
			conditions++
			break
		}
	}
	if conditions != 4 {
		msgCD := "BFWE01010402"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.Password) < 8 || utf8.RuneCountInString(ipt.Password) > 12 {
		msgCD := "BFWE01010402"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u operator) ValidateAppV1AuthResetRequest(ctx context.Context, ipt input.ResetRequest) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1AuthResetRequest : func (u operator) ValidateAppV1AuthResetRequest(ctx context.Context, ipt input.ResetRequest) error start.")))

	if ipt.MailAddr == "" {
		msgCD := "BFWE01010201"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[\\.a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_\.]{0,61}@[a-zA-Z0-9-]{2,}(\.[a-zA-Z0-9]{2,4}){1,2}$`)
	iptEmailAdress := strings.Trim(ipt.MailAddr, " ")
	if !emailRegex.MatchString(iptEmailAdress) {
		msgCD := "BFWE01010202"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u operator) ValidateAppV1AuthReset(ctx context.Context, ipt input.AuthReset) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1AuthReset : func (u operator) ValidateAppV1AuthReset(ctx context.Context, ipt input.AuthReset) error start.")))

	if ipt.Key == "" {
		msgCD := "BFWE01010301"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.Key) {
		msgCD := "BFWE01010302"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.Password == "" {
		msgCD := "BFWE01010303"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// if !util.MatchAlphanumericSymbolStringV1(ipt.Password) {
	// 	msgCD := "BFWE01010304"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	// 英数字と記号　!#%&'/=`*+?{}^$-|.　を許容
	// New          !~<>,;:_=?*+#."&§%°()|[]-$^@/
	// 英字、数字、記号それぞれ1文字以上含むこと
	// 8文字以上20文字以下
	// 旧パスワードと同じ場合エラー（1つ前のパスワードはエラー）

	if !util.MatchStringV1("^[a-zA-Z0-9!~<>,;:_=?*+#.\"&§%°()|[\\]\\-$^@\\/]+$", ipt.Password) {
		msgCD := "BFWE01010304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	newPasswordSlice := strings.Split(ipt.Password, "")

	conditions := 0
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice { // !~<>,;:_=?*+#.\"&§%°()|[]-$^@/
		if util.StringArrayContainsV1([]string{"!", "#", "%", "&", "'", "/", "=",
			"~", "\"", "°", "(", ")", "[", "]", "@", "§",
			"`", "*", "+", "?", "{", "}", "^", "$", "-", "|",
			".", "<", ">", ",", ";", ":", "_", "\\"}, newPassword) {
			conditions++
			break
		}
	}
	if conditions != 4 {
		msgCD := "BFWE01010304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.Password) < 8 || utf8.RuneCountInString(ipt.Password) > 12 {
		msgCD := "BFWE01010304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u operator) ValidateExpiredAuthReset(ctx context.Context, resetKey *model.ResetKey) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateExpiredAuthReset : func (u common) ValidateExpiredAuthReset(ctx context.Context, resetKey *model.ResetKey) error start.")))

	timeNow := time.Now().Unix()

	// C) 検索結果0件の場合
	if resetKey == nil {
		msgCD := "01010302"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// B) 結果が取得でき、かつ現在時刻が取得した項目の「expired※unixtime表記」の時刻を過ぎていた場合
	if timeNow > resetKey.Expired {
		msgCD := "01010301"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u operator) ValidateAppV1GetOperatorList(ctx context.Context, ipt input.OperatorList) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1GetOperatorList : func (u operator) ValidateAppV1GetOperatorList(ctx context.Context, ipt input.OperatorList) error start.")))

	if ipt.OperatorID != "" {

		if utf8.RuneCountInString(ipt.OperatorID) != 8 {
			msgCD := "BFWE01050410"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchStringV1("^[0-9]*$", ipt.OperatorID) {
			msgCD := "BFWE01050410"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	} else {
		if ipt.SearchText != "" {
			if !util.MatchAlphanumericSymbolStringV1(ipt.SearchText) {
				msgCD := "BFWE01050401"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if ipt.SortKey == "" {
			msgCD := "BFWE01050402"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchNumberStringV1(ipt.SortKey) {
			msgCD := "BFWE01050403"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		arrSortKey := []string{"01", "02", "03", "04"}
		if !util.StringArrayContainsV1(arrSortKey, ipt.SortKey) {
			msgCD := "BFWE01050403"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortOrder == "" {
			msgCD := "BFWE01050404"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortOrder != "asc" && ipt.SortOrder != "desc" {
			msgCD := "BFWE01050405"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageNum == "" {
			msgCD := "BFWE01050406"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		iptPageNum := strings.Trim(ipt.PageNum, " ")
		if !util.MatchStringV1("^[0-9]*$", iptPageNum) {
			msgCD := "BFWE01050407"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		pageInt, err := strconv.Atoi(ipt.PageNum)
		if err != nil {
			msgCD := "BFWE01050407"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if pageInt <= 0 {
			msgCD := "BFWE01050407"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageSize == "" {
			msgCD := "BFWE01050408"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		iptPageSize := strings.Trim(ipt.PageSize, " ")
		if !util.MatchStringV1("^[0-9]*$", iptPageSize) {
			msgCD := "BFWE01050409"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		maximumRecordsPerPageInt, err := strconv.Atoi(ipt.PageSize)
		if err != nil {
			msgCD := "BFWE01050409"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt <= 0 {
			msgCD := "BFWE01050409"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt > config.GetProjectConfig().AppEnvConf.App.Operator.MaximumRecordsPerPage01 {
			msgCD := "BFWE01050409"
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

func (u operator) ValidateAppV1OperatorUpdate(ctx context.Context, ipt input.OperatorUpdate) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/employee.go : ValidateAppV1OperatorUpdate : func (u operator) ValidateAppV1OperatorUpdate(ctx context.Context, ipt input.OperatorUpdate) error start.")))

	if ipt.OperatorID == "" {
		msgCD := "BFWE01050201"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptOperatorID := strings.Trim(ipt.OperatorID, " ")
	if !util.MatchStringV1("^[0-9]*$", iptOperatorID) {
		msgCD := "BFWE01050202"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	nameRegex := regexp.MustCompile("^[a-zA-Z ]*$")
	iptFirstName := strings.Trim(ipt.FirstName, " ")
	iptLastName := strings.Trim(ipt.LastName, " ")

	if iptLastName != "" {
		if !nameRegex.MatchString(iptLastName) {
			msgCD := "BFWE01050203"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if iptFirstName == "" {
			msgCD := "BFWE01050220"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if iptFirstName != "" {
		if !nameRegex.MatchString(iptFirstName) {
			msgCD := "BFWE01050204"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if iptLastName == "" {
			msgCD := "BFWE01050221"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	iptBelong := strings.Trim(ipt.Belong, " ")
	if iptBelong != "" {
		if !util.MatchStringV1("^[a-zA-Z0-9!@#%&'/=~`*+?{}\\^$\\-|.]+$", iptBelong) {
			msgCD := "BFWE01050205"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.MailAddress != "" {
		if !util.MatchAlphanumericSymbolStringV1(ipt.MailAddress) {
			msgCD := "BFWE01050206"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		// emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[\\.a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		emailRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_\.]{0,61}@[a-zA-Z0-9-]{2,}(\.[a-zA-Z0-9]{2,4}){1,2}$`)
		iptEmailAdress := strings.Trim(ipt.MailAddress, " ")
		if (!emailRegex.MatchString(iptEmailAdress)) || (utf8.RuneCountInString(ipt.MailAddress) > 255) {
			msgCD := "BFWE01050207"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.PhoneNumber != "" {
		if !util.MatchStringV1("^[0-9\\-]+$", ipt.PhoneNumber) {
			msgCD := "BFWE01050208"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		// if utf8.RuneCountInString(ipt.PhoneNumber) > 21 {
		// 	msgCD := "BFWE01050208"
		// 	return u.errm.Wrap1(
		// 		errors.New(""),
		// 		contextw.GetRequestLang(ctx),
		// 		msgCD,
		// 		"",
		// 	)
		// }

		// if utf8.RuneCountInString(ipt.PhoneNumber) < 10 || utf8.RuneCountInString(ipt.PhoneNumber) > 11 {
		// 	msgCD := "BFWE01050219"
		// 	return u.errm.Wrap1(
		// 		errors.New(""),
		// 		contextw.GetRequestLang(ctx),
		// 		msgCD,
		// 		"",
		// 	)
		// }
	}
	newRole := &input.Role{}
	if *newRole != ipt.Role {
		if reflect.DeepEqual(input.Role{}, ipt.Role) {
			msgCD := "BFWE01050211"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
		arrRolesType := []string{"01", "02", "03", "04"}
		if ipt.Role.CardMemberRole != "" {
			if !util.StringArrayContainsV1(arrRolesType, ipt.Role.CardMemberRole) {
				msgCD := "BFWE01050211"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if ipt.Role.ClubOwnerRole != "" {
			if !util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubOwnerRole) {
				msgCD := "BFWE01050211"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if ipt.Role.ClubRole != "" {
			if !util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubRole) {
				msgCD := "BFWE01050211"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if ipt.Role.MessageRole != "" {
			if !util.StringArrayContainsV1(arrRolesType, ipt.Role.MessageRole) {
				msgCD := "BFWE01050211"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if ipt.Role.OperatorRole != "" {
			if !util.StringArrayContainsV1(arrRolesType, ipt.Role.OperatorRole) {
				msgCD := "BFWE01050211"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		// if !util.StringArrayContainsV1(arrRolesType, ipt.Role.CardMemberRole) ||
		// 	!util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubRole) ||
		// 	!util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubOwnerRole) ||
		// 	!util.StringArrayContainsV1(arrRolesType, ipt.Role.MessageRole) ||
		// 	!util.StringArrayContainsV1(arrRolesType, ipt.Role.OperatorRole) {
		// 	msgCD := "BFWE01050211"
		// 	return u.errm.Wrap1(
		// 		errors.New(""),
		// 		contextw.GetRequestLang(ctx),
		// 		msgCD,
		// 		"",
		// 	)
		// }
	}

	if ipt.FirstLoginFlg == nil {
		msgCD := "BFWE01050212"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	//3.3. パターンチェック
	// チェック仕様No2-1～2-2を実施する（追加の入力パターンが存在する場合は、更新する）
	if ipt.Password != "" {
		// 英数字と記号　!#%&'/=~`*+?{}^$-|.　を許容
		// New          !~<>,;:_=?*+#."&§%°()|[]-$^@/
		// 英字、数字、記号それぞれ1文字以上含むこと
		// 8文字以上20文字以下
		// 旧パスワードと同じ場合エラー（1つ前のパスワードはエラー）

		if !util.MatchStringV1("^[a-zA-Z0-9!~<>,;:_=?*+#.\"&§%°()|[\\]\\-$^@\\/]+$", ipt.Password) {
			msgCD := "BFWE01050209"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
		newPasswordSlice := strings.Split(ipt.Password, "")

		conditions := 0
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice {
			if util.StringArrayContainsV1([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, newPassword) {
				conditions++
				break
			}
		}
		for _, newPassword := range newPasswordSlice { // !~<>,;:_=?*+#.\"&§%°()|[]-$^@/
			if util.StringArrayContainsV1([]string{"!", "#", "%", "&", "/", "=",
				"~", "\"", "°", "(", ")", "[", "]", "@", "§",
				"*", "+", "?", "^", "$", "-", "|",
				".", "<", ">", ",", ";", ":", "_", "\\"}, newPassword) {
				conditions++
				break
			}
		}
		if conditions != 4 {
			msgCD := "BFWE01050210"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if utf8.RuneCountInString(ipt.Password) < 8 || utf8.RuneCountInString(ipt.Password) > 20 {
			msgCD := "BFWE01050210"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.FirstLoginFlg != nil && *ipt.FirstLoginFlg != true {
			msgCD := "BFWE01050213"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	} else {
		if ipt.FirstLoginFlg != nil && *ipt.FirstLoginFlg == true {
			msgCD := "BFWE01050213"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.OperatorID != "" {
		if ipt.FirstName == "" && ipt.LastName == "" && ipt.Belong == "" && ipt.MailAddress == "" && ipt.Password == "" && ipt.PhoneNumber == "" {
			msgCD := "BFWE01050214"
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

func (u operator) ValidateAppV1OperatorDelete(ctx context.Context, ipt input.OperatorDelete) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/employee.go : ValidateAppV1OperatorDelete : func (u operator) ValidateAppV1OperatorDelete(ctx context.Context, ipt input.OperatorDelete) error start.")))

	if ipt.OperatorID == "" {
		msgCD := "BFWE01050301"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptOperatorID := strings.Trim(ipt.OperatorID, " ")
	if !util.MatchStringV1("^[0-9]*$", iptOperatorID) {
		msgCD := "BFWE01050302"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.AccountStatus == "" {
		msgCD := "BFWE01050303"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptAccountStatus := strings.Trim(ipt.AccountStatus, " ")
	if !util.MatchStringV1("^[0-9]*$", iptAccountStatus) && iptAccountStatus != "03" {
		msgCD := "BFWE01050304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u operator) ValidateAppV1OperatorRegistration(ctx context.Context, ipt input.OperatorRegistration) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/common.go : ValidateAppV1OperatorRegistration : func (u common) ValidateAppV1OperatorRegistration(ctx context.Context, ipt input.OperatorRegistration) error start.")))

	//check message_config
	// if ipt.UserName == "" {
	// 	msgCD := "BFWE01050101"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	// iptUserName := strings.Trim(ipt.UserName, " ")
	// if !util.MatchStringV1("^[a-zA-Z ]*$", iptUserName) || (utf8.RuneCountInString(iptUserName) > 255) {
	// 	msgCD := "BFWE01050102"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	if ipt.BeLong == "" {
		msgCD := "BFWE01050105"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// iptBeLong := strings.Trim(ipt.BeLong, " ")
	// if !util.MatchStringV1("^[0-9]+$", iptBeLong) || (utf8.RuneCountInString(iptBeLong) != 2) {
	// 	msgCD := "BFWE01050106"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	if !util.MatchAlphanumericSymbolStringV1(ipt.BeLong) {
		msgCD := "BFWE01050106"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	arrBeLong := []string{"01", "02", "03"}
	if !util.StringArrayContainsV1(arrBeLong, ipt.BeLong) {
		msgCD := "BFWE01050106"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.MailAddress == "" {
		msgCD := "BFWE01050107"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.MailAddress) {
		msgCD := "BFWE01050108"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[\\.a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_\.]{0,61}@[a-zA-Z0-9-]{2,}(\.[a-zA-Z0-9]{2,4}){1,2}$`)
	iptEmailAdress := strings.Trim(ipt.MailAddress, " ")
	if (!emailRegex.MatchString(iptEmailAdress)) || (utf8.RuneCountInString(ipt.MailAddress) > 255) {
		msgCD := "BFWE01050109"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PhoneNumber == "" {
		msgCD := "BFWE01050110"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchStringV1("^[0-9\\-]+$", ipt.PhoneNumber) {
		msgCD := "BFWE01050111"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.PhoneNumber) < 10 || utf8.RuneCountInString(ipt.PhoneNumber) > 11 {
		msgCD := "BFWE01050123"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.Password == "" {
		msgCD := "BFWE01050112"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// 英数字と記号　!#%&'/=~`*+?{}^$-|.　を許容
	// New          !~<>,;:_=?*+#."&§%°()|[]-$^@/
	// 英字、数字、記号それぞれ1文字以上含むこと
	// 8文字以上20文字以下
	// 旧パスワードと同じ場合エラー（1つ前のパスワードはエラー）

	if !util.MatchStringV1("^[a-zA-Z0-9!~<>,;:_=?*+#.\"&§%°()|[\\]\\-$^@\\/]+$", ipt.Password) {
		msgCD := "BFWE01050113"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	newPasswordSlice := strings.Split(ipt.Password, "")

	conditions := 0
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice {
		if util.StringArrayContainsV1([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, newPassword) {
			conditions++
			break
		}
	}
	for _, newPassword := range newPasswordSlice { // !~<>,;:_=?*+#.\"&§%°()|[]-$^@/
		if util.StringArrayContainsV1([]string{"!", "#", "%", "&", "/", "=",
			"~", "\"", "°", "(", ")", "[", "]", "@", "§",
			"*", "+", "?", "^", "$", "-", "|",
			".", "<", ">", ",", ";", ":", "_", "\\"}, newPassword) {
			conditions++
			break
		}
	}
	if conditions != 4 {
		msgCD := "BFWE01050114"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.Password) < 8 || utf8.RuneCountInString(ipt.Password) > 20 {
		msgCD := "BFWE01050114"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if reflect.DeepEqual(input.OperatorRoleRegistration{}, ipt.Role) {
		msgCD := "BFWE01050115"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	arrRolesType := []string{"01", "02", "03", "04"}

	if !util.StringArrayContainsV1(arrRolesType, ipt.Role.CardMemberRole) ||
		!util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubRole) ||
		!util.StringArrayContainsV1(arrRolesType, ipt.Role.ClubOwnerRole) ||
		!util.StringArrayContainsV1(arrRolesType, ipt.Role.MessageRole) ||
		!util.StringArrayContainsV1(arrRolesType, ipt.Role.OperatorRole) {
		msgCD := "BFWE01050116"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.FirstLoginFlg == nil {
		msgCD := "BFWE01050117"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.FirstLoginFlg != nil && *ipt.FirstLoginFlg != true {
		msgCD := "BFWE01050118"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.FirstName == "" {
		msgCD := "BFWE01050124"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptFirstName := strings.Trim(ipt.FirstName, " ")
	if !util.MatchStringV1("^[a-zA-Z ]*$", iptFirstName) || (utf8.RuneCountInString(iptFirstName) > 255) {
		msgCD := "BFWE01050125"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.LastName == "" {
		msgCD := "BFWE01050126"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptLastName := strings.Trim(ipt.LastName, " ")
	if !util.MatchStringV1("^[a-zA-Z ]*$", iptLastName) || (utf8.RuneCountInString(iptLastName) > 255) {
		msgCD := "BFWE01050127"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}
