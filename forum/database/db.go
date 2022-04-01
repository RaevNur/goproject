package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"forum/configs"

	_ "github.com/mattn/go-sqlite3"
)

// initialize connection with sqlite db
// returns (connection, error)
func InitDB() (*sql.DB, error) {
	err := os.MkdirAll(configs.DB_PATH, os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	dbPathField := filepath.Join(configs.DB_PATH, configs.DB_NAME)
	// connection field from configs
	if configs.DB_USERNAME != "" && configs.DB_PASSWORD != "" {
		authField := fmt.Sprintf("?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s", configs.DB_USERNAME, configs.DB_PASSWORD, configs.DB_AUTHCRYPT)
		dbPathField += authField
	}
	db, err := sql.Open(configs.DBDriverName, dbPathField)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	err = checkDB(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}
	db.SetMaxIdleConns(100)
	return db, err
}

// check the scheme
func checkDB(db *sql.DB) error {
	_, err := db.Exec(createUserTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createSessionsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createTagsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createPostsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createCommentsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createQuestionsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createLikesTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createTagsQuestionsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	_, err = db.Exec(createQuestionsAnswersTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}
	return nil
}
