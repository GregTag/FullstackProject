package handler

import (
	"backend/internal/entity"
	"backend/pkg/helpers"
	"errors"
	"log"
	"net/http"

	"clevergo.tech/jsend"
	"github.com/jinzhu/copier"
)

func (h *Handler) mediaLoad(w http.ResponseWriter, r *http.Request) {
	id, exists := getIdParam(w, r)
	if !exists {
		return
	}

	media, err := h.mediaService.Load(id)

	switch {
	case errors.Is(err, entity.ErrMediaNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("userLoad: ", err.Error())
		return
	}

	jsend.Success(w, media, http.StatusOK)
}

var mediaChangableFields = helpers.MakeStringSet("title", "description", "image", "category", "releaseYear", "duration", "genres")

func mediaCheckFields(data *map[string]interface{}, allow_id bool) bool {
	for key := range *data {
		if !mediaChangableFields.Contains(key) || (allow_id && key == "id") {
			return false
		}
	}
	return true
}

func parseMedia(w http.ResponseWriter, r *http.Request, media *entity.Media, allow_id bool) bool {
	var data map[string]interface{}
	if parseBody(w, r, &data) != nil {
		return false
	}
	if !mediaCheckFields(&data, allow_id) {
		jsend.Error(w, "Invalid fields", http.StatusBadRequest)
		return false
	}
	err := copier.Copy(media, &data)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("parseMedia: ", err.Error())
		return false
	}

	if !entity.Categories.Contains(media.Category) {
		jsend.Error(w, "Invalid category", http.StatusBadRequest)
		return false
	}
	return true
}

func (h *Handler) mediaAdd(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
	var media entity.Media
	if !parseMedia(w, r, &media, false) {
		return
	}

	err := h.mediaService.Add(&media)
	switch {
	case errors.Is(err, entity.ErrMediaAlreadyExists):
		jsend.Error(w, err.Error(), http.StatusConflict)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("mediaAdd: ", err.Error())
		return
	}

	jsend.Success(w, media, http.StatusOK)
}

func (h *Handler) mediaEdit(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
	var media entity.Media
	if !parseMedia(w, r, &media, true) {
		return
	}

	err := h.mediaService.Edit(&media)
	switch {
	case errors.Is(err, entity.ErrMediaNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("mediaEdit: ", err.Error())
		return
	}

	jsend.Success(w, media, http.StatusOK)
}

func (h *Handler) mediaDelete(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
	id, exists := getIdParam(w, r)
	if !exists {
		return
	}

	err := h.mediaService.Delete(id)
	switch {
	case errors.Is(err, entity.ErrMediaNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("mediaDelete: ", err.Error())
		return
	}

	jsend.Success(w, nil, http.StatusOK)
}

func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	var filter entity.Filter
	if parseBody(w, r, &filter) != nil {
		return
	}

	results, err := h.mediaService.Search(&filter)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("search: ", err.Error())
		return
	}

	jsend.Success(w, results, http.StatusOK)
}
