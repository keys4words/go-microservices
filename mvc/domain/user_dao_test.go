package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting user with this id")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "\nuser 0 not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Mike", user.FirstName)
	assert.EqualValues(t, "Tyson", user.LastName)
	assert.EqualValues(t, "puncher@mail.com", user.Email)
}
