package service

import (
	"social-network-api/pkg/clients/kafka"
	"social-network-api/pkg/repository"
)


type Posts interface {
	Create(userID, title, content string) (string, error)
	GetAll() ([]Post, error)
	GetByID(postID string) (Post, error)
	GetByUserID(userID string) ([]Post, error)
	Update(postID, title, content string) error
	Delete(postID string) error
}

type Users interface {
	GetUserPosts(userID string) ([]Post, error)
}

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Service struct {
	Posts
	Users
	kafkaProducer *kafka.Producer
}

func NewService(repos *repository.Repository, kafkaProducer *kafka.Producer) *Service {
	return &Service{
		kafkaProducer: kafkaProducer,
	}
}
