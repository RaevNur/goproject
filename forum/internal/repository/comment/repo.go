package comment

import (
	"fmt"

	"forum/configs"
	"forum/internal/helper"

	model "forum/internal/models"
)

func (c *CommentRepo) Create(postId int64, comment *model.Comment) error {
	query := `INSERT INTO comments (
		content, 
		user_id, 
		created_at,
		post_id
	) 
	VALUES (?, ?, ?, ?);`

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	encodedTime := helper.EncodeTime((*comment).CreatedAt, configs.TimeFormatRFC3339)
	res, err := stmt.Exec((*comment).Content, (*comment).UserId, encodedTime, postId)
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	(*comment).Id = lastId
	return nil
}

// take page???
// ordered by created time (asc)
func (c *CommentRepo) GetPostComments(postId int64) ([]*model.Comment, error) {
	query := `SELECT id, content, user_id, created_at FROM comments 
	WHERE post_id = ?
	ORDER BY created_at ASC`

	rows, err := c.db.Query(query, postId)
	if err != nil {
		return nil, fmt.Errorf("CommentRepo.GetPostComments: %w", err)
	}

	comments := make([]*model.Comment, 0)
	for rows.Next() {
		t := model.Comment{}
		var decodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Content,
			&t.UserId,
			&decodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("CommentRepo.GetPostComments: %w", err)
		}

		t.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("CommentRepo.GetPostComments: %w", err)
		}
		comments = append(comments, &t)
	}

	return comments, nil
}
