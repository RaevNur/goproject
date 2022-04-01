package user

import (
	"fmt"

	"forum/configs"
	"forum/internal/helper"

	model "forum/internal/models"
)

func (u *UserRepo) Create(user *model.User) error {
	query := `INSERT INTO users (
		nickname, 
		email, 
		created_at, 
		password, 
		firstname, 
		lastname
	) 
	VALUES (?, ?, ?, ?, ?, ?);`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	encodedTime := helper.EncodeTime((*user).CreatedAt, configs.TimeFormatRFC3339)
	res, err := stmt.Exec((*user).Nickname, (*user).Email, encodedTime, (*user).Password, (*user).Fistname, (*user).Lastname)
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	(*user).Id = lastId
	return nil
}

func (u *UserRepo) GetByID(id int64) (*model.User, error) {
	query := `SELECT id, nickname, email, created_at, firstname, lastname FROM users WHERE id = ?`
	row := u.db.QueryRow(query, id)

	user := model.User{}
	decodedTime := ""

	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &decodedTime, &user.Fistname, &user.Lastname)
	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetById: %w", err)
	}

	user.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetById: %w", err)
	}

	return &user, nil
}

func (u *UserRepo) GetByNickname(nickname string) (*model.User, error) {
	query := `SELECT id, nickname, email, created_at, firstname, lastname FROM users WHERE nickname = ?`
	row := u.db.QueryRow(query, nickname)

	user := model.User{}
	var decodedTime string

	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &decodedTime, &user.Fistname, &user.Lastname)
	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetByNickname: %w", err)
	}

	user.CreatedAt, err = helper.DecodeTime(decodedTime, configs.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("UserRepo.GetById: %w", err)
	}

	return &user, nil
}

func (u *UserRepo) UserExist(nickname, email string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE nickname = ? OR email = ?`
	row := u.db.QueryRow(query, nickname, email)

	var count int

	err := row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("UserRepo.GetByNickname: %w", err)
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
