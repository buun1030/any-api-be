package repository

import "any-api/internal/models"

type MessageRepository interface {
	GetMessage() *models.Message
}

type InMemoryMessageRepository struct{}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{}
}

func (r *InMemoryMessageRepository) GetMessage() *models.Message {
	return &models.Message{Text: "Hello from Repository!"}
}