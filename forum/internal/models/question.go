package models

type Question struct {
	Id      int64
	Title   string
	Views   int
	Post    *Post
	Tags    []*Tag
	Answers []*Post
}

type IQuestionRepo interface {
	Create(question *Question) error
	// Delete(id int) error
	// Update(question *Question) error
	GetById(id int64) (*Question, error)
	GetMostViewed() (*Question, error)
	GetByUserId(userId int64, page int) ([]*Question, error)
	GetByTag(tagId int64, page int) ([]*Question, error)
	GetRecentQuestions(page int) ([]*Question, error)
	AddView(id int64) error
}

type IQuestionService interface {
	Create(question *Question) error
	GetById(id int64) (*Question, error)
	GetMostViewed() (*Question, error)
	GetMostLiked() (*Question, error)
	GetByTag(tagId int64) ([]*Question, error)
	GetRecentQuestions(page int) ([]*Question, error)
	Viewed(id int64) error
}
