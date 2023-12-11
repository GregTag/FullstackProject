package repository_sqlite

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type MediaSQLite struct {
	db *gorm.DB
}

func NewMediaSQLite(db *gorm.DB) *MediaSQLite {
	return &MediaSQLite{db: db}
}

func (r *MediaSQLite) Create(media *entity.Media) error {
	err := r.db.Create(media).Error
	if err != nil {
		if checkPrimaryKeyError(err) {
			return entity.ErrMediaAlreadyExists
		} else {
			return err
		}
	}
	return nil
}

func (r *MediaSQLite) Update(media *entity.Media) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(media).Association("Genres").Replace(media.Genres)
		if err != nil {
			return err
		}
		result := tx.Model(media).Updates(media)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return entity.ErrMediaNotFound
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *MediaSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Media{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrMediaNotFound
	}
	return nil
}

func (r *MediaSQLite) Get(id uint) (*entity.Media, error) {
	var media entity.Media

	result := r.db.Preload("Genres").First(&media, id)
	if result.Error == nil {
		return &media, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrMediaNotFound
	} else {
		return nil, result.Error
	}
}

func (r *MediaSQLite) Load(id uint) (*entity.MediaView, error) {
	media, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	comments, err := NewCommentSQLite(r.db).LoadAll(id)
	if err != nil {
		return nil, err
	}

	return &entity.MediaView{Media: *media, Comments: *comments}, nil
}

func (r *MediaSQLite) Search(filter *entity.Filter) (*[]entity.MediaResult, error) {
	var results []entity.MediaResult

	query := r.db.Model(&entity.Media{})

	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if len(filter.Genres) > 0 {
		query = query.Joins("JOIN genres ON genres.media_id = media.id").Where("genres.genre IN ?", filter.Genres)
	}
	if filter.YearFrom != 0 {
		query = query.Where("release_year >= ?", filter.YearFrom)
	}
	if filter.YearTo != 0 {
		query = query.Where("release_year <= ?", filter.YearTo)
	}
	if filter.DurationFrom != 0 {
		query = query.Where("duration >= ?", filter.DurationFrom)
	}
	if filter.DurationTo != 0 {
		query = query.Where("duration <= ?", filter.DurationTo)
	}
	if filter.RatingFrom != 0 {
		query = query.Where("cumulative_rating >= number_of_votes * ?", filter.RatingFrom)
	}
	if filter.RatingTo != 0 {
		query = query.Where("cumulative_rating <= number_of_votes * ?", filter.RatingTo)
	}

	err := query.Select("id, title, image, cumulative_rating / number_of_ratings AS rating").Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return &results, nil
}
