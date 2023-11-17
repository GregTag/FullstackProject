package entity

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	MediaID   uint
	SenderID  uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentView struct {
	Comment
	SenderLogin    string
	SenderAvatar   string
	SenderFullname string
}

type CommentRepository interface {
	Create(*Comment) error
	Update(*Comment) error
	Delete(id uint) error
	Get(id uint) (*Comment, error)
	LoadAll(media_id uint) (*[]CommentView, error)
}
