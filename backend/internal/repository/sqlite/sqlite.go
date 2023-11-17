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

	comment_view := db.Model(&entity.Comment{}).Joins("left join users on users.id = comments.sender_id").Select("comments.*, users.login as sender_login, users.avatar as sender_avatar, users.fullname as sender_fullname")
	err = db.Migrator().CreateView("comment_view", gorm.ViewOption{Query: comment_view})

	if err != nil {
		fmt.Println("Failed to create CommentView")
	}

	media_view := db.Model(&entity.Media{}).Joins("left join comment_view on media.id = comment_view.media_id")
	err = db.Migrator().CreateView("media_view", gorm.ViewOption{Query: media_view})

	if err != nil {
		fmt.Println("Failed to create MediaView")
	}

	media_track_view := db.Model(&entity.MediaTrack{}).Joins("left join media on media.id = media_tracks.media_id")
	err = db.Migrator().CreateView("media_track_view", gorm.ViewOption{Query: media_track_view})

	if err != nil {
		fmt.Println("Failed to create MediaTrackView")
	}

	return db, nil
}
