package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/internal/service/question"
	"forum/internal/service/user"
)

type Service struct {
	User     models.IUserService
	Question models.IQuestionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     user.NewUserService(repo.User),
		Question: question.NewQuestionService(repo.Question),
	}
}
