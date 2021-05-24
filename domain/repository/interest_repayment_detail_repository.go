package repository

import (
	"context"
)

// InterestRepaymentDetailRepository InterestRepaymentDetail's Repository
type InterestRepaymentDetailRepository interface {
	PutInterestRepaymentDetailRepository01(ctx context.Context, userID string, interestRepaymentDate string, interestRepaymentAmount int64, totalInterest int64, totalRepaymentBalance int64, createdAt string, createdPgmID string, updatedAt string, updatedPgmID string) error
}
