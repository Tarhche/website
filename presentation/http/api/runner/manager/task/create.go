package task

import (
	"encoding/json"
	"net/http"

	createtask "github.com/khanzadimahdi/testproject/application/runner/manager/task/createTask"
)

const (
	DefaultMaxDiskSize   = 200 << 20 // 200 MB
	DefaultMaxMemorySize = 10 << 20  // 10 MB
	DefaultMaxCpu        = 0.05
)

type createHandler struct {
	useCase *createtask.UseCase
}

func NewCreateHandler(useCase *createtask.UseCase) *createHandler {
	return &createHandler{
		useCase: useCase,
	}
}

func (h *createHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request createtask.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.ResourceLimits.Disk == 0 {
		request.ResourceLimits.Disk = DefaultMaxDiskSize
	}

	if request.ResourceLimits.Memory == 0 {
		request.ResourceLimits.Memory = DefaultMaxMemorySize
	}

	if request.ResourceLimits.Cpu == 0 {
		request.ResourceLimits.Cpu = DefaultMaxCpu
	}

	request.OwnerUUID = "guest"

	response, err := h.useCase.Execute(&request)

	switch {
	case err != nil:
		rw.WriteHeader(http.StatusInternalServerError)
	case response != nil && len(response.ValidationErrors) > 0:
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response)
	default:
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(response)
	}
}
