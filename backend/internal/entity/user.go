package entity

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	UserRegister
	Avatar      string
	CreatedAt   time.Time
	LastLoginAt time.Time
}

type UserLogin struct {
	Login    string
	Password string
}

type UserRegister struct {
	UserLogin
	Fullname string
	Email    string
}

type UserRepository interface {
	Create(*User) error
	Update(*User) error
	Delete(id uint) error
	Get(id uint) (*User, error)
	GetByLogin(login string) (*User, error)
}
