package service

import (
	"backend/internal/entity"

	"github.com/jinzhu/copier"
)

type MediaTrackService struct {
	mediaTrackRepository entity.MediaTrackRepository
}

func NewMediaTrackService(mediaTrackRepository entity.MediaTrackRepository) *MediaTrackService {
	return &MediaTrackService{
		mediaTrackRepository: mediaTrackRepository,
	}
}

func (s *MediaTrackService) Track(track *entity.MediaTrackBase) error {
	media_track := entity.MediaTrack{MediaTrackBase: *track}
	err := s.mediaTrackRepository.Create(&media_track)
	if err != nil {
		return err
	}
	err = copier.Copy(track, &media_track)
	return err
}

func (s *MediaTrackService) Change(track *entity.MediaTrackBase) error {
	media_track := entity.MediaTrack{MediaTrackBase: *track}
	err := s.mediaTrackRepository.Update(&media_track)
	if err != nil {
		return err
	}
	err = copier.Copy(track, &media_track)
	return err
}

func (s *MediaTrackService) Untrack(user_id uint, media_id uint) error {
	err := s.mediaTrackRepository.Delete(user_id, media_id)
	return err
}

func (s *MediaTrackService) LoadAll(user_id uint) (*[]entity.MediaTrackView, error) {
	tracks, err := s.mediaTrackRepository.LoadAll(user_id)
	return tracks, err
}
