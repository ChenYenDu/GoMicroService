package domain

import (
	"fmt"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Fede", LastName: "Leon", Email: "tm1@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationEror) {

	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationEror{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
