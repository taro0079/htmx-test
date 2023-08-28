package repository

import (
	"database/sql"
	"taro/test/src/models"
)

type User struct {
    UserId int
    Name string
    Email string
}


func Save(u models.User) {

}

func FindByUserId(user_id int64) User {
	db, _ := sql.Open("mysql", "root@tcp(localhost:3306)/test")
    row := db.QueryRow(`select * from user where user_id = ?`, user_id)
    u := &User{}
    row.Scan(&u.UserId, &u.Name, &u.Email)
    return *u
}

func FindAllUser() {
    db, _ := sql.Open("mysql", "root@tcp(localhost:3306)/test")
    rows := db.Query(`select * from user`)

}
