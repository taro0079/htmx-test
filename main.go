package main

import (
	"net/http"
	"taro/test/src/repository"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string
	UserId string
	Email  string
}

//	var users = []User{
//	    {UserId: "1", Name: "森田太郎", Email: "morita0079@gmail.com"},
//	    {UserId: "2", Name: "森田文人", Email: "morita0081@gmail.com"},
//	}
func getUsers(c *gin.Context) {
	var users = repository.FindByUserId(1)
	c.JSON(http.StatusOK, users)

}
func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run()
}
