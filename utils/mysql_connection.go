package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	// sql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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
	db = d
}

//GetDB db 반환
func GetDB() *sqlx.DB {
	return db
}
