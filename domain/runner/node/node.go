package node

// Node represents a node in the cluster
type Node struct {
	Name      string
	Role      Role
	API       string
	Resources Resource
}

// Resource represents the hardware resources of the node
type Resource struct {
	Cpu    float64
	Memory uint64
	Disk   uint64
}

// Role represents the role of the node
type Role string

const (
	// Worker is a node that runs tasks
	Worker Role = "worker"

	// Manager is a node that manages the cluster
	Manager Role = "manager"
)
