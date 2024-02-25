package scanErr

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

func IdentifyErr(sourceErr error) (string, error) {
	var pgErr *pgconn.PgError
	var errMsg string

	if errors.As(sourceErr, &pgErr) {
		errMsg = fmt.Sprint("PostgreSQL error: ", pgErr)
		return errMsg, pgErr
	} else {
		errMsg = fmt.Sprint("Error: ", sourceErr)
		return errMsg, sourceErr
	}
}
