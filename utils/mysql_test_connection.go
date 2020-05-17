package utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// SQLDB 테스트용과 서비스용 둘다 사용가능한 인터페이스
type SQLDB interface {
	sqlx.Ext
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	MustBegin() *sqlx.Tx
	Beginx() (*sqlx.Tx, error)
}

func init() {
	mysql.Config
}
