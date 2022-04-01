package user

import model "forum/internal/models"

func (s *UserService) Register(user *model.User) error {
	return nil
}

func (s *UserService) Login(user *model.User) error {
	return nil
}

func (s *UserService) Logout(user *model.User) error {
	return nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	return nil, nil
}

func (s *UserService) GetByNickname(nickname string) (*model.User, error) {
	return nil, nil
}

func (s *UserService) GetLikedQuestions(user *model.User) ([]*model.Question, error) {
	return nil, nil
}

func (s *UserService) GetCreatedQuestions(user *model.User) ([]*model.Question, error) {
	return nil, nil
}
