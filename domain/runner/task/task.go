package task

import (
	"time"

	"github.com/khanzadimahdi/testproject/domain/runner/port"
)

// Task represents a task specification
type Task struct {
	UUID           string
	Name           string
	State          State
	Image          string
	PortBindings   []port.PortMap
	RestartPolicy  string
	RestartCount   uint
	HealthCheck    string
	AttachStdin    bool
	AttachStdout   bool
	AttachStderr   bool
	Environment    []string
	Command        []string
	Entrypoint     []string
	Mounts         []Mount
	ResourceLimits ResourceLimits
	ContainerID    string
	OwnerUUID      string
	CreatedAt      time.Time
	StartedAt      time.Time
	FinishedAt     time.Time
}

// Mount represents a mount point of volume
type Mount struct {
	Source   string
	Target   string
	Type     string
	ReadOnly bool
}

// ResourceLimits represents the resource limits of the container
type ResourceLimits struct {
	Cpu    float64
	Memory uint64
	Disk   uint64
}

// Manager represents a manager of tasks
type Manager interface {
	Run(task *Task) (containerID string, err error)
	Stop(taskUUID string) error
	Inspect(taskUUID string) error
}

// Repository represents a repository of tasks
type Repository interface {
	GetAll(offset uint, limit uint) ([]Task, error)
	GetOne(UUID string) (Task, error)
	Save(*Task) (uuid string, err error)
	Delete(UUID string) error
	Count() (uint, error)
}
