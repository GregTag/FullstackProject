package entity

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	UserRegister
	Avatar      string
	CreatedAt   time.Time `gorm:"<-:false"`
	LastLoginAt time.Time
}

type UserLogin struct {
	Login    string `gorm:"unique;<-:create"`
	Password string
}

type UserRegister struct {
	UserLogin
	Fullname string
	Email    string `gorm:"unique"`
}

type UserView struct {
	ID          uint
	Login       string
	Fullname    string
	Avatar      string
	CreatedAt   time.Time
	LastLoginAt time.Time
}

type UserInfo struct {
	User   UserView
	Tracks []MediaTrackView
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
