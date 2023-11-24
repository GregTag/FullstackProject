package entity

import "time"

type CommentBase struct {
	ID        uint `gorm:"primaryKey"`
	MediaID   uint `gorm:"index:"`
	SenderID  uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	CommentBase
	Sender User `gorm:"foreignKey:SenderID"`
}

type CommentView struct {
	CommentBase
	Sender UserView
}

type CommentRepository interface {
	Create(*Comment) error
	Update(*Comment) error
	Delete(id uint) error
	Load(id uint) (*CommentView, error)
	LoadAll(media_id uint) (*[]CommentView, error)
}

type CommentService interface {
	Add(*CommentBase) (*CommentView, error)
	Edit(*CommentBase) (*CommentView, error)
	Delete(id uint) error
	LoadAll(media_id uint) (*[]CommentView, error)
}
