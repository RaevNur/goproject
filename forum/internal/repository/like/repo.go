package like

import (
	"fmt"

	model "forum/internal/models"
)

func (l *LikeRepo) Create(like *model.Like) error {
	query := `INSERT INTO likes (
		user_id, 
		post_id, 
		liked
	) 
	VALUES (?, ?, ?);`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	res, err := stmt.Exec((*like).UserId, (*like).PostId, (*like).Liked)
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	(*like).Id = lastId
	return nil
}

// updates only 'liked' value
func (l *LikeRepo) Update(like *model.Like) error {
	query := `UPDATE likes SET liked = ? WHERE id = ?`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}

	res, err := stmt.Exec((*like).Liked, (*like).Id)
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("LikeRepo.Update affected rows more than 1: %d", affect)
	}

	return nil
}

func (l *LikeRepo) Delete(id int64) error {
	query := `DELETE FROM likes WHERE id = ?`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("LikeRepo.Delete affected rows more than 1: %d", affect)
	}

	return nil
}

func (l *LikeRepo) GetById(id int64) (*model.Like, error) {
	query := `SELECT id, user_id, post_id, liked FROM likes WHERE id = ?`
	row := l.db.QueryRow(query, id)

	like := model.Like{}

	err := row.Scan(&like.Id, &like.UserId, &like.PostId, &like.Liked)
	if err != nil {
		return nil, fmt.Errorf("LikeRepo.GetById: %w", err)
	}

	return &like, nil
}
