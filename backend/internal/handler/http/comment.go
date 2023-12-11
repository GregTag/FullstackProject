package handler

import (
	"backend/internal/entity"
	"errors"
	"log"
	"net/http"

	"clevergo.tech/jsend"
)

func (h *Handler) commentAdd(w http.ResponseWriter, r *http.Request) {
	user_id, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	var data map[string]interface{}
	if parseBody(w, r, &data) != nil {
		return
	}
	media_id, ok1 := data["media_id"].(float64)
	content, ok2 := data["content"].(string)
	if !ok1 || !ok2 {
		jsend.Error(w, "invalid comment", http.StatusBadRequest)
		return
	}

	comment := entity.CommentBase{
		SenderID: user_id,
		MediaID:  uint(media_id),
		Content:  content,
	}

	view, err := h.commentService.Add(&comment)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("commentAdd: ", err.Error())
		return

	}

	jsend.Success(w, view, http.StatusOK)
}

func (h *Handler) commentEdit(w http.ResponseWriter, r *http.Request) {
	user_id, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	var data map[string]interface{}
	if parseBody(w, r, &data) != nil {
		return
	}

	id, ok1 := data["id"].(float64)
	content, ok2 := data["content"].(string)
	if !ok1 || !ok2 {
		jsend.Error(w, "invalid comment", http.StatusBadRequest)
		return
	}

	comment := entity.CommentBase{
		ID:       uint(id),
		SenderID: user_id,
		Content:  content,
	}

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
	user_id, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	id, exists := getIdParam(w, r)
	if !exists {
		return
	}

	err := h.commentService.Delete(id, user_id)
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
