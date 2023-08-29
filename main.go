package main

import (
	"net/http"
	"taro/test/src/repository"

	"github.com/gin-contrib/cors"
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
	var users = repository.FindAllUser()
	c.JSON(http.StatusOK, users)

}

func createUser(c *gin.Context)  {
    u := repository.User {
        Name: c.PostForm("name"),
        Email: c.PostForm("email"),
    }
    repository.CreateUser(u)
    users:=repository.FindAllUser()
    c.JSON(http.StatusOK, users)
}


func main() {
	router := gin.Default()
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "*",
        },
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTION",
        },
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Access-Control-Allow-Origin",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
            "hx-current-url",
            "hx-request",
            "hx-target",
        },
    }))
	router.GET("/users", getUsers)
    router.POST("/users/new", createUser)
    router.Run(":8080")
}
