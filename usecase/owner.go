package usecase

import (
	"bff-web/config"
	"bff-web/usecase/input"
	"bff-web/util"
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/nri4nudge/api-common/contextw"
	commonrepository "github.com/nri4nudge/api-common/domain/repository"
	apperror "github.com/nri4nudge/api-common/infra/error"
)

// Card ...
type Owner interface {
	ValidateAppV1OwnerRegistration(ctx context.Context, ipt input.OwnerRegistration) error
	ValidateAppV1OwnerList(ctx context.Context, ipt input.OwnerList) error
	ValidateAppV1OwnerUpdate(ctx context.Context, ipt input.OwnerUpdate) error
}

// NewOperator ...
func NewOwner(
	commonTransactionRepository commonrepository.TransactionRepository,
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository,
	errm apperror.ErrorManager,
) Owner {
	return owner{
		commonTransactionRepository,
		commonReadReplicaTransactionRepository,
		errm,
	}
}

type owner struct {
	commonTransactionRepository            commonrepository.TransactionRepository
	commonReadReplicaTransactionRepository commonrepository.TransactionRepository
	errm                                   apperror.ErrorManager
}

func (u owner) ValidateAppV1OwnerRegistration(ctx context.Context, ipt input.OwnerRegistration) error {

	if ipt.ClubOwnerName == "" {
		msgCD := "BFWE01030101"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.ClubOwnerName) > 255 {
		msgCD := "BFWE01030102"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.MailAddress == "" {
		msgCD := "BFWE01030103"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}
	if !util.MatchAlphanumericSymbolStringV1(ipt.MailAddress) {
		msgCD := "BFWE01030104"
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
		msgCD := "BFWE01030104"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if ipt.PhoneNumber != "" {
		if !util.MatchStringV1("^[0-9\\-]+$", ipt.PhoneNumber) {
			msgCD := "BFWE01030106"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.CompanyName != "" {
		if utf8.RuneCountInString(ipt.CompanyName) > 255 {
			msgCD := "BFWE01030108"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
	}

	if ipt.InChargeName != "" {
		if utf8.RuneCountInString(ipt.InChargeName) > 100 {
			msgCD := "BFWE01030110"
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

func (u owner) ValidateAppV1OwnerUpdate(ctx context.Context, ipt input.OwnerUpdate) error {

	if ipt.ClubOwnerID == "" {
		msgCD := "BFWE01030201"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.ClubOwnerID) {
		msgCD := "BFWE01030202"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchAlphanumericSymbolStringV1(ipt.ImageURL) {
		msgCD := "BFWE01030203"
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
		msgCD := "BFWE01030204"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if !util.MatchNumberStringV1(ipt.PhoneNumber) {
		msgCD := "BFWE01030205"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.CompanyName) > 255 {
		msgCD := "BFWE01030206"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.InChargeName) > 100 {
		msgCD := "BFWE01030207"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.BankName) > 50 {
		msgCD := "BFWE01030208"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if (!util.MatchAlphanumericSymbolStringV1(ipt.BranchCode)) || (utf8.RuneCountInString(ipt.BranchCode) > 4) {
		msgCD := "BFWE01030209"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.BranchName) > 50 {
		msgCD := "BFWE01030210"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.AccountType) > 50 {
		msgCD := "BFWE01030211"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if (!util.MatchNumberStringV1(ipt.AccountNo)) || (utf8.RuneCountInString(ipt.BranchCode) > 10) {
		msgCD := "BFWE01030212"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.AccountName) > 50 {
		msgCD := "BFWE01030213"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if utf8.RuneCountInString(ipt.AccountNameKana) > 50 {
		msgCD := "BFWE01030214"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if (!util.MatchNumberStringV1(ipt.ReleaseStatus)) || (utf8.RuneCountInString(ipt.ReleaseStatus) > 1) {
		msgCD := "BFWE01030215"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if (!util.MatchNumberStringV1(ipt.ReleaseStartDate)) || (utf8.RuneCountInString(ipt.ReleaseStartDate) > 8) {
		msgCD := "BFWE01030216"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	if (!util.MatchNumberStringV1(ipt.ReleaseEndDate)) || (utf8.RuneCountInString(ipt.ReleaseStartDate) > 8) {
		msgCD := "BFWE01030217"
		return u.errm.Wrap1(
			errors.New(""),
			contextw.GetRequestLang(ctx),
			msgCD,
			"",
		)
	}

	return nil
}

func (u owner) ValidateAppV1OwnerList(ctx context.Context, ipt input.OwnerList) error {

	if ipt.ClubOwnerID == "" {
		if ipt.SearchText != "" && ipt.PageNum == "" && ipt.PageSize == "" && ipt.SortKey == "" && ipt.SortOrder == "" && ipt.Type == "" {
			msgCD := "BFWE01030311"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
		if ipt.Type == "" {
			msgCD := "BFWE01030301"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.MatchNumberStringV1(ipt.Type) {
			msgCD := "BFWE01030301"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortKey == "" {
			msgCD := "BFWE01030302"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}
		if !util.MatchNumberStringV1(ipt.SortKey) {
			msgCD := "BFWE01030302"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SortOrder == "" {
			msgCD := "BFWE01030303"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if !util.StringArrayContainsV1([]string{"desc", "asc"}, ipt.SortOrder) {
			msgCD := "BFWE01030303"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageNum == "" {
			msgCD := "BFWE01030304"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		pageInt, err := strconv.Atoi(ipt.PageNum)
		if err != nil {
			msgCD := "BFWE01030304"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if pageInt <= 0 {
			msgCD := "BFWE01030304"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.PageSize == "" {
			msgCD := "BFWE01030305"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		maximumRecordsPerPageInt, err := strconv.Atoi(ipt.PageSize)
		if err != nil {
			msgCD := "BFWE01030305"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt <= 0 {
			msgCD := "BFWE01030305"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if maximumRecordsPerPageInt > config.GetProjectConfig().AppEnvConf.App.Operator.MaximumRecordsPerPage01 {
			msgCD := "BFWE01030305"
			return u.errm.Wrap1(
				errors.New(""),
				contextw.GetRequestLang(ctx),
				msgCD,
				"",
			)
		}

		if ipt.SearchText != "" {
			if !util.MatchAlphanumericSymbolStringV1(ipt.SearchText) {
				msgCD := "BFWE01030306"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}

		if !util.StringArrayContainsV1([]string{"0", "1"}, ipt.Type) {
			if ipt.SearchText == "" {
				msgCD := "BFWE01030312"
				return u.errm.Wrap1(
					errors.New(""),
					contextw.GetRequestLang(ctx),
					msgCD,
					"",
				)
			}
		}
	} else {
		if ipt.Type != "" || ipt.SortKey != "" || ipt.SortOrder != "" || ipt.PageNum != "" || ipt.PageSize != "" || ipt.SearchText != "" {
			msgCD := "BFWE01030310"
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
