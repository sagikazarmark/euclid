package euclidhttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moogar0880/problems"
)

// Sequence of IDs.
type Sequence interface {
	// Next returns the next ID in the sequence.
	Next(name string) (uint, error)
}

// IDResponse is the returned response.
type IDResponse struct {
	ID uint `json:"id"`
}

// NewIDHandler returns an HTTP handler that returns a new ID.
func NewIDHandler(sequence Sequence, errorHandler ErrorHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)
		
		next, err := sequence.Next(name["name"])
		if err != nil {
			errorHandler.Handle(err)

			problem := problems.NewDetailedProblem(http.StatusInternalServerError, "failed to generate next ID")
			w.Header().Set("Content-Type", problems.ProblemMediaType)
			_ = json.NewEncoder(w).Encode(problem)
		}

		response := IDResponse{next}

		body, err := json.Marshal(response)
		if err != nil {
			errorHandler.Handle(err)

			problem := problems.NewDetailedProblem(http.StatusInternalServerError, "failed to marshal ID")
			w.Header().Set("Content-Type", problems.ProblemMediaType)
			_ = json.NewEncoder(w).Encode(problem)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(body)
	})
}
