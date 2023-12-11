package entity

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserRegister
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `gorm:"<-:create" json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type UserLogin struct {
	Login    string `gorm:"unique;<-:create" json:"login"`
	Password string `json:"password"`
}

type UserRegister struct {
	UserLogin
	Fullname string `json:"fullname"`
	Email    string `gorm:"unique" json:"email"`
}

type UserView struct {
	ID          uint      `json:"id"`
	Login       string    `json:"login"`
	Fullname    string    `json:"fullname"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type UserInfo struct {
	User   UserView         `json:"user"`
	Tracks []MediaTrackView `json:"tracks"`
}

type UserRepository interface {
	Create(*User) error
	Update(*User) error
	Delete(id uint) error
	Get(id uint) (*User, error)
	GetByLogin(login string) (*User, error)
	GetViewByLogin(login string) (*UserView, error)
	FillTracks(*UserView) (*UserInfo, error)
}

type UserService interface {
	Register(*UserRegister) (*UserInfo, error)
	Login(*UserLogin) (*UserInfo, error)
	Load(string) (*UserInfo, error)
	Change(*UserView) error
}
