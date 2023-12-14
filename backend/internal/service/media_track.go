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
	mediaTrack := entity.MediaTrack{MediaTrackBase: *track}
	err := s.mediaTrackRepository.Create(&mediaTrack)
	if err != nil {
		return err
	}
	err = copier.Copy(track, &mediaTrack)
	return err
}

func (s *MediaTrackService) Change(track *entity.MediaTrackBase) error {
	mediaTrack := entity.MediaTrack{MediaTrackBase: *track}
	err := s.mediaTrackRepository.Update(&mediaTrack)
	if err != nil {
		return err
	}
	err = copier.Copy(track, &mediaTrack)
	return err
}

func (s *MediaTrackService) Untrack(userID uint, mediaID uint) error {
	err := s.mediaTrackRepository.Delete(userID, mediaID)
	return err
}

func (s *MediaTrackService) LoadAll(userID uint) (*[]entity.MediaTrackView, error) {
	tracks, err := s.mediaTrackRepository.LoadAll(userID)
	return tracks, err
}
