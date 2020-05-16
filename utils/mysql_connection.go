package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	// sql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// go test 로 실행
var isTestRun bool = false

func getRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	dirName := filepath.Dir(filename)
	return filepath.Dir(dirName)
}

type connectionJSON struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (c *connectionJSON) read() error {
	rootPath := getRootPath()
	b, err := ioutil.ReadFile(filepath.Join(rootPath, "mysql.connection.json"))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, c); err != nil {
		return err
	}

	return nil
}

var db *sqlx.DB

func init() {
	conn := connectionJSON{}
	err := conn.read()
	if err != nil {
		log.Fatal(err)
	}

	d, err := sqlx.Connect("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conn.UserName,
			conn.Password,
			conn.IP,
			conn.Port,
			conn.Database))

	if err != nil {
		log.Fatal(err)
	}
	if isTestRun {
		d.SetMaxOpenConns(1)
	}
	db = d
}

//GetDB db 반환
func GetDB() *sqlx.DB {
	return db
}

// BeginTransactionForTest 테스트 환경일 때 sql begin transaction 을 실행함
func BeginTransactionForTest() {
	if isTestRun {
		db.MustExec("start transaction;")
	}
}

// RollbackTransactionForTest 테스트 환경일 때 rollback 을 실행함
func RollbackTransactionForTest() {
	if isTestRun {
		db.MustExec("rollback;")
	}
}

func init() {
	if flag.Lookup("test.v") != nil || strings.HasSuffix(os.Args[0], ".test") {
		isTestRun = true
	}
}
