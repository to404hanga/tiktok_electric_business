package errs

import "github.com/go-sql-driver/mysql"

const MYSQL_DUPLICATE_ERROR uint16 = 1062

func IsDuplicateError(err error) bool {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == MYSQL_DUPLICATE_ERROR {
			return true
		}
	}
	return false
}
