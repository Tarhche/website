package createtask

import (
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/runner/task"
)

// UseCase creates a task
type UseCase struct {
	taskRepository task.Repository
	validator      domain.Validator
}

// NewUseCase creates an instance of the UseCase
func NewUseCase(
	taskRepository task.Repository,
	validator domain.Validator,
) *UseCase {
	return &UseCase{
		taskRepository: taskRepository,
		validator:      validator,
	}
}

// Execute executes the use case
func (uc *UseCase) Execute(request *Request) (*Response, error) {
	if validationErrors := uc.validator.Validate(request); len(validationErrors) > 0 {
		return &Response{
			ValidationErrors: validationErrors,
		}, nil
	}

	t := task.Task{
		Name:          request.Name,
		State:         task.Created,
		Image:         request.Image,
		PortBindings:  request.ConvertPortBindings(),
		RestartPolicy: request.RestartPolicy,
		RestartCount:  request.RestartCount,
		HealthCheck:   request.HealthCheck,
		AttachStdin:   request.AttachStdin,
		AttachStdout:  request.AttachStdout,
		AttachStderr:  request.AttachStderr,
		Environment:   request.Environment,
		Command:       request.Command,
		Entrypoint:    request.Entrypoint,
		Mounts:        request.ConvertMounts(),
		ResourceLimits: task.ResourceLimits{
			Cpu:    request.ResourceLimits.Cpu,
			Memory: request.ResourceLimits.Memory,
			Disk:   request.ResourceLimits.Disk,
		},
		OwnerUUID: request.OwnerUUID,
	}

	_, err := uc.taskRepository.Save(&t)

	return &Response{UUID: t.UUID}, err
}
