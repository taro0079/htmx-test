package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFind(t * testing.T) {
    user := FindByUserId(1)
    assert.Equal(t, 1, user.UserId)

}
