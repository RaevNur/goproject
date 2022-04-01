package question

import (
	"forum/internal/models"
)

type QuestionService struct {
	repo models.IQuestionRepo
}

func NewQuestionService(repo models.IQuestionRepo) *QuestionService {
	return &QuestionService{repo}
}
