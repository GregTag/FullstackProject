package entity

type MediaTrack struct {
	UserID      uint `gorm:"primaryKey"`
	MediaID     uint `gorm:"primaryKey"`
	Rating      uint8
	TrackStatus string
}

type MediaTrackView struct {
	MediaTrack
	MediaTitle string
}

type MediaTrackRepository interface {
	Create(*MediaTrack) error
	Update(*MediaTrack) error
	Delete(user_id uint, media_id uint) error
	Get(user_id uint, media_id uint) (*MediaTrack, error)
	LoadAll(user_id uint) (*[]MediaTrackView, error)
}
