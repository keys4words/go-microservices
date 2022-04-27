package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Mike", LastName: "Tyson", email: "puncher@mail.com"},
	}
)

func GetUser(userId int64) (User, error) {
	user := users[userId]
	if user == nil {
		return User{}, errors.New(fmt.Sprintf("user %v not found", userId))
	}
}
