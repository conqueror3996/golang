package memory

import (
	"database/sql"

	apperrorsub "github.com/nri4nudge/api-common/adapter/error"
	apperror "github.com/nri4nudge/api-common/infra/error"
)

var errm apperror.ErrorManager

func init() {
	errm = apperrorsub.NewErrorManager()
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
