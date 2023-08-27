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

