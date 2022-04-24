package services

import (
	"mvc/domain"
	"mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationEror) {
	user, err := domain.UserDao.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
