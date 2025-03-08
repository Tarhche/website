package gettasks

import (
	"github.com/khanzadimahdi/testproject/domain/runner/container"
	"github.com/khanzadimahdi/testproject/domain/runner/task"
)

const limit = 10

type UseCase struct {
	containerManager container.Manager
}

func NewUseCase(containerManager container.Manager) *UseCase {
	return &UseCase{
		containerManager: containerManager,
	}
}

func (uc *UseCase) Execute() (*Response, error) {
	allContainers, err := uc.containerManager.GetAll()
	if err != nil {
		return nil, err
	}

	tasks := make([]task.Task, len(allContainers))
	for i, c := range allContainers {
		tasks[i] = task.Task{
			UUID:        c.Labels[container.TaskUUIDLabelKey],
			Name:        c.Labels[container.TaskNameLabelKey],
			Image:       c.Image,
			ContainerID: c.ID,
			CreatedAt:   c.CreatedAt,
		}
	}

	return NewResponse(tasks), nil
}
