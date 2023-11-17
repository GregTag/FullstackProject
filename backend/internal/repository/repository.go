package repository

import (
	"backend/internal/entity"
	repository_sqlite "backend/internal/repository/sqlite"

	"gorm.io/gorm"
)

type Repository struct {
	User       entity.UserRepository
	Media      entity.MediaRepository
	Comment    entity.CommentRepository
	MediaTrack entity.MediaTrackRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       repository_sqlite.NewUserSQLite(db),
		Media:      repository_sqlite.NewMediaSQLite(db),
		Comment:    repository_sqlite.NewCommentSQLite(db),
		MediaTrack: repository_sqlite.NewMediaTrackSQLite(db),
	}
}
