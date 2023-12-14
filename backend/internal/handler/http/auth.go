package handler

import (
	"backend/internal/entity"
	"log"
	"net/http"
	"strconv"
	"strings"

	"clevergo.tech/jsend"
)

var AUTH_HEADER_NAME = "Authorization"

func (h *Handler) getClaimsFromAuthHeader(r *http.Request) (*map[string]string, error) {
	jwtClaims := &map[string]string{}

	authzHeader := r.Header.Get(AUTH_HEADER_NAME)
	if authzHeader == "" {
		return jwtClaims, entity.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(authzHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return jwtClaims, entity.ErrInvalidAuthHeader
	}

	jwtClaims, err := h.auth.FetchToken(headerParts[1])
	return jwtClaims, err
}

func (h *Handler) checkAuth(w http.ResponseWriter, r *http.Request) (uint, bool) {
	claims, err := h.getClaimsFromAuthHeader(r)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusUnauthorized)
		return 0, false
	}
	userID, err := strconv.ParseUint((*claims)["sub"], 10, 32)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusUnauthorized)
		log.Printf("checkAuth: %s\n", err.Error())
		return 0, false
	}
	return uint(userID), true
}
