package entity

import "backend/pkg/helpers"

type MediaTrackBase struct {
	UserID      uint `gorm:"primaryKey"`
	MediaID     uint `gorm:"primaryKey"`
	Rating      uint8
	TrackStatus string
}

const MaxRating = 10

var TrackStatuses = helpers.MakeStringSet("planned", "in_progress", "completed", "dropped")

type MediaTrack struct {
	MediaTrackBase
	Media Media `gorm:"foreignKey:MediaID"`
}

type MediaTrackView struct {
	MediaTrackBase
	MediaTitle string
}

type MediaTrackRepository interface {
	Create(*MediaTrack) error
	Update(*MediaTrack) error
	Delete(user_id uint, media_id uint) error
	Get(user_id uint, media_id uint) (*MediaTrack, error)
	LoadAll(user_id uint) (*[]MediaTrackView, error)
}

type MediaTrackService interface {
	Track(*MediaTrackBase) error
	Change(*MediaTrackBase) error
	Untrack(user_id uint, media_id uint) error
	LoadAll(user_id uint) (*[]MediaTrackView, error)
}
