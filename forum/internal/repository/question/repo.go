package question

import (
	"fmt"

	"forum/configs"

	model "forum/internal/models"
)

func (q *QuestionRepo) Create(question *model.Question) error {
	query := `INSERT INTO questions (
		post_id, 
		title, 
		views
	) 
	VALUES (?, ?, 0);`

	stmt, err := q.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("QuestionRepo.Create: %w", err)
	}

	res, err := stmt.Exec((*question).Post.Id, (*question).Title)
	if err != nil {
		return fmt.Errorf("QuestionRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("QuestionRepo.Create: %w", err)
	}

	(*question).Id = lastId
	return nil
}

func (q *QuestionRepo) GetById(id int64) (*model.Question, error) {
	query := `SELECT id, post_id, title, views FROM questions WHERE id = ?`
	row := q.db.QueryRow(query, id)

	question := model.Question{}
	postId := int64(0)

	err := row.Scan(&question.Id, &postId, &question.Title, &question.Views)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepo.GetById: %w", err)
	}

	question.Post = &model.Post{
		Id: postId,
	}

	return &question, nil
}

func (q *QuestionRepo) GetMostViewed() (*model.Question, error) {
	query := `SELECT id, post_id, title, views FROM questions ORDER BY views DESC LIMIT 1`
	row := q.db.QueryRow(query)

	question := &model.Question{}
	postId := int64(0)

	err := row.Scan(&question.Id, &postId, &question.Title, &question.Views)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepo.GetMostViewed: %w", err)
	}

	question.Post = &model.Post{
		Id: postId,
	}

	return question, nil
}

// takes a page number (1, 2, 3...)
func (q *QuestionRepo) GetByUserId(userId int64, page int) ([]*model.Question, error) {
	query := `SELECT questions.id, questions.post_id, questions.title, questions.views FROM questions 
	INNER JOIN posts ON questions.post_id = posts.id WHERE posts.user_id = ? 
	ORDER BY posts.created_at DESC LIMIT ? OFFSET ?`

	offset := (page - 1) * configs.LimitQuestionPerPage
	rows, err := q.db.Query(query, userId, configs.LimitQuestionPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepo.GetByUserId: %w", err)
	}

	questions := make([]*model.Question, 0, configs.LimitQuestionPerPage)
	for rows.Next() {
		t := model.Question{}
		postId := int64(0)

		err = rows.Scan(
			&t.Id,
			&postId,
			&t.Title,
			&t.Views,
		)
		if err != nil {
			return nil, fmt.Errorf("QuestionRepo.GetByUserId: %w", err)
		}

		t.Post = &model.Post{
			Id: postId,
		}
		questions = append(questions, &t)
	}

	return questions, nil
}

// takes a page number (1, 2, 3...)
func (q *QuestionRepo) GetByTag(tagId int64, page int) ([]*model.Question, error) {
	query := `SELECT questions.id, questions.post_id, questions.title, questions.views FROM questions 
	INNER JOIN posts ON questions.post_id = posts.id 
	INNER JOIN tags_questions ON questions.id = tags_questions.question_id WHERE tags_questions.tag_ig = ? 
	ORDER BY posts.created_at DESC LIMIT ? OFFSET ?`

	offset := (page - 1) * configs.LimitQuestionPerPage
	rows, err := q.db.Query(query, tagId, configs.LimitQuestionPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepo.GetByTag: %w", err)
	}

	questions := make([]*model.Question, 0, configs.LimitQuestionPerPage)
	for rows.Next() {
		t := model.Question{}
		postId := int64(0)

		err = rows.Scan(
			&t.Id,
			&postId,
			&t.Title,
			&t.Views,
		)
		if err != nil {
			return nil, fmt.Errorf("QuestionRepo.GetByTag: %w", err)
		}

		t.Post = &model.Post{
			Id: postId,
		}
		questions = append(questions, &t)
	}

	return questions, nil
}

// takes a page number (1, 2, 3...)
func (q *QuestionRepo) GetRecentQuestions(page int) ([]*model.Question, error) {
	query := `SELECT questions.id, questions.post_id, questions.title, questions.views FROM questions 
	INNER JOIN posts ON questions.post_id = posts.id ORDER BY posts.created_at DESC LIMIT ? OFFSET ?`

	offset := (page - 1) * configs.LimitQuestionPerPage
	rows, err := q.db.Query(query, configs.LimitQuestionPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("QuestionRepo.GetRecentQuestions: %w", err)
	}

	questions := make([]*model.Question, 0, configs.LimitQuestionPerPage)
	for rows.Next() {
		t := model.Question{}
		postId := int64(0)

		err = rows.Scan(
			&t.Id,
			&postId,
			&t.Title,
			&t.Views,
		)
		if err != nil {
			return nil, fmt.Errorf("QuestionRepo.GetRecentQuestions: %w", err)
		}

		t.Post = &model.Post{
			Id: postId,
		}
		questions = append(questions, &t)
	}

	return questions, nil
}

func (q *QuestionRepo) AddView(id int64) error {
	query := `UPDATE questions SET views = views + 1 WHERE id = ?`
	stmt, err := q.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("QuestionRepo.AddView: %w", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("QuestionRepo.AddView: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("QuestionRepo.AddView: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("QuestionRepo.AddView affected rows more than 1: %d", affect)
	}
	return nil
}
