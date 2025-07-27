package services

import (
	"any-api/internal/models"
	"any-api/internal/repository"
)

type HelloService struct {
	Repo repository.MessageRepository
}

func NewHelloService(repo repository.MessageRepository) *HelloService {
	return &HelloService{Repo: repo}
}

func (s *HelloService) GetHelloMessage() *models.Message {
	return s.Repo.GetMessage()
}