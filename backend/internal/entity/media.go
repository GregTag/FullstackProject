package entity

import (
	"backend/pkg/helpers"
	"time"
)

type MediaBase struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	Category    string        `json:"category"`
	ReleaseYear uint32        `json:"release_year"`
	Duration    time.Duration `json:"duration"`
	Genres      []Genre       `gorm:"foreignKey:MediaID" json:"genres"`
}

type Media struct {
	MediaBase
	CumulativeRating int32  `json:"cumulative_rating"`
	NumberOfRatings  uint32 `json:"number_of_ratings"`
	NumberOfTracks   uint32 `json:"number_of_tracks"`
}

type MediaView struct {
	Media
	Comments []CommentView `json:"comments"`
}

type MediaResult struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Image  string  `json:"image"`
	Rating float32 `json:"rating"`
}

type Filter struct {
	Title        string        `json:"title"`
	Category     string        `json:"category"`
	Genres       []string      `json:"genres"`
	YearFrom     uint32        `json:"year_from"`
	YearTo       uint32        `json:"year_to"`
	DurationFrom time.Duration `json:"duration_from"`
	DurationTo   time.Duration `json:"duration_to"`
	RatingFrom   float32       `json:"rating_from"`
	RatingTo     float32       `json:"rating_to"`
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
