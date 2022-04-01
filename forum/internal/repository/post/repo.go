package post

import (
	"fmt"

	"forum/configs"
	"forum/internal/helper"

	model "forum/internal/models"
)

func (p *PostRepo) Create(post *model.Post) error {
	query := `INSERT INTO posts (
		content, 
		user_id, 
		created_at
	) 
	VALUES (?, ?, ?);`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	encodedTime := helper.EncodeTime((*post).CreatedAt, configs.TimeFormatRFC3339)
	res, err := stmt.Exec((*post).Content, (*post).UserId, encodedTime)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	(*post).Id = lastId
	return nil
}

// creates relation in questions_answers table
func (p *PostRepo) CreateRelation(post *model.Post, questionId int64) error {
	query := `INSERT INTO questions_answers (
		post_id, 
		question_id
	) 
	VALUES (?, ?);`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("PostRepo.CreateRelation: %w", err)
	}

	res, err := stmt.Exec((*post).Id, questionId)
	if err != nil {
		return fmt.Errorf("PostRepo.CreateRelation: %w", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("PostRepo.CreateRelation: %w", err)
	}

	return nil
}

// returns questions body
func (p *PostRepo) GetQuestion(postId int64) (*model.Post, error) {
	query := `SELECT id, content, user_id, created_at FROM posts WHERE id = ?`
	row := p.db.QueryRow(query, postId)

	post := model.Post{}
	decodedTime := ""

	err := row.Scan(&post.Id, &post.Content, &post.UserId, &decodedTime)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetQuestion: %w", err)
	}

	post.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetQuestion: %w", err)
	}

	return &post, nil
}

// takes a page number (1, 2, 3...)
// ordered by like rate(desc) and created date(asc)
func (p *PostRepo) GetQuestionAnswers(questionId int64, page int) ([]*model.Post, error) {
	query := `SELECT posts.id, posts.content, posts.user_id, posts.created_at, IFNULL(SUM(subq1.rate), 0) as "rate" FROM posts, 
	INNER JOIN questions_answers AS answ ON answ.post_id = posts.id 
	LEFT JOIN (SELECT COUNT(*) AS "rate", post_id FROM likes WHERE liked = ? GROUP BY post_id, liked
		UNION ALL
	SELECT COUNT(*)*(-1) AS "rate", post_id FROM likes WHERE liked = ? GROUP BY post_id, liked) AS subq1 posts.id = subq1.post_id 
	WHERE answ.question_id = ?
	GROUP BY posts.id
	ORDER BY rate DESC, posts.created_at ASC 
	LIMIT ? OFFSET ?`

	offset := (page - 1) * configs.LimitAnswersPerPage
	rows, err := p.db.Query(query, configs.LikeValue, configs.DislikeValue, questionId, configs.LimitAnswersPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetQuestionAnswers: %w", err)
	}

	answers := make([]*model.Post, 0, configs.LimitAnswersPerPage)
	for rows.Next() {
		t := model.Post{}
		var decodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Content,
			&t.UserId,
			&decodedTime,
			&t.Rate,
		)
		if err != nil {
			return nil, fmt.Errorf("PostRepo.GetQuestionAnswers: %w", err)
		}

		t.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("PostRepo.GetQuestionAnswers: %w", err)
		}
		answers = append(answers, &t)
	}

	return answers, nil
}

// returns question id with body
func (p *PostRepo) GetMostLikedQuestion() (int64, *model.Post, error) {
	query := `SELECT questions.id, posts.id, posts.content, posts.user_id, posts.created_at, IFNULL(SUM(subq1.rate), 0) as "rate" FROM posts
	INNER JOIN questions ON questions.post_id = posts.id
	LEFT JOIN (SELECT COUNT(*) AS "rate", post_id FROM likes WHERE liked = ? GROUP BY post_id, liked
		UNION ALL
	SELECT COUNT(*)*(-1) AS "rate", post_id FROM likes WHERE liked = ? GROUP BY post_id, liked) AS subq1 ON posts.id = subq1.post_id
	GROUP BY posts.id
	ORDER BY rate DESC, posts.created_at ASC
	LIMIT 1`
	row := p.db.QueryRow(query, configs.LikeValue, configs.DislikeValue)

	questionId := int64(0)
	post := &model.Post{}
	var decodedTime string

	err := row.Scan(&questionId, &post.Id, &post.Content, &post.UserId, &decodedTime, &post.Rate)
	if err != nil {
		return int64(0), nil, fmt.Errorf("PostRepo.GetMostLikedQuestion: %w", err)
	}

	post.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return int64(0), nil, fmt.Errorf("PostRepo.GetMostLikedQuestion: %w", err)
	}

	return questionId, post, nil
}
