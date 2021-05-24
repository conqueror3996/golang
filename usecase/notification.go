package usecase

import (
	"bff-web/config"
	"bff-web/usecase/input"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
	"github.com/nri4nudge/api-common/util"
)

// Notification ...
type Notification interface {
	ValidateAppV1NotificationList(ctx context.Context, ipt input.NotificationList) error
	ValidateAppV1NotificationRegistration(ctx context.Context, ipt input.NotificationRegistration) error
}

// NewNotification ...
func NewNotification(
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	errm apperror.ErrorManager,
) Notification {
	return notification{
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		errm,
	}
}

type notification struct {
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	errm                                   apperror.ErrorManager
}

func (n notification) ValidateAppV1NotificationRegistration(ctx context.Context, ipt input.NotificationRegistration) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1CardGetMemberList : func (u operator) ValidateAppV1CardGetMemberList(ctx context.Context, ipt input.CardMemberList) error start.")))

	if ipt.MessageSection == "" {
		msgCD := "BFWE01060101"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.MessageSection) {
		msgCD := "BFWE01060102"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.StringArrayContainsV1([]string{"01", "02", "03"}, ipt.MessageSection) {
		msgCD := "BFWE01060102"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.MessageSection == "02" {
		if ipt.MemberID == "" {
			msgCD := "BFWE01060111"
			return n.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchNumberStringV1(ipt.MemberID) {
			msgCD := "BFWE01060103"
			return n.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.MessageSection == "03" {
		if ipt.ClubID == "" {
			msgCD := "BFWE01060112"
			return n.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchNumberStringV1(ipt.ClubID) {
			msgCD := "BFWE01060104"
			return n.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.Title == "" {
		msgCD := "BFWE01060105"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.Title) > 255 {
		msgCD := "BFWE01060106"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.Content == "" {
		msgCD := "BFWE01060107"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.Content) > 255 {
		msgCD := "BFWE01060108"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}
func (n notification) ValidateAppV1NotificationList(ctx context.Context, ipt input.NotificationList) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/svc_operator.go : ValidateAppV1NotificationList : func (n notification) ValidateAppV1NotificationList(ctx context.Context, ipt input.NotificationList) error start.")))

	if ipt.MessageSection == "" {
		msgCD := "BFWE01060201"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.MessageSection) {
		msgCD := "BFWE01060202"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.StringArrayContainsV1([]string{"01", "02", "03"}, ipt.MessageSection) {
		msgCD := "BFWE01060202"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.SortKey == "" {
		msgCD := "BFWE01060203"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.StringArrayContainsV1([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}, ipt.SortKey) {
		msgCD := "BFWE01060204"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.SortOrder == "" {
		msgCD := "BFWE01060205"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.StringArrayContainsV1([]string{"asc", "desc"}, ipt.SortOrder) {
		msgCD := "BFWE01060206"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PageNum == "" {
		msgCD := "BFWE01060207"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptPageNum := strings.Trim(ipt.PageNum, " ")
	if !util.MatchStringV1("^[0-9]*$", iptPageNum) {
		msgCD := "BFWE01060208"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	pageInt, err := strconv.Atoi(ipt.PageNum)
	if err != nil {
		msgCD := "BFWE01060208"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if pageInt <= 0 {
		msgCD := "BFWE01060208"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PageSize == "" {
		msgCD := "BFWE01060209"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	iptPageSize := strings.Trim(ipt.PageSize, " ")
	if !util.MatchStringV1("^[0-9]*$", iptPageSize) {
		msgCD := "BFWE01060210"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	maximumRecordsPerPageInt, err := strconv.Atoi(ipt.PageSize)
	if err != nil {
		msgCD := "BFWE01060210"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if maximumRecordsPerPageInt <= 0 {
		msgCD := "BFWE01060210"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if maximumRecordsPerPageInt > config.GetProjectConfig().AppEnvConf.App.Operator.MaximumRecordsPerPage01 {
		msgCD := "BFWE01060210"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.MemberID) {
		msgCD := "BFWE01060211"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.MemberID) > 36 {
		msgCD := "BFWE01060211"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubID) {
		msgCD := "BFWE01060212"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.ClubID) > 12 {
		msgCD := "BFWE01060212"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// 3.4 パターンチェック
	if ipt.MessageSection == "02" && ipt.MemberID == "" {
		msgCD := "BFWE01060215"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.MessageSection == "03" && ipt.ClubID == "" {
		msgCD := "BFWE01060216"
		return n.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}
