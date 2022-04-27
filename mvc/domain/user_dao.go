package domain

import (
	"fmt"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Mike", LastName: "Tyson", Email: "puncher@mail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Message:    fmt.Sprintf("\nuser %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
