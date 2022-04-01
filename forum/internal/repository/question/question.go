package question

import "database/sql"

type QuestionRepo struct {
	db *sql.DB
}

func NewQuestionRepo(db *sql.DB) *QuestionRepo {
	return &QuestionRepo{db}
}
