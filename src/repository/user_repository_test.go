package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFind(t * testing.T) {
    user := FindByUserId(1)
    assert.Equal(t, 1, user.UserId)
}

func TestFindAllUser(t * testing.T) {
    users := FindAllUser()
    assert.Equal(t, 1, users[0].UserId)
}

func TestCreateUser(t * testing.T) {
    user := User{
        Name: "testUser",
        Email:"test@test.com",
    }
    res := CreateUser(user)
    assert.NotNil(t, res)
    deleteUser := User{
        UserId: int(res),
    }
    DeleteUser(deleteUser)
}

