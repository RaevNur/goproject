package models

type Tag struct {
	Id    int64
	Name  string
	Count int
}

type ITagRepo interface {
	Create(tag *Tag) error
	CreateRelation(questionId int64, tag *Tag) error
	GetTags(page int) ([]*Tag, error)
	GetTagsByQuestion(questionId int64) ([]*Tag, error)
}

type ITagService interface {
	GetTags(page int) ([]*Tag, error)
}
