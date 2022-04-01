package session

import (
	"database/sql"
	"fmt"

	"forum/configs"
	"forum/internal/helper"

	model "forum/internal/models"
)

func (s *SessionRepo) Create(session *model.Session) error {
	query := `INSERT INTO sessions (
		uuid, 
		user_id, 
		created_at
	) 
	VALUES (?, ?, ?);`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	encodedTime := helper.EncodeTime((*session).CreatedAt, configs.TimeFormatRFC3339)
	res, err := stmt.Exec((*session).Uuid, (*session).UserId, encodedTime)
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	(*session).Id = lastId
	return nil
}

func (s *SessionRepo) Delete(id int64) error {
	query := `DELETE FROM sessions WHERE id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("SessionRepo.Delete affected rows more than 1: %d", affect)
	}

	return nil
}

func (s *SessionRepo) GetByUserId(userId int64) (*model.Session, error) {
	query := `SELECT id, uuid, user_id, created_at FROM sessions WHERE user_id = ?`
	row := s.db.QueryRow(query, userId)

	session := model.Session{}
	decodedTime := ""

	err := row.Scan(&session.Id, &session.Uuid, &session.UserId, &decodedTime)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUserId: %w", err)
	}

	session.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUserId: %w", err)
	}

	return &session, nil
}

func (s *SessionRepo) GetByUuid(uuid string) (*model.Session, error) {
	query := `SELECT id, uuid, user_id, created_at FROM sessions WHERE uuid = ?`
	row := s.db.QueryRow(query, uuid)

	session := model.Session{}
	decodedTime := ""

	err := row.Scan(&session.Id, &session.Uuid, &session.UserId, &decodedTime)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUuid: %w", err)
	}

	session.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUuid: %w", err)
	}

	return &session, nil
}
