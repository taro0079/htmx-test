package repository

import (
	// "context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRegisteredDriver(t *testing.T) {
	assert.Equal(t, []string{"mysql"}, sql.Drivers())
}

func TestExecuteDDL(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	assert.Nil(t, err)
	defer db.Close()
	_, err = db.Exec(`create table if not exists book(isbn varchar(14), title varchar(200), price int, primary key(isbn))`)
	assert.Nil(t, err)
	_, err = db.Exec(`drop table if exists book`)
	assert.Nil(t, err)
}

func TestExecuteQueryUsingInterpolateParams(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	assert.Nil(t, err)
	defer db.Close()
	_, err = db.Exec(`create table if not exists book(isbn varchar(14), title varchar(200), price int, primary key(isbn))`)
	assert.Nil(t, err)
	// query such as insert
	result, err := db.Exec(`insert into book(isbn, title, price) values(?,?,?)`, "978-4798161488", "MySQL徹底入門 第4版", 4180)
	assert.Nil(t, err)
	rowsAffected, err := result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)

	result, err = db.Exec(`insert into book(isbn, title, price) values(?,?,?)`, "978-4798147406", "詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド", 3960)
	assert.Nil(t, err)
	rowsAffected, err = result.RowsAffected()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)
	_, err = db.Exec(`drop table if exists book`)

}
