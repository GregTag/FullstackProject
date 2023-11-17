package repository_sqlite

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type MediaTrackSQLite struct {
	db *gorm.DB
}

func NewMediaTrackSQLite(db *gorm.DB) *MediaTrackSQLite {
	return &MediaTrackSQLite{db: db}
}

func (r *MediaTrackSQLite) Create(track *entity.MediaTrack) error {
	result := r.db.Create(track)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MediaTrackSQLite) Update(track *entity.MediaTrack) error {
	result := r.db.Model(track).Updates(track)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MediaTrackSQLite) Delete(user_id uint, media_id uint) error {
	result := r.db.Delete(&entity.MediaTrack{}, user_id, media_id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *MediaTrackSQLite) Get(user_id uint, media_id uint) (*entity.MediaTrack, error) {
	var track entity.MediaTrack

	result := r.db.First(&track, user_id, media_id)
	if result.Error == nil {
		return &track, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &track, entity.ErrMediaTrackNotFound
	} else {
		return &track, result.Error
	}
}

func (r *MediaTrackSQLite) LoadAll(user_id uint) (*[]entity.MediaTrackView, error) {
	var tracks []entity.MediaTrackView

	result := r.db.Table("media_track_view").Where("user_id = ?", user_id).Find(&tracks)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &tracks, nil
	}
}
