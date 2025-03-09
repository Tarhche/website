package stopTask

import (
	"encoding/json"

	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/runner/task/events"
)

const StopTaskName = "stopTask"

// StopTaskHandler handles the StopTask command
type StopTaskHandler struct {
	useCase *UseCase
}

var _ domain.MessageHandler = &StopTaskHandler{}

// NewStopTaskHandler creates a new StopTaskHandler
func NewStopTaskHandler(useCase *UseCase) *StopTaskHandler {
	return &StopTaskHandler{
		useCase: useCase,
	}
}

// Handle handles the StopTask command
func (h *StopTaskHandler) Handle(data []byte) error {
	var taskStoppageRequested events.TaskStoppageRequested
	if err := json.Unmarshal(data, &taskStoppageRequested); err != nil {
		return err
	}

	request := &Request{UUID: taskStoppageRequested.UUID}

	_, err := h.useCase.Execute(request)

	return err
}
