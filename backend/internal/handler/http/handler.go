package handler

import (
	"backend/internal/entity"
	"backend/internal/service"
	"backend/pkg/auth"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"clevergo.tech/jsend"
	"github.com/gorilla/mux"
)

type Handler struct {
	auth              entity.AuthManager
	userService       entity.UserService
	mediaService      entity.MediaService
	commentService    entity.CommentService
	mediaTrackService entity.MediaTrackService
}

func NewHandler(s *service.Service, auth *auth.AuthManager) *Handler {
	return &Handler{
		auth:              auth,
		userService:       s.User,
		mediaService:      s.Media,
		commentService:    s.Comment,
		mediaTrackService: s.MediaTrack,
	}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			if r.Method == http.MethodOptions {
				return
			}
			next.ServeHTTP(w, r)
		})
	})
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Method: %s, URL: %s, RemoteAddr: %s\n", r.Method, r.URL, r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/user/register", h.userRegister).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/login", h.userLogin).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/logout", h.userLogout).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/edit", h.userEdit).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/user/load/{login}", h.userLoad).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/media/load/{id}", h.mediaLoad).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/media/add", h.mediaAdd).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/media/edit", h.mediaEdit).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/media/delete/{id}", h.mediaDelete).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/search", h.search).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/comment/add", h.commentAdd).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/comment/edit", h.commentEdit).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/comment/delete/{id}", h.commentDelete).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/track/add", h.trackAdd).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/track/edit", h.trackEdit).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/track/delete/{id}", h.trackDelete).Methods(http.MethodDelete, http.MethodOptions)

	return r
}

func parseBody(w http.ResponseWriter, r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		jsend.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

func getParam(w http.ResponseWriter, r *http.Request, param string) (string, bool) {
	vars := mux.Vars(r)
	value, exists := vars[param]
	if !exists {
		jsend.Error(w, param+" not specified", http.StatusBadRequest)
	}
	return value, exists
}

func getIdParam(w http.ResponseWriter, r *http.Request) (uint, bool) {
	value, exists := getParam(w, r, "id")
	if !exists {
		return 0, false
	}
	id, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		jsend.Error(w, "ID must be integer value", http.StatusBadRequest)
		log.Println("getIntParam: ", err.Error())
		return 0, false
	}
	return uint(id), true

}
