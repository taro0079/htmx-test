package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func main() {
	var err error
	databaseSourceName := fmt.Sprintf("root@tcp(localhost:3306)/test")
	Db, err = sql.Open("mysql", databaseSourceName)
	if err != nil {
		panic(err)
	}

}
