package task

import (
	"encoding/json"
	"errors"
	"net/http"

	runtask "github.com/khanzadimahdi/testproject/application/runner/manager/task/runTask"
	"github.com/khanzadimahdi/testproject/domain"
)

type runHandler struct {
	useCase *runtask.UseCase
}

func NewRunHandler(useCase *runtask.UseCase) *runHandler {
	return &runHandler{
		useCase: useCase,
	}
}

func (h *runHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	UUID := r.PathValue("uuid")
	request := &runtask.Request{
		UUID: UUID,
	}

	response, err := h.useCase.Execute(request)
	switch {
	case errors.Is(err, domain.ErrNotExists):
		rw.WriteHeader(http.StatusNotFound)
	case err != nil:
		rw.WriteHeader(http.StatusInternalServerError)
	case response != nil && len(response.ValidationErrors) > 0:
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response)
	default:
		rw.WriteHeader(http.StatusAccepted)
	}
}
