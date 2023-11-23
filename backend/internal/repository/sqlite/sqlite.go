package repository_sqlite

import (
	"backend/internal/entity"
	"fmt"

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
