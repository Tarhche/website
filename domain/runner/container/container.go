package container

import (
	"time"

	"github.com/khanzadimahdi/testproject/domain/runner/port"
	"github.com/khanzadimahdi/testproject/domain/runner/stats"
)

// Container represents a container specification
type Container struct {
	ID               string
	Name             string
	Status           string
	Image            string
	ResourceLimits   ResourceLimits
	RestartPolicy    string
	RestartCount     uint
	WorkingDirectory string
	ExposedPorts     port.PortSet
	PortBindings     port.PortMap
	HealthCheck      string
	Environment      []string
	Entrypoint       []string
	Command          []string
	CreatedAt        time.Time
}

// ResourceLimits represents the resource limits of the container
type ResourceLimits struct {
	Cpu    float64
	Memory uint64
	Disk   uint64
}

// Manager represents a manager of containers
type Manager interface {
	Create(container *Container) (containerUUID string, err error)
	Start(containerUUID string) error
	Stop(containerUUID string) error
	Delete(containerUUID string) error
	Inspect(containerUUID string) (Container, error)
	Stats(containerUUID string) (stats.Stats, error)
}
