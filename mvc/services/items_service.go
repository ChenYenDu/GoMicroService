package services

import (
	"mvc/domain"
	"mvc/utils"
	"net/http"
)

type itemService struct{}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationEror) {
	return nil, &utils.ApplicationEror{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "internal_server_error",
	}
}
