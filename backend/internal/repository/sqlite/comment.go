package repository_sqlite

import (
	"backend/internal/entity"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CommentSQLite struct {
	db *gorm.DB
}

func NewCommentSQLite(db *gorm.DB) *CommentSQLite {
	return &CommentSQLite{db: db}
}

func (r *CommentSQLite) Create(comment *entity.Comment) error {
	err := r.db.Model(comment).Omit("Sender").Create(comment).Error
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
	result := r.db.Model(comment).Omit("Sender").Updates(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrCommentNotFound
	}
	return nil
}

func (r *CommentSQLite) Delete(id, senderID uint) error {
	result := r.db.Model(&entity.Comment{}).Where("id = ? AND sender_id = ?", id, senderID).Delete(&entity.Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrCommentNotFound
	}
	return nil
}

func (r *CommentSQLite) Load(id, senderID uint) (*entity.CommentView, error) {
	var comment entity.Comment

	result := r.db.Model(&entity.Comment{}).Preload("Sender").Where("id = ? AND sender_id = ?", id, senderID).First(&comment)
	if result.RowsAffected == 0 {
		return nil, entity.ErrCommentNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	var view entity.CommentView
	err := copier.Copy(&view, &comment)
	if err != nil {
		return nil, err
	}
	return &view, nil

}

func (r *CommentSQLite) LoadAll(mediaID uint) (*[]entity.CommentView, error) {
	var comments []entity.Comment

	result := r.db.Model(&entity.Comment{}).Preload("Sender").Where("media_id = ?", mediaID).Find(&comments)
	if result.RowsAffected == 0 {
		return nil, entity.ErrCommentNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	var views []entity.CommentView
	err := copier.Copy(&views, &comments)
	if err != nil {
		return nil, err
	}
	return &views, nil
}
