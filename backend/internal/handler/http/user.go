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
	jsend.Success(w, info, http.StatusOK)
}

func (h *Handler) userLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: add JWT
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
	jsend.Success(w, info, http.StatusOK)
}

func (h *Handler) userLogout(w http.ResponseWriter, r *http.Request) {
	// TODO: JWT expiration
	jsend.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *Handler) userEdit(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization, take user_id from JWT
	var user_edit map[string]string
	if parseBody(w, r, &user_edit) != nil {
		return
	}
	jsend.Error(w, "Not implemented", http.StatusNotImplemented)

}
func (h *Handler) userLoad(w http.ResponseWriter, r *http.Request) {
	// TODO: add authorization
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
