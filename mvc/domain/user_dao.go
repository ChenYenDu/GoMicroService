package domain

import (
	"fmt"
	"log"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "tm1@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(UserId int64) (*User, *utils.ApplicationEror)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationEror) {
	log.Println("we are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationEror{
		Message:    fmt.Sprintf("user %v is not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
