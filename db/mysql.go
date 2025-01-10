package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDbMysql() *sql.DB {

	dbCon, err := sql.Open("mysql", os.Getenv("GOOSE_DBSTRING"))

	if err != nil {
		panic(err)
	}

	dbCon.SetConnMaxLifetime(time.Minute * 3)
	dbCon.SetMaxOpenConns(10)
	dbCon.SetMaxIdleConns(10)

	return dbCon

}
