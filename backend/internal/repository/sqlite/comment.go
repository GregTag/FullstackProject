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
	result := r.db.Create(comment)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CommentSQLite) Update(comment *entity.Comment) error {
	result := r.db.Model(comment).Updates(comment)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CommentSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Comment{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CommentSQLite) Get(id uint) (*entity.Comment, error) {
	var comment entity.Comment

	result := r.db.First(&comment, id)
	if result.Error == nil {
		return &comment, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &comment, entity.ErrCommentNotFound
	} else {
		return &comment, result.Error
	}
}

func (r *CommentSQLite) LoadAll(media_id uint) (*[]entity.CommentView, error) {
	var comments []entity.CommentView

	result := r.db.Table("comment_view").Where("media_id = ?", media_id).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &comments, nil
	}
}
