package util

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsMysqlError(err error, id uint16) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == id
}
