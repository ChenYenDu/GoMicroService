package services

import (
	"mvc/domain"
	"mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationEror) {
	return domain.GetUser(userId)
}
