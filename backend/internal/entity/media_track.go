package entity

import "backend/pkg/helpers"

type MediaTrackBase struct {
	UserID      uint   `gorm:"primaryKey" json:"user_id"`
	MediaID     uint   `gorm:"primaryKey" json:"media_id"`
	Rating      uint8  `json:"rating"`
	TrackStatus string `json:"track_status"`
}

const MaxRating = 10

var TrackStatuses = helpers.MakeStringSet("planned", "in_progress", "completed", "dropped")

type MediaTrack struct {
	MediaTrackBase
	Media Media `gorm:"foreignKey:MediaID"`
}

type MediaTrackView struct {
	MediaTrackBase
	MediaTitle    string `json:"media_title"`
	MediaCategory string `json:"media_category"`
}

type MediaTrackRepository interface {
	Create(*MediaTrack) error
	Update(*MediaTrack) error
	Delete(userID, mediaID uint) error
	Get(userID, mediaID uint) (*MediaTrack, error)
	LoadAll(userID uint) (*[]MediaTrackView, error)
}

type MediaTrackService interface {
	Track(*MediaTrackBase) error
	Change(*MediaTrackBase) error
	Untrack(userID uint, mediaID uint) error
	LoadAll(userID uint) (*[]MediaTrackView, error)
}
