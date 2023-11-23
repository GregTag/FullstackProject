package service

import "backend/internal/entity"

type CommentService struct {
	commentRepository entity.CommentRepository
}

func NewCommentService(commentRepository entity.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (s *CommentService) Add(comment *entity.CommentBase) error {
	err := s.commentRepository.Create(comment)
	return err
}

func (s *CommentService) Edit(comment *entity.CommentBase) error {
	err := s.commentRepository.Update(comment)
	return err
}

func (s *CommentService) Delete(id uint) error {
	err := s.commentRepository.Delete(id)
	return err
}

func (s *CommentService) LoadAll(media_id uint) (*[]entity.CommentView, error) {
	comments, err := s.commentRepository.LoadAll(media_id)
	return comments, err
}
