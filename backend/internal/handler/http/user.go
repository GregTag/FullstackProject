package handler

import (
	"backend/internal/entity"
	"errors"
	"log"
	"net/http"

	"clevergo.tech/jsend"
)

func (h *Handler) userRegister(w http.ResponseWriter, r *http.Request) {
	var user_reg entity.UserRegister
	if parseBody(w, r, &user_reg) != nil {
		return
	}

	info, err := h.userService.Register(&user_reg)
	switch {
	case errors.Is(err, entity.ErrUserAlreadyExists):
		jsend.Error(w, err.Error(), http.StatusConflict)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("userRegister: ", err.Error())
		return
	}
	h.provideAuthHeader(w, info.User.ID)
	jsend.Success(w, info, http.StatusOK)
}

func (h *Handler) userLogin(w http.ResponseWriter, r *http.Request) {
	var user_login entity.UserLogin
	if parseBody(w, r, &user_login) != nil {
		return
	}

	info, err := h.userService.Login(&user_login)
	switch {
	case errors.Is(err, entity.ErrUserNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("userLogin: ", err.Error())
		return
	}
	h.provideAuthHeader(w, info.User.ID)
	jsend.Success(w, info, http.StatusOK)
}

func (h *Handler) provideAuthHeader(w http.ResponseWriter, user_id uint) bool {
	token, err := h.auth.MakeToken(user_id)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("auth header: ", err.Error())
		return false
	}
	w.Header().Add(AUTH_HEADER_NAME, "Bearer "+token)
	return true
}

func (h *Handler) userLogout(w http.ResponseWriter, r *http.Request) {
	_, verified := h.checkAuth(w, r)
	if !verified {
		return
	}
	w.Header().Del(AUTH_HEADER_NAME)
	jsend.Success(w, nil, http.StatusOK)
}

func (h *Handler) userEdit(w http.ResponseWriter, r *http.Request) {
	user_id, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	var user_edit map[string]string
	if parseBody(w, r, &user_edit) != nil {
		return
	}

	view := entity.UserView{
		ID:       user_id,
		Fullname: user_edit["fullname"],
		Avatar:   user_edit["avatar"],
	}

	err := h.userService.Change(&view)
	switch {
	case errors.Is(err, entity.ErrUserNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("userEdit: ", err.Error())
		return
	}
	jsend.Success(w, view, http.StatusOK)
}
func (h *Handler) userLoad(w http.ResponseWriter, r *http.Request) {
	_, verified := h.checkAuth(w, r)
	if !verified {
		return
	}

	login, exists := getParam(w, r, "login")
	if !exists {
		return
	}

	info, err := h.userService.Load(login)
	switch {
	case errors.Is(err, entity.ErrUserNotFound):
		jsend.Error(w, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		jsend.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("userLoad: ", err.Error())
		return
	}
	jsend.Success(w, info, http.StatusOK)
}
