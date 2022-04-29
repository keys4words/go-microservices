package domain

import (
	"fmt"
	"log"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Mike", LastName: "Tyson", Email: "puncher@mail.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.AppError)
}
type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.AppError) {
	log.Println("We are accessing the DB")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Message:    fmt.Sprintf("\nuser %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
