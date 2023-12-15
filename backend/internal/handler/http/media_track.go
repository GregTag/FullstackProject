package handler

import (
	"backend/internal/entity"
	"errors"
	"log"
	"net/http"

	"clevergo.tech/jsend"
)

func validateTrack(w http.ResponseWriter, track *entity.MediaTrackBase) bool {
	if track.Rating > entity.MaxRating {
		jsend.Error(w, "Rating must be in range [0, 10]", http.StatusBadRequest)
		return false
	}

	if track.TrackStatus != "" && !entity.TrackStatuses.Contains(track.TrackStatus) {
		jsend.Error(w, "Invalid track status", http.StatusBadRequest)
		return false
	}

	return true
}

func (h *Handler) trackAdd(w http.ResponseWriter, r *http.Request) {
	userID, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	var track entity.MediaTrackBase
	if parseBody(w, r, &track) != nil {
		return
	}
	if !validateTrack(w, &track) {
		return
	}

	track.UserID = userID

	err := h.mediaTrackService.Track(&track)
	switch {
	case errors.Is(err, entity.ErrMediaTrackAlreadyExists):
		jsend.Error(w, err.Error(), http.StatusConflict)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("trackAdd: ", err.Error())
		return
	}

	jsend.Success(w, track, http.StatusOK)
}

func (h *Handler) trackEdit(w http.ResponseWriter, r *http.Request) {
	userID, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	var track entity.MediaTrackBase
	if parseBody(w, r, &track) != nil {
		return
	}
	if !validateTrack(w, &track) {
		return
	}

	track.UserID = userID

	err := h.mediaTrackService.Change(&track)
	switch {
	case errors.Is(err, entity.ErrMediaTrackNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("trackEdit: ", err.Error())
		return
	}

	jsend.Success(w, track, http.StatusOK)
}

func (h *Handler) trackDelete(w http.ResponseWriter, r *http.Request) {
	userID, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	mediaID, exists := getIdParam(w, r)
	if !exists {
		return
	}

	err := h.mediaTrackService.Untrack(userID, mediaID)
	switch {
	case errors.Is(err, entity.ErrMediaTrackNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("trackDelete: ", err.Error())
		return
	}

	jsend.Success(w, nil, http.StatusOK)
}
