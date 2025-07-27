package services

import (
	"any-api/internal/models"
	"any-api/internal/repository"
)

type ItemService struct {
	Repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) CreateItem(item *models.Item) (*models.Item, error) {
	return s.Repo.CreateItem(item)
}