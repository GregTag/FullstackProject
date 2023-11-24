package entity

import "errors"

var ErrUserNotFound = errors.New("User not found")
var ErrMediaNotFound = errors.New("Media not found")
var ErrCommentNotFound = errors.New("Comment not found")
var ErrMediaTrackNotFound = errors.New("MediaTrack not found")

var ErrUserAlreadyExists = errors.New("User already exists")
var ErrMediaAlreadyExists = errors.New("Media already exists")
var ErrCommentAlreadyExists = errors.New("Comment already exists")
var ErrMediaTrackAlreadyExists = errors.New("MediaTrack already exists")
