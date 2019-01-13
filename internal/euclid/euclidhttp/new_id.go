package euclidhttp

import (
	"encoding/json"
	"net/http"

	"github.com/moogar0880/problems"
	"github.com/pkg/errors"
	"github.com/sagikazarmark/euclid/.gen/openapi/go"
)

// Sequence of IDs.
type Sequence interface {
	// Next returns the next ID in the sequence.
	Next(name string) (uint, error)
}

// NewIDHandler returns an HTTP handler that returns a new ID.
func NewIDHandler(sequence Sequence, errorHandler ErrorHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var apiRequest api.IdRequest

		err := json.NewDecoder(r.Body).Decode(&apiRequest)
		if err != nil {
			errorHandler.Handle(errors.Wrap(err, "failed to decode request"))

			problem := problems.NewStatusProblem(http.StatusBadRequest)
			w.Header().Set("Content-Type", problems.ProblemMediaType)
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(problem)

			return
		}

		next, err := sequence.Next(apiRequest.Name)
		if err != nil {
			errorHandler.Handle(err)

			problem := problems.NewDetailedProblem(http.StatusInternalServerError, "failed to generate next ID")
			w.Header().Set("Content-Type", problems.ProblemMediaType)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(problem)

			return
		}

		apiResponse := api.IdResponse{
			Id: int32(next),
		}

		body, err := json.Marshal(apiResponse)
		if err != nil {
			errorHandler.Handle(err)

			problem := problems.NewDetailedProblem(http.StatusInternalServerError, "failed to marshal ID")
			w.Header().Set("Content-Type", problems.ProblemMediaType)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(problem)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(body)
	})
}
