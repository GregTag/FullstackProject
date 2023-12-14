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
		err := tx.Omit("Media").Create(track).Error
		if err != nil {
			if checkPrimaryKeyError(err) {
				return entity.ErrMediaTrackAlreadyExists
			} else {
				return err
			}
		}
		err = tx.Model(track).Association("Media").Find(&track.Media)
		if err != nil {
			return err
		}
		if track.Media.ID == 0 {
			return entity.ErrMediaNotFound
		}

		if track.Rating != 0 {
			track.Media.CumulativeRating += int32(track.Rating)
			track.Media.NumberOfRatings += 1
		}
		track.Media.NumberOfTracks += 1

		err = tx.Model(track.Media).Select("cumulative_rating", "number_of_ratings", "number_of_tracks").Updates(track.Media).Error
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
		err = tx.Model(track).Omit("Media").Updates(track).Error
		if err != nil {
			return err
		}
		err = tx.Model(track).Association("Media").Find(&track.Media)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return entity.ErrMediaNotFound
			} else {
				return err
			}
		}

		track.Media.CumulativeRating += int32(track.Rating) - int32(rating)
		if rating == 0 {
			track.Media.NumberOfRatings += 1
		}
		if track.Rating == 0 {
			track.Media.NumberOfRatings -= 1
		}

		result := tx.Model(track.Media).Select("cumulative_rating", "number_of_ratings", "number_of_tracks").Updates(track.Media)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return entity.ErrMediaNotFound
		}
		return nil
	})
	return err
}

func (r *MediaTrackSQLite) Delete(userID uint, mediaID uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		track, err := getImpl(tx, userID, mediaID)
		if err != nil {
			return err
		}

		if track.Rating != 0 {
			track.Media.CumulativeRating -= int32(track.Rating)
			track.Media.NumberOfRatings -= 1
		}
		track.Media.NumberOfTracks -= 1

		err = tx.Model(track.Media).Select("cumulative_rating", "number_of_ratings", "number_of_tracks").Updates(track.Media).Error
		if err != nil {
			return err
		}

		result := tx.Delete(track)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return entity.ErrMediaNotFound
		}
		return nil
	})
	return err
}

func (r *MediaTrackSQLite) Get(userID uint, mediaID uint) (*entity.MediaTrack, error) {
	return getImpl(r.db, userID, mediaID)
}

func getImpl(db *gorm.DB, userID uint, mediaID uint) (*entity.MediaTrack, error) {
	var track entity.MediaTrack

	result := db.Model(&entity.MediaTrack{}).Preload("Media").Where("user_id = ? AND media_id = ?", userID, mediaID).First(&track)
	if result.Error == nil {
		return &track, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &track, entity.ErrMediaTrackNotFound
	} else {
		return &track, result.Error
	}
}

func (r *MediaTrackSQLite) LoadAll(userID uint) (*[]entity.MediaTrackView, error) {
	var tracks []entity.MediaTrackView

	result := r.db.Table("media_tracks").Select("media_tracks.*, media.title as media_title").Joins("JOIN media ON media_tracks.media_id = media.id").Where("media_tracks.user_id = ?", userID).Scan(&tracks)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &tracks, nil
	}
}
