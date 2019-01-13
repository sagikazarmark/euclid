package euclidhttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewApp returns a new HTTP application.
func NewApp(sequence Sequence, errorHandler ErrorHandler) http.Handler {
	newID := NewIDHandler(sequence, errorHandler)

	router := mux.NewRouter()

	router.Path("/new/{name}").Methods("POST").Handler(newID)

	return router
}
