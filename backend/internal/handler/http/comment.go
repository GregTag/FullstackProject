package handler

import (
	"backend/internal/entity"
	"errors"
	"log"
	"net/http"
	"strconv"

	"clevergo.tech/jsend"
)

func (h *Handler) commentAdd(w http.ResponseWriter, r *http.Request) {
	// TODO add authorization, take user_id from JWT
	var data map[string]string
	if parseBody(w, r, &data) != nil {
		return
	}

	media_id, err := strconv.ParseUint(data["media_id"], 10, 0)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("commentAdd: ", err.Error())
		return
	}

	var comment entity.CommentBase
	// comment.UserID =
	comment.MediaID = uint(media_id)
	comment.Content = data["content"]

	// view, err = h.commentService.Add(&comment)
	// if err != nil {
	// 	jsend.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Println("commentAdd: ", err.Error())
	// 	return

	// }

	// jsend.Success(w, view, http.StatusOK)
	jsend.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handler) commentEdit(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
	var data map[string]string
	if parseBody(w, r, &data) != nil {
		return
	}

	id, err := strconv.ParseUint(data["id"], 10, 0)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("commentAdd: ", err.Error())
		return
	}

	var comment entity.CommentBase

	comment.ID = uint(id)
	comment.Content = data["content"]

	view, err := h.commentService.Edit(&comment)
	switch {
	case errors.Is(err, entity.ErrCommentNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("commentEdit: ", err.Error())
		return
	}

	jsend.Success(w, view, http.StatusOK)
}

func (h *Handler) commentDelete(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
	id, exists := getIdParam(w, r)
	if !exists {
		return
	}

	err := h.commentService.Delete(id)
	switch {
	case errors.Is(err, entity.ErrCommentNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("commentDelete: ", err.Error())
		return
	}

	jsend.Success(w, nil, http.StatusOK)

}
