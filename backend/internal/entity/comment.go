package entity

import "time"

type CommentBase struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SenderID  uint      `gorm:"primaryKey;autoIncrement:false" json:"sender_id"`
	MediaID   uint      `gorm:"index:" json:"media_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	CommentBase
	Sender User `gorm:"foreignKey:SenderID"`
}

type CommentView struct {
	CommentBase
	Sender UserView `json:"sender"`
}

type CommentRepository interface {
	Create(*Comment) error
	Update(*Comment) error
	Delete(id, senderID uint) error
	Load(id, senderID uint) (*CommentView, error)
	LoadAll(mediaID uint) (*[]CommentView, error)
}

type CommentService interface {
	Add(*CommentBase) (*CommentView, error)
	Edit(*CommentBase) (*CommentView, error)
	Delete(id, senderID uint) error
	LoadAll(mediaID uint) (*[]CommentView, error)
}
