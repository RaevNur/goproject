package question

import model "forum/internal/models"

func (s *QuestionService) Create(question *model.Question) error {
	return nil
}

func (s *QuestionService) GetById(id int64) (*model.Question, error) {
	return nil, nil
}

func (s *QuestionService) GetMostViewed() (*model.Question, error) {
	return nil, nil
}

func (s *QuestionService) GetMostLiked() (*model.Question, error) {
	return nil, nil
}

func (s *QuestionService) GetByTag(tagId int64) ([]*model.Question, error) {
	return nil, nil
}

func (s *QuestionService) GetRecentQuestions(page int) ([]*model.Question, error) {
	return nil, nil
}

func (s *QuestionService) Viewed(id int64) error {
	return nil
}
