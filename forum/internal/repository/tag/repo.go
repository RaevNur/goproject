package tag

import (
	"fmt"

	"forum/configs"

	model "forum/internal/models"
)

func (t *TagRepo) Create(tag *model.Tag) error {
	query := `INSERT INTO tags (
		name
	) 
	VALUES (?);`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	res, err := stmt.Exec((*tag).Name)
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	(*tag).Id = lastId
	return nil
}

// creates realtion in tags_questions table
func (t *TagRepo) CreateRelation(questionId int64, tag *model.Tag) error {
	query := `INSERT INTO tags_questions (
		tag_ig,
		question_id
	) 
	VALUES (?, ?);`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	res, err := stmt.Exec((*tag).Id, questionId)
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	return nil
}

// gets tags ordered by it's amount(desc) and name(asc)
func (t *TagRepo) GetTags(page int) ([]*model.Tag, error) {
	query := `SELECT tags.id, tags.name, COUNT(tg.question_id) AS "amount" FROM tags 
	INNER JOIN tags_questions AS tg ON tags.id = tg.tag_id 
	GROUP BY tags.id
	ORDER BY amount DESC, tags.name ASC 
	LIMIT ? OFFSET ?`

	offset := (page - 1) * configs.LimitTagsPerPage
	rows, err := t.db.Query(query, configs.LimitTagsPerPage, offset)
	if err != nil {
		return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
	}

	tags := make([]*model.Tag, 0, configs.LimitTagsPerPage)
	for rows.Next() {
		t := model.Tag{}

		err = rows.Scan(
			&t.Id,
			&t.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}

func (t *TagRepo) GetTagsByQuestion(questionId int64) ([]*model.Tag, error) {
	query := `SELECT tags.id, tags.name FROM tags 
	INNER JOIN tags_questions AS tg ON tags.id = tg.tag_id
	WHERE tg.question_id = ? 
	ORDER BY tags.name ASC`

	rows, err := t.db.Query(query, questionId)
	if err != nil {
		return nil, fmt.Errorf("TagRepo.GetTagsByQuestion: %w", err)
	}

	tags := make([]*model.Tag, 0)
	for rows.Next() {
		t := model.Tag{}

		err = rows.Scan(
			&t.Id,
			&t.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("TagRepo.GetTagsByQuestion: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
