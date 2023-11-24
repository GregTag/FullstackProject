package service

import (
	"backend/internal/entity"

	"github.com/jinzhu/copier"
)

type CommentService struct {
	commentRepository entity.CommentRepository
}

func NewCommentService(commentRepository entity.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (s *CommentService) Add(base *entity.CommentBase) (*entity.CommentView, error) {
	comment := entity.Comment{CommentBase: *base}
	err := s.commentRepository.Create(&comment)
	if err != nil {
		return nil, err
	}
	view := new(entity.CommentView)
	copier.Copy(view, &comment)
	return view, nil
}

func (s *CommentService) Edit(base *entity.CommentBase) (*entity.CommentView, error) {
	comment := entity.Comment{CommentBase: *base}
	err := s.commentRepository.Update(&comment)
	if err != nil {
		return nil, err
	}
	view := new(entity.CommentView)
	copier.Copy(view, &comment)
	return view, nil
}

func (s *CommentService) Delete(id uint) error {
	err := s.commentRepository.Delete(id)
	return err
}

func (s *CommentService) LoadAll(media_id uint) (*[]entity.CommentView, error) {
	comments, err := s.commentRepository.LoadAll(media_id)
	return comments, err
}
