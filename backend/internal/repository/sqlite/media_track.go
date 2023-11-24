package repository_sqlite

import (
	"backend/internal/entity"
	"database/sql"
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
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(track).Error
		if err != nil {
			if checkPrimaryKeyError(err) {
				return entity.ErrMediaTrackAlreadyExists
			} else {
				return err
			}
		}

		if track.Media.ID == 0 {
			return entity.ErrMediaNotFound
		}

		if track.Rating != 0 {
			track.Media.CumulativeRating += uint64(track.Rating)
			track.Media.NumberOfRatings += 1
		}
		track.Media.NumberOfTracks += 1

		err = tx.Model(track.Media).Updates(track.Media).Error
		return err
	})
	return err

}

func (r *MediaTrackSQLite) Update(track *entity.MediaTrack) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var rating uint8
		err := tx.Model(&entity.MediaTrack{}).Select("rating").Where("user_id = ? AND media_id = ?", track.UserID, track.MediaID).Row().Scan(&rating)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return entity.ErrMediaTrackNotFound
			} else {
				return err
			}
		}
		err = tx.Model(track).Updates(track).Error
		if err != nil {
			return err
		}

		track.Media.CumulativeRating += uint64(track.Rating - rating)
		if rating == 0 {
			track.Media.NumberOfRatings += 1
		}
		if track.Rating == 0 {
			track.Media.NumberOfRatings -= 1
		}

		result := tx.Model(track.Media).Updates(track.Media)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return entity.ErrUserNotFound
		}
		return nil
	})
	return err
}

func (r *MediaTrackSQLite) Delete(user_id uint, media_id uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		track, err := r.Get(user_id, media_id)
		if err != nil {
			return err
		}

		if track.Rating != 0 {
			track.Media.CumulativeRating -= uint64(track.Rating)
			track.Media.NumberOfRatings -= 1
		}
		track.Media.NumberOfTracks -= 1

		err = tx.Model(track.Media).Updates(track.Media).Error
		if err != nil {
			return err
		}

		result := tx.Delete(track)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return entity.ErrUserNotFound
		}
		return nil
	})
	return err
}

func (r *MediaTrackSQLite) Get(user_id uint, media_id uint) (*entity.MediaTrack, error) {
	var track entity.MediaTrack

	result := r.db.Model(&entity.MediaTrack{}).Preload("Media").Where("user_id = ? AND media_id = ?", user_id, media_id).First(&track)
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

	result := r.db.Table("media_tracks").Select("media_tracks.*, media.title as media_title").Joins("JOIN media ON media_tracks.media_id = media.id").Where("media_tracks.user_id = ?", user_id).Scan(&tracks)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &tracks, nil
	}
}
