package runtask

import (
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/runner/task"
)

const RunTaskName = "runTask"

// RunTask command
type RunTask struct {
	UUID string `json:"uuid"`
}

// RunTaskHandler handles RunTask command
type RunTaskHandler struct {
	taskRepository task.Repository
}

var _ domain.MessageHandler = &RunTaskHandler{}

func NewRunTaskHandler(taskRepository task.Repository) *RunTaskHandler {
	return &RunTaskHandler{taskRepository: taskRepository}
}

func (h *RunTaskHandler) Handle(data []byte) error {
	// get command
	// schedule task on one of the workers

	return nil
}
