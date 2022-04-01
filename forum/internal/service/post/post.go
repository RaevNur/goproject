package post

import "forum/internal/models"

type PostService struct {
	repo models.IPostService
}

func NewPostService(repo models.IPostService) *PostService {
	return &PostService{repo}
}
