package repository_sqlite

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type CommentSQLite struct {
	db *gorm.DB
}

func NewCommentSQLite(db *gorm.DB) *CommentSQLite {
	return &CommentSQLite{db: db}
}

func (r *CommentSQLite) Create(comment *entity.Comment) error {
	err := r.db.Model(comment).Create(comment).Error
	if err != nil {
		if checkPrimaryKeyError(err) {
			return entity.ErrCommentAlreadyExists
		} else {
			return err
		}
	}
	return nil
}

func (r *CommentSQLite) Update(comment *entity.Comment) error {
	result := r.db.Model(comment).Updates(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrCommentNotFound
	}
	return nil
}

func (r *CommentSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrCommentNotFound
	}
	return nil
}

func (r *CommentSQLite) Load(id uint) (*entity.CommentView, error) {
	var comment entity.CommentView

	result := r.db.Model(&entity.Comment{}).First(&comment, id)
	if result.Error == nil {
		return &comment, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrCommentNotFound
	} else {
		return nil, result.Error
	}
}

func (r *CommentSQLite) LoadAll(media_id uint) (*[]entity.CommentView, error) {
	var comments []entity.CommentView

	result := r.db.Model(&entity.Comment{}).Where("media_id = ?", media_id).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &comments, nil
	}
}
