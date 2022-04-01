package models

type Like struct {
	Id     int64
	UserId int64
	PostId int64
	Liked  int
}

type ILikeRepo interface {
	Create(like *Like) error
	Update(like *Like) error
	Delete(id int64) error
	GetById(id int64) (*Like, error)
}
