package services

import (
	"mvc/domain"
	"mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.AppError) {
	return nil, &utils.AppError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
