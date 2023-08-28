package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	databaseSourceName := fmt.Sprintf("root@tcp(localhost:3306)/test")
	Db, err = sql.Open("mysql", databaseSourceName)
    if err != nil {
        panic(err)
    }

}
func InitDb() *sql.DB {
	var err error
	databaseSourceName := fmt.Sprintf("root@tcp(localhost:3306)/test")
	Db, err = sql.Open("mysql", databaseSourceName)
    if err != nil {
        panic(err)
    }
    return Db

}
