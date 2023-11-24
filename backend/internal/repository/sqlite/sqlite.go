package repository_sqlite

import (
	"backend/internal/entity"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(db_path string) (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(db_path),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Media{},
		&entity.Comment{},
		&entity.MediaTrack{},
		&entity.Genre{},
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database initialized")
	return db, nil
}

func checkPrimaryKeyError(err error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		return (sqliteErr.ExtendedCode & sqlite3.ErrConstraintPrimaryKey) == sqlite3.ErrConstraintPrimaryKey
	}
	return false
}
