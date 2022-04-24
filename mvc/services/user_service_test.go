package services

import (
	"mvc/domain"
	"mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationEror)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationEror) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDataBase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationEror) {
		return nil, &utils.ApplicationEror{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 is not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 is not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationEror) {
		return &domain.User{
			Id:        123,
			FirstName: "Fede",
			LastName:  "Leon",
			Email:     "tm1@gmail.com",
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
