package services

import (
	"mvc/domain"
	"mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.AppError)
)

type usersDaoMock struct{}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.AppError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.AppError) {
		return nil, &utils.AppError{
			StatusCode: http.StatusNotFound,
			Message:    "\nuser 0 not found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "\nuser 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.AppError) {
		return &domain.User{Id: 1}, nil
	}
	user, err := UsersService.GetUser(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
}
