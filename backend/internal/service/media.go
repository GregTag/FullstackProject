package service

import "backend/internal/entity"

type MediaService struct {
	mediaRepository entity.MediaRepository
}

func NewMediaService(mediaRepository entity.MediaRepository) *MediaService {
	return &MediaService{
		mediaRepository: mediaRepository,
	}
}

func (s *MediaService) Add(media *entity.Media) error {
	err := s.mediaRepository.Create(media)
	return err
}

func (s *MediaService) Edit(media *entity.Media) error {
	err := s.mediaRepository.Update(media)
	return err
}

func (s *MediaService) Delete(id uint) error {
	err := s.mediaRepository.Delete(id)
	return err
}

func (s *MediaService) Load(id uint) (*entity.MediaView, error) {
	media, err := s.mediaRepository.Load(id)
	return media, err
}

func (s *MediaService) Search(filter *entity.Filter) (*[]entity.MediaResult, error) {
	media, err := s.mediaRepository.Search(filter)
	return media, err
}
