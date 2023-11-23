package service

import (
	"backend/internal/entity"
	"backend/internal/repository"
)

type Service struct {
	User       entity.UserService
	Media      entity.MediaService
	Comment    entity.CommentService
	MediaTrack entity.MediaTrackService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:       NewUserService(repository.User),
		Media:      NewMediaService(repository.Media),
		Comment:    NewCommentService(repository.Comment),
		MediaTrack: NewMediaTrackService(repository.MediaTrack),
	}
}
