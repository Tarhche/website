package stoptask

import (
	"github.com/khanzadimahdi/testproject/domain/runner/task"
)

const StopTaskName = "stopTask"

// StopTask command
type StopTask struct {
	UUID string `json:"uuid"`
}

// StopTaskHandler handles the StopTask command
type StopTaskHandler struct {
	taskRepository task.Repository
}

// NewStopTaskHandler creates a new StopTaskHandler
func NewStopTaskHandler(taskRepository task.Repository) *StopTaskHandler {
	return &StopTaskHandler{
		taskRepository: taskRepository,
	}
}

// Handle handles the StopTask command
func (h *StopTaskHandler) Handle(data []byte) error {
	// Implementation for stopping a task
	return nil
}
