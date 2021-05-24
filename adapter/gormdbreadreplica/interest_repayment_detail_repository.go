package gormdbreadreplica

import (
	"context"
	"errors"
	"fmt"

	"bff-web/domain/repository"

	"github.com/nri4nudge/api-common/contextw"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

// NewInterestRepaymentDetailRepository get repository
func NewInterestRepaymentDetailRepository() repository.InterestRepaymentDetailRepository {
	return interestRepaymentDetailRepository{}
}

// interestRepaymentDetailRepository InterestRepaymentDetail's Repository Sub
type interestRepaymentDetailRepository struct {
}

// PutInterestRepaymentDetailRepository01 insert interest_repayment_details Data
func (r interestRepaymentDetailRepository) PutInterestRepaymentDetailRepository01(ctx context.Context, userID string, interestRepaymentDate string, interestRepaymentAmount int64, totalInterest int64, totalRepaymentBalance int64, createdAt string, createdPgmID string, updatedAt string, updatedPgmID string) error {
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : func (r interestRepaymentDetailRepository) PutInterestRepaymentDetailRepository01(ctx context.Context, interestRepaymentDetailID string, interestRepaymentDetailName string, interestRepaymentDetailNameKana string, birthdate string, phoneNumber1 string, phoneNumber2 string, mailAddress string, postalCode string, address string, branchNumber string, interestRepaymentDetailMemo string, operatorID string, interestRepaymentDetailRegistrationDate string, interestRepaymentDetailDeletedDate string, interestRepaymentDetailStatus string, createDate string, updateDate string, updatePgmID string, requestID string, version string) error start.")))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : userID=%s", userID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : interestRepaymentDate=%s", interestRepaymentDate)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : interestRepaymentAmount=%d", interestRepaymentAmount)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : totalInterest=%d", totalInterest)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : totalRepaymentBalance=%d", totalRepaymentBalance)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : createdAt=%s", createdAt)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : createdPgmID=%s", createdPgmID)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : updatedAt=%s", updatedAt)))
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdbreadreplica/interest_repayment_detail_repository.go : PutInterestRepaymentDetailRepository01 : updatedPgmID=%s", updatedPgmID)))

	return errm.Wrap1(
		errors.New(""),
		contextw.GetRequestLang(ctx),
		"250000500900207",
		"",
	)
}
