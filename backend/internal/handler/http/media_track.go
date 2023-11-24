package handler

import (
	"net/http"

	"clevergo.tech/jsend"
)

// func parseTrack(w http.ResponseWriter, data *map[string]string, track *entity.MediaTrackBase) bool {
// 	media_id, err1 := strconv.ParseUint((*data)["mediaID"], 10, 0)
// 	rating, err2 := strconv.ParseUint((*data)["rating"], 10, 0)
// 	if err1 != nil || err2 != nil {
// 		if err1 == nil {
// 			err1 = err2
// 		}
// 		jsend.Error(w, err1.Error(), http.StatusBadRequest)
// 		log.Println("trackEdit: ", err1.Error())
// 		return false
// 	}

// 	if rating > entity.MaxRating {
// 		jsend.Error(w, "Rating must be in range [0, 10]", http.StatusBadRequest)
// 		return false
// 	}

// 	track.MediaID = uint(media_id)
// 	track.Rating = uint8(rating)
// 	track.TrackStatus = (*data)["trackStatus"]

// 	if !entity.TrackStatuses.Contains(track.TrackStatus) {
// 		jsend.Error(w, "Invalid track status", http.StatusBadRequest)
// 		return false
// 	}

// 	return true
// }

func (h *Handler) trackAdd(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization, take user_id from JWT

	// var data map[string]string
	// if parseBody(w, r, &data) != nil {
	// 	return
	// }

	// var track entity.MediaTrackBase
	// if !parseTrack(w, &data, &track) {
	// 	return
	// }

	// track.UserID = user_id

	// err := h.mediaTrackService.Track(&track)
	// switch {
	// case errors.Is(err, entity.ErrMediaTrackAlreadyExists):
	// 	jsend.Error(w, err.Error(), http.StatusConflict)
	// 	return
	// case err != nil:
	// 	jsend.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Println("trackAdd: ", err.Error())
	// 	return
	// }

	// jsend.Success(w, track, http.StatusOK)

	jsend.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handler) trackEdit(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization, take user_id from JWT

	// var data map[string]string
	// if parseBody(w, r, &data) != nil {
	// 	return
	// }

	// var track entity.MediaTrackBase
	// if !parseTrack(w, &data, &track) {
	// 	return
	// }

	// track.UserID = user_id

	// err := h.mediaTrackService.Change(&track)
	// switch {
	// case errors.Is(err, entity.ErrMediaTrackNotFound):
	// 	jsend.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// case err != nil:
	// 	jsend.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Println("trackEdit: ", err.Error())
	// 	return
	// }

	// jsend.Success(w, track, http.StatusOK)

	jsend.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handler) trackDelete(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization, take user_id from JWT

	// media_id, exists := getIdParam(w, r)
	// if !exists {
	// 	return
	// }

	// err := h.mediaTrackService.Untrack(user_id, media_id)
	// switch {
	// case errors.Is(err, entity.ErrMediaTrackNotFound):
	// 	jsend.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// case err != nil:
	// 	jsend.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Println("trackDelete: ", err.Error())
	// 	return
	// }

	// jsend.Success(w, nil, http.StatusOK)

	jsend.Error(w, "Not implemented", http.StatusNotImplemented)
}
