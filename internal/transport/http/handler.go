package http

import (
	"encoding/json"
	"net/http"

	"github.com/cmcki85/RESTurant-Full-Stack/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores the pointer to the comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from API
type Response struct {
	Message string
	Error   string
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// LoggingMiddleware - adds middleware for our application.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).Info("Request Handled.")
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes - Sets up the routes for the application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := sendOKResponse(w, Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

func sendOKResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
