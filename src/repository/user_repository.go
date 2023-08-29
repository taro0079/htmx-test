package repository

import (
	"database/sql"
	"taro/test/src/models"
)

type User struct {
	UserId int
	Name   string
	Email  string
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

func FindAllUser() []User {
	var users []User
	db, _ := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	rows, _ := db.Query(`select * from user`)
	defer rows.Close()

	for rows.Next() {
		var u User
		rows.Scan(&u.UserId, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users
}

func CreateUser(u User) int64 {
	db, _ := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	ins, _ := db.Prepare("INSERT INTO user (name,email) VALUES (?,?)")
	defer ins.Close()
	res, _ := ins.Exec(u.Name, u.Email)
	lastInsertUserId, _ := res.LastInsertId()
	return lastInsertUserId
}
func DeleteUser(u User) int64 {
	db, _ := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	del, _ := db.Prepare("DELETE FROM user WHERE user_id = ?")
    defer del.Close()

    res, _ := del.Exec(u.UserId)
    delId,_ := res.LastInsertId()
    return delId
}
