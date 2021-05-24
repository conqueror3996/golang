package gormdb

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"

	"gorm.io/gorm"

	apperrorsub "github.com/nri4nudge/api-common/adapter/error"
	apperror "github.com/nri4nudge/api-common/infra/error"
	"github.com/nri4nudge/api-common/infra/gormdatabase"
	"github.com/nri4nudge/api-common/logger/apllog"
	apllogdefaultutil "github.com/nri4nudge/api-common/logger/apllog/defaultutil"
)

var dbm *gormdatabase.DBM

var errm apperror.ErrorManager

func init() {
	dbm = gormdatabase.GetDBM()
	errm = apperrorsub.NewErrorManager()
}

// begin begin transaction
func begin(ctx context.Context) (context.Context, error) {
	return (*dbm).Begin(ctx)
}

// rollback rollback transaction
func rollback(ctx context.Context) (context.Context, error) {
	return (*dbm).Rollback(ctx)
}

// commit commit transaction
func commit(ctx context.Context) (context.Context, error) {
	return (*dbm).Commit(ctx)
}

// table ...
func table(ctx context.Context, name string) *gorm.DB {
	return (*dbm).Table(ctx, name)
}

// gmodel ...
func gmodel(ctx context.Context, value interface{}) *gorm.DB {
	return (*dbm).Model(ctx, value)
}

// find ...
func find(ctx context.Context, out interface{}, where ...interface{}) *gorm.DB {
	return (*dbm).Find(ctx, out, where)
}

// exec ...
func exec(ctx context.Context, sql string, values ...interface{}) *gorm.DB {
	return (*dbm).Exec(ctx, sql, values)
}

// first ...
func first(ctx context.Context, out interface{}, where ...interface{}) *gorm.DB {
	return (*dbm).First(ctx, out, where)
}

// raw ...
func raw(ctx context.Context, sql string, values ...interface{}) *gorm.DB {
	return (*dbm).Raw(ctx, sql, values)
}

// create ...
func create(ctx context.Context, value interface{}) *gorm.DB {
	return (*dbm).Create(ctx, value)
}

// save ...
func save(ctx context.Context, value interface{}) *gorm.DB {
	return (*dbm).Save(ctx, value)
}

// delete ...
func delete(ctx context.Context, value interface{}) *gorm.DB {
	return (*dbm).Delete(ctx, value)
}

// where ...
func where(ctx context.Context, query interface{}, args ...interface{}) *gorm.DB {
	return (*dbm).Where(ctx, query, args)
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// isDatabaseIsClosedError ...
func isDatabaseIsClosedError(err error) bool {
	if err == nil {
		return false
	}
	errString := fmt.Sprintf("%v", err)
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdb/common.go : isDatabaseIsClosedError : errString=%v", errString)))
	return regexp.MustCompile(".*database is closed.*").MatchString(errString)
}

// isDeadlockError ...
func isDeadlockError(err error) bool {
	if err == nil {
		return false
	}
	errString := fmt.Sprintf("%v", err)
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdb/common.go : isDeadlockError : errString=%v", errString)))
	return regexp.MustCompile(".*Error.*1213.*").MatchString(errString)
}

// isLockWaitTimeoutError ...
func isLockWaitTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	errString := fmt.Sprintf("%v", err)
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdb/common.go : isLockWaitTimeoutError : errString=%v", errString)))
	return regexp.MustCompile(".*Error.*1205.*").MatchString(errString)
}

// isDuplicateError ...
func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}
	errString := fmt.Sprintf("%v", err)
	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("adapter/gormdb/common.go : isDuplicateError : errString=%v", errString)))
	return regexp.MustCompile(".*Error.*1062.*").MatchString(errString)
}
