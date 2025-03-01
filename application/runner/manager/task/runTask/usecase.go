package runtask

import (
	"context"
	"encoding/json"

	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/runner/node"
	"github.com/khanzadimahdi/testproject/domain/runner/port"
	"github.com/khanzadimahdi/testproject/domain/runner/task"
	"github.com/khanzadimahdi/testproject/domain/runner/task/events"
	"github.com/khanzadimahdi/testproject/domain/translator"
)

const (
	nominatedNodesLimit = 10
)

type UseCase struct {
	taskRepository  task.Repository
	nodeRepository  node.Repository
	scheduler       task.Scheduler
	asyncCommandBus domain.PublishSubscriber
	translator      translator.Translator
}

func NewUseCase(
	taskRepository task.Repository,
	nodeRepository node.Repository,
	scheduler task.Scheduler,
	asyncCommandBus domain.PublishSubscriber,
	translator translator.Translator,
) *UseCase {
	return &UseCase{
		taskRepository:  taskRepository,
		nodeRepository:  nodeRepository,
		scheduler:       scheduler,
		asyncCommandBus: asyncCommandBus,
		translator:      translator,
	}
}

func (uc *UseCase) Execute(request *Request) (*Response, error) {
	t, err := uc.taskRepository.GetOne(request.UUID)
	if err != nil {
		return nil, err
	}

	destinationState := task.Scheduled
	if !task.ValidStateTransition(t.State, destinationState) {
		return &Response{
			ValidationErrors: domain.ValidationErrors{
				"task_id": uc.translator.Translate("invalid_state_transition"),
			},
		}, nil
	}

	nodes, err := uc.nodeRepository.GetAll(0, nominatedNodesLimit)
	if err != nil {
		return nil, err
	}
	selectedNode := uc.scheduler.Pick(&t, nodes)

	t.State = destinationState
	if _, err = uc.taskRepository.Save(&t); err != nil {
		return nil, err
	}

	event := events.TaskScheduled{
		UUID:          request.UUID,
		Name:          t.Name,
		Image:         t.Image,
		AutoRemove:    t.AutoRemove,
		PortBindings:  convertPortBindings(t.PortBindings),
		RestartPolicy: t.RestartPolicy,
		RestartCount:  t.RestartCount,
		HealthCheck:   t.HealthCheck,
		AttachStdin:   t.AttachStdin,
		AttachStdout:  t.AttachStdout,
		AttachStderr:  t.AttachStderr,
		Environment:   t.Environment,
		Command:       t.Command,
		Entrypoint:    t.Entrypoint,
		Mounts:        convertMounts(t.Mounts),
		ResourceLimits: events.ResourceLimits{
			Cpu:    t.ResourceLimits.Cpu,
			Memory: t.ResourceLimits.Memory,
			Disk:   t.ResourceLimits.Disk,
		},
		NominatedNode: selectedNode.Name,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	if err = uc.asyncCommandBus.Publish(context.Background(), events.TaskScheduledName, payload); err != nil {
		return nil, err
	}

	return nil, nil
}

func convertPortBindings(domainPorts []port.PortMap) []events.PortMap {
	result := make([]events.PortMap, len(domainPorts))
	for i, p := range domainPorts {
		portMap := make(events.PortMap)
		for portNum, bindings := range p {
			portBindings := make([]events.PortBinding, len(bindings))
			for j, b := range bindings {
				portBindings[j] = events.PortBinding{
					HostIP:   b.HostIP,
					HostPort: b.HostPort,
				}
			}
			portMap[portNum] = portBindings
		}
		result[i] = portMap
	}
	return result
}

func convertMounts(domainMounts []task.Mount) []events.Mount {
	result := make([]events.Mount, len(domainMounts))
	for i, m := range domainMounts {
		result[i] = events.Mount{
			Source:   m.Source,
			Target:   m.Target,
			Type:     m.Type,
			ReadOnly: m.ReadOnly,
		}
	}
	return result
}
