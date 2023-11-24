package entity

import (
	"backend/pkg/helpers"
	"time"
)

type Media struct {
	ID               uint `gorm:"primaryKey"`
	Title            string
	Description      string
	Image            string
	Category         string
	CumulativeRating uint64
	NumberOfRatings  uint32
	NumberOfTracks   uint32
	ReleaseYear      uint32
	Duration         time.Duration
	Genres           []Genre `gorm:"foreignKey:MediaID"`
}

type MediaView struct {
	Media
	Comments []CommentView
}

type MediaResult struct {
	ID     uint
	Title  string
	Image  string
	Rating float32
}

type Filter struct {
	Title        string
	Category     string
	Genres       []string
	YearFrom     uint32
	YearTo       uint32
	DurationFrom time.Duration
	DurationTo   time.Duration
	RatingFrom   float32
	RatingTo     float32
}

var Categories = helpers.MakeStringSet("Serials", "Movies", "Books", "Lectures")

type MediaRepository interface {
	Create(*Media) error
	Update(*Media) error
	Delete(id uint) error
	Get(id uint) (*Media, error)
	Load(id uint) (*MediaView, error)
	Search(*Filter) (*[]MediaResult, error)
}

type MediaService interface {
	Add(*Media) error
	Edit(*Media) error
	Delete(id uint) error
	Load(id uint) (*MediaView, error)
	Search(*Filter) (*[]MediaResult, error)
}
