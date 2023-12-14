package entity

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found")
	ErrMediaNotFound      = errors.New("Media not found")
	ErrCommentNotFound    = errors.New("Comment not found")
	ErrMediaTrackNotFound = errors.New("MediaTrack not found")

	ErrUserAlreadyExists       = errors.New("User already exists")
	ErrMediaAlreadyExists      = errors.New("Media already exists")
	ErrCommentAlreadyExists    = errors.New("Comment already exists")
	ErrMediaTrackAlreadyExists = errors.New("MediaTrack already exists")

	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
)
