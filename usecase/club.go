package usecase

import (
	"bff-web/usecase/input"
	"bff-web/util"
	"context"
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

// Club ...
type Club interface {
	ValidateAppV1ClubUpdate(ctx context.Context, ipt input.ClubUpdate) error
	ValidateAppV1RewardRegistration(ctx context.Context, ipt input.RewardRegistration) error
	ValidateAppV1RewardUpdate(ctx context.Context, ipt input.RewardUpdate) error
	ValidateAppV1RewardList(ctx context.Context, ipt input.RewardInquiry) error
}

// NewClub ...
func NewClub(
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	errm apperror.ErrorManager,
) Club {
	return club{
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		errm,
	}
}

type club struct {
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	errm                                   apperror.ErrorManager
}

func (u club) ValidateAppV1ClubUpdate(ctx context.Context, ipt input.ClubUpdate) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/club.go : ValidateAppV1ClubUpdate : func (u club) ValidateAppV1ClubUpdate(ctx context.Context, ipt input.ClubUpdate) error start.")))

	if !util.MatchNumberStringV1(ipt.ClubID) || utf8.RuneCountInString(ipt.ClubID) > 12 {
		msgCD := "BFWE01040201"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.ClubID == "" {
		msgCD := "BFWE01040202"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ContentType) {
		msgCD := "BFWE01040203"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.ClubImageURL) || utf8.RuneCountInString(ipt.ClubImageURL) > 2048 {
		msgCD := "BFWE01040204"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.CoverImageURL) || utf8.RuneCountInString(ipt.CoverImageURL) > 2048 {
		msgCD := "BFWE01040205"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.ClubName) || utf8.RuneCountInString(ipt.ClubName) > 255 {
		msgCD := "BFWE01040206"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubOwnerID) || utf8.RuneCountInString(ipt.ClubOwnerID) > 12 {
		msgCD := "BFWE01040207"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ThemeColor) || utf8.RuneCountInString(ipt.ThemeColor) > 12 {
		msgCD := "BFWE01040209"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.CategoryCode) || utf8.RuneCountInString(ipt.ThemeColor) > 12 {
		msgCD := "BFWE01040210"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ReleaseStatus) || utf8.RuneCountInString(ipt.ReleaseStatus) > 1 {
		msgCD := "BFWE01040211"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ReleaseStartDate) || utf8.RuneCountInString(ipt.ReleaseStartDate) > 8 {
		msgCD := "BFWE01040212"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ReleaseEndDate) || utf8.RuneCountInString(ipt.ReleaseEndDate) > 8 {
		msgCD := "BFWE01040213"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.MaxMembers) || utf8.RuneCountInString(ipt.MaxMembers) > 10 {
		msgCD := "BFWE01040214"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubParticipants) || utf8.RuneCountInString(ipt.ClubParticipants) > 12 {
		msgCD := "BFWE01040215"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.MaxMembershipNo) || utf8.RuneCountInString(ipt.MaxMembershipNo) > 12 {
		msgCD := "BFWE01040216"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u club) ValidateAppV1RewardUpdate(ctx context.Context, ipt input.RewardUpdate) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/club.go : ValidateAppV1ClubUpdate : func (u club) ValidateAppV1ClubUpdate(ctx context.Context, ipt input.ClubUpdate) error start.")))

	if !util.MatchNumberStringV1(ipt.RewardID) || utf8.RuneCountInString(ipt.RewardID) > 12 {
		msgCD := "BFWE01040701"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.RewardID == "" {
		msgCD := "BFWE01040702"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.ContentType == "" {
		if ipt.RewardID != "" && ipt.ClubID == "" && ipt.RewardName == "" && ipt.DistributionPeriod == "" && ipt.DistributionStatus == "" &&
			ipt.TotalOrderQuantity == "" && ipt.DownloadURL == "" && ipt.RewardContent == "" && ipt.RewardImage == "" && ipt.RewardThumbnail == "" {
			msgCD := "BFWE01040715"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if !util.MatchNumberStringV1(ipt.ContentType) {
		msgCD := "BFWE01040703"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubID) {
		msgCD := "BFWE01040704"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.RewardName) || utf8.RuneCountInString(ipt.RewardName) > 255 {
		msgCD := "BFWE01040705"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchStringV1("^[0-9\\-]+$", ipt.DistributionPeriod) {
		msgCD := "BFWE01040706"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.DistributionStatus) || utf8.RuneCountInString(ipt.DistributionStatus) > 1 {
		msgCD := "BFWE01040707"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.TotalOrderQuantity) {
		msgCD := "BFWE01040708"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.DownloadURL) || utf8.RuneCountInString(ipt.DownloadURL) > 2048 {
		msgCD := "BFWE01040709"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.RewardImage) || utf8.RuneCountInString(ipt.RewardImage) > 2048 {
		msgCD := "BFWE010407011"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.RewardThumbnail) {
		msgCD := "BFWE010407012"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u club) ValidateAppV1RewardRegistration(ctx context.Context, ipt input.RewardRegistration) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/club.go : ValidateAppV1RewardRegistration : func (u club) ValidateAppV1RewardRegistration(ctx context.Context, ipt input.RewardRegistration) error start.")))

	if ipt.ClubID == "" {
		msgCD := "BFWE01040601"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubID) || utf8.RuneCountInString(ipt.ClubID) > 12 {
		msgCD := "BFWE01040602"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.RewardName == "" {
		msgCD := "BFWE01040603"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.RewardName) > 255 {
		msgCD := "BFWE01040604"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.Section == "" {
		msgCD := "BFWE01040605"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.Section) || utf8.RuneCountInString(ipt.Section) > 1 {
		msgCD := "BFWE01040606"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.AchievementAmount == "" {
		msgCD := "BFWE01040607"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.AchievementAmount) || utf8.RuneCountInString(ipt.AchievementAmount) > 20 {
		msgCD := "BFWE01040608"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.Created == "" {
		msgCD := "BFWE01040609"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.Created) || utf8.RuneCountInString(ipt.Created) > 255 {
		msgCD := "BFWE01040610"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.ShipmentSource == "" {
		msgCD := "BFWE01040611"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ShipmentSource) || utf8.RuneCountInString(ipt.ShipmentSource) > 255 {
		msgCD := "BFWE01040612"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchStringV1("^[0-9]{8}-[0-9]{8}+$", ipt.DistributionPeriod) {
		msgCD := "BFWE01040613"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.DistributionStatus) || utf8.RuneCountInString(ipt.ShipmentSource) > 1 {
		msgCD := "BFWE01040614"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.TotalOrderQuantity) {
		msgCD := "BFWE01040615"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	// if ipt.RewardContent == "" {
	// 	msgCD := "BFWE01040616"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	// if utf8.RuneCountInString(ipt.RewardContent) > 255 {
	// 	msgCD := "BFWE01040617"
	// 	return u.errm.Wrap1(
	// 		errors.New(""),
	// 		contextw.GetRequestLang(ctx),
	// 		msgCD,
	// 		"",
	// 	)
	// }

	if !util.MatchAlphanumericSymbolStringV1(ipt.RewardImage) || utf8.RuneCountInString(ipt.RewardImage) > 2048 {
		msgCD := "BFWE01040618"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.RewardThumbnail) || utf8.RuneCountInString(ipt.RewardThumbnail) > 2048 {
		msgCD := "BFWE01040619"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u club) ValidateAppV1RewardList(ctx context.Context, ipt input.RewardInquiry) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("usecase/club.go : ValidateAppV1ClubUpdate : func (u club) ValidateAppV1ClubUpdate(ctx context.Context, ipt input.ClubUpdate) error start.")))

	if ipt.ClubID == "" {
		msgCD := "BFWE01040801"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubID) || utf8.RuneCountInString(ipt.ClubID) > 12 {
		msgCD := "BFWE01040802"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}
