package models

import "time"

// User -
type User struct {
	Id        int64
	Nickname  string
	Fistname  string
	Lastname  string
	Email     string
	Password  string
	CreatedAt time.Time
	Avatar    string
}

type IUserRepo interface {
	Create(user *User) error
	// Update(user *User) error
	// Delete(id int) error
	GetByID(id int64) (*User, error)
	GetByNickname(nickname string) (*User, error)
	UserExist(nickname, email string) (bool, error)
}

type IUserService interface {
	Register(user *User) error
	Login(user *User) error
	Logout(user *User) error
	GetByID(id int64) (*User, error)
	GetByNickname(nickname string) (*User, error)
	GetLikedQuestions(user *User) ([]*Question, error)
	GetCreatedQuestions(user *User) ([]*Question, error)
}
