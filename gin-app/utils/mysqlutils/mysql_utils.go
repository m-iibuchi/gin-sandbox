package mysqlutils

import (
	"gin-sandbox/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	noSearchResult = "record not found"
)

func ParseError(err error) *errors.ApiErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noSearchResult) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
