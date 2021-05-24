package usecase

import (
	"bff-web/config"
	"bff-web/usecase/input"
	"bff-web/util"
	"context"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

// Card ...
type Card interface {
	ValidateAppV1CardGetMemberList(ctx context.Context, ipt input.CardMemberList) error
	ValidateAppV1CardGetHistoryUsage(ctx context.Context, ipt input.CardHistoryUsage) error
	ValidateAppV1CardGetHistoryRepayment(ctx context.Context, ipt input.CardHistoryRepayment) error
	ValidateAppV1CardGetHistoryBenefit(ctx context.Context, ipt input.CardHistoryBenefit) error
	ValidateAppV1CardGetHistoryNotification(ctx context.Context, ipt input.CardHistoryNotification) error
}

// NewCard ...
func NewCard(
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	errm apperror.ErrorManager,
) Card {
	return card{
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		errm,
	}
}

type card struct {
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	errm                                   apperror.ErrorManager
}

func (u card) ValidateAppV1CardGetMemberList(ctx context.Context, ipt input.CardMemberList) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetMemberList : func (u operator) ValidateAppV1CardGetMemberList(ctx context.Context, ipt input.CardMemberList) error start.")))

	if ipt.MemberID != "" {

		if !util.MatchAlphanumericSymbolStringV1(ipt.MemberID) {
			msgCD := "BFWE01020106"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if utf8.RuneCountInString(ipt.MemberID) > 36 {
			msgCD := "BFWE01020106"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortKey != "" {
			if ipt.SortKey != "06" {
				if ipt.SearchText != "" && ipt.MemberID != "" && (ipt.SortOrder == "" || ipt.PageNum == "" || ipt.PageSize == "") {
					msgCD := "BFWE01020111"
					return u.errm.Wrap1(
						errors.New(""),
						contextw.GetRequestLang(ctx),
						msgCD,
						"",
					)
				}
			} else {
				if ipt.SearchText == "" && ipt.SortOrder == "" {
					msgCD := "BFWE01020112"
					return u.errm.Wrap1(
						errors.New(""),
						contextw.GetRequestLang(ctx),
						msgCD,
						"",
					)
				}
			}
		}

		// チェック仕様No.3-1を実施する
		if ipt.SearchText != "" || ipt.SortKey != "" || ipt.SortOrder != "" || ipt.PageNum != "" || ipt.PageSize != "" {
			msgCD := "BFWE01020109"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	} else {

		if !util.MatchAlphanumericSymbolStringV1(ipt.SearchText) {
			msgCD := "BFWE01020101"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		// // チェック仕様No.3-2を実施する
		if ipt.SortKey == "" {
			msgCD := "BFWE01020110"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchNumberStringV1(ipt.SortKey) {
			msgCD := "BFWE01020102"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		arrSortKey := []string{"01", "02", "03", "04", "05", "06"}
		if !util.StringArrayContainsV1(arrSortKey, ipt.SortKey) {
			msgCD := "BFWE01020102"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortOrder == "" {
			msgCD := "BFWE01020103"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.StringArrayContainsV1([]string{"desc", "asc"}, ipt.SortOrder) {
			msgCD := "BFWE01020103"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageNum == "" {
			msgCD := "BFWE01020104"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		pageInt, err := strconv.Atoi(ipt.PageNum)
		if err != nil {
			msgCD := "BFWE01020104"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if pageInt <= 0 {
			msgCD := "BFWE01020104"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageSize == "" {
			msgCD := "BFWE01020105"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		maximumRecordsPerPageInt, err := strconv.Atoi(ipt.PageSize)
		if err != nil {
			msgCD := "BFWE01020105"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt <= 0 {
			msgCD := "BFWE01020105"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt > config.GetProjectConfig().AppEnvConf.App.Operator.MaximumRecordsPerPage01 {
			msgCD := "BFWE01020105"
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

func (u card) ValidateAppV1CardGetHistoryUsage(ctx context.Context, ipt input.CardHistoryUsage) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetHistoryUsage : func (u operator) ValidateAppV1CardGetHistoryUsage(ctx context.Context, ipt input.CardHistoryUsage) error start.")))

	if ipt.MemberID == "" {
		msgCD := "BFWE01020301"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.MemberID) > 36 {
		msgCD := "BFWE01020302"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	if !util.MatchAlphanumericSymbolStringV1(ipt.MemberID) {
		msgCD := "BFWE01020302"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.YearMonth == "" {
		msgCD := "BFWE01020303"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.YearMonth) > 6 {
		msgCD := "BFWE01020304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	if !util.MatchStringV1("^[0-9]*$", ipt.YearMonth) {
		msgCD := "BFWE01020304"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	return nil
}

func (u card) ValidateAppV1CardGetHistoryRepayment(ctx context.Context, ipt input.CardHistoryRepayment) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetHistoryRepayment : func (u operator) ValidateAppV1CardGetHistoryRepayment(ctx context.Context, ipt input.CardHistoryRepayment) error start.")))

	if ipt.UserID == "" {
		msgCD := "BFWE01020401"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.UserID) {
		msgCD := "BFWE01020402"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.UserID) > 36 {
		msgCD := "BFWE01020402"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.RepayMonth == "" {
		msgCD := "BFWE01020403"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.RepayMonth) {
		msgCD := "BFWE01020404"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.RepayMonth) > 6 {
		msgCD := "BFWE01020404"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u card) ValidateAppV1CardGetHistoryBenefit(ctx context.Context, ipt input.CardHistoryBenefit) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetHistoryBenefit : func (u operator) ValidateAppV1CardGetHistoryBenefit(ctx context.Context, ipt input.CardHistoryBenefit) error start.")))
	if ipt.MemberID == "" {
		msgCD := "BFWE01020501"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.MemberID) {
		msgCD := "BFWE01020502"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.MemberID) > 36 {
		msgCD := "BFWE01020502"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u card) ValidateAppV1CardGetHistoryNotification(ctx context.Context, ipt input.CardHistoryNotification) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetHistoryNotification : func (u operator) ValidateAppV1CardGetHistoryNotification(ctx context.Context, ipt input.CardHistoryNotification) error start.")))

	if ipt.MemberID == "" {
		msgCD := "BFWE01020601"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.MemberID) {
		msgCD := "BFWE01020602"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.MemberID) > 36 {
		msgCD := "BFWE01020602"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.SortKey == "" {
		msgCD := "BFWE01020603"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.SortKey) {
		msgCD := "BFWE01020604"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	arrSortKey := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}
	if !util.StringArrayContainsV1(arrSortKey, ipt.SortKey) {
		msgCD := "BFWE01020604"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.SortOrder == "" {
		msgCD := "BFWE01020605"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchStringV1("^[a-zA-Z]*$", ipt.SortOrder) {
		msgCD := "BFWE01020606"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.StringArrayContainsV1([]string{"desc", "asc"}, ipt.SortOrder) {
		msgCD := "BFWE01020606"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PageNum == "" {
		msgCD := "BFWE01020607"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	pageInt, err := strconv.Atoi(ipt.PageNum)
	if err != nil {
		msgCD := "BFWE01020608"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if pageInt <= 0 {
		msgCD := "BFWE01020608"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PageSize == "" {
		msgCD := "BFWE01020609"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	maximumRecordsPerPageInt, err := strconv.Atoi(ipt.PageSize)
	if err != nil {
		msgCD := "BFWE01020610"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if maximumRecordsPerPageInt <= 0 {
		msgCD := "BFWE01020610"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.DeletedData == "" {
		msgCD := "BFWE01020611"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.DeletedData) {
		msgCD := "BFWE01020612"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}
