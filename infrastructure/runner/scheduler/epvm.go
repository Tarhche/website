package scheduler

// import (
// 	"math"

// 	"github.com/khanzadimahdi/testproject/domain/runner/task"
// )

// const (
// 	// LIEB square ice constant
// 	// https://en.wikipedia.org/wiki/Lieb%27s_square_ice_constant
// 	LIEB = 1.53960071783900203869
// )

// // Enhanced PVM (Parallel Virtual Machine) Algorithm
// //
// // Implementation of the E-PVM algorithm laid out in http://www.cnds.jhu.edu/pub/papers/mosix.pdf.
// // The algorithm calculates the "marginal cost" of assigning a task to a machine. In the paper and
// // in this implementation, the only resources considered for calculating a task's marginal cost are
// // memory and cpu.
// type Epvm struct {
// 	Name string
// }

// func (e *Epvm) NominatingNodes(t task.Task, nodes []node.State) []node.Node {
// 	var candidates []node.Node
// 	for node := range nodes {
// 		if checkDisk(t, nodes[node].Disk-nodes[node].DiskAllocated) {
// 			candidates = append(candidates, nodes[node])
// 		}

// 	}

// 	return candidates
// }

// func (e *Epvm) Score(t task.Task, nodes []node.Node) map[string]float64 {
// 	nodeScores := make(map[string]float64)
// 	maxJobs := 4.0

// 	for _, node := range nodes {
// 		cpuUsage := calculateCpuUsage(node)
// 		cpuLoad := calculateLoad(*cpuUsage, math.Pow(2, 0.8))

// 		memoryAllocated := float64(node.Stats.MemUsedKb()) + float64(node.MemoryAllocated)
// 		memoryPercentAllocated := memoryAllocated / float64(node.Memory)

// 		newMemPercent := (calculateLoad(memoryAllocated+float64(t.Memory/1000), float64(node.Memory)))
// 		memCost := math.Pow(LIEB, newMemPercent) + math.Pow(LIEB, (float64(node.TaskCount+1))/maxJobs) - math.Pow(LIEB, memoryPercentAllocated) - math.Pow(LIEB, float64(node.TaskCount)/float64(maxJobs))
// 		cpuCost := math.Pow(LIEB, cpuLoad) + math.Pow(LIEB, (float64(node.TaskCount+1))/maxJobs) - math.Pow(LIEB, cpuLoad) - math.Pow(LIEB, float64(node.TaskCount)/float64(maxJobs))

// 		nodeScores[node.Name] = memCost + cpuCost
// 	}

// 	return nodeScores
// }

// func (e *Epvm) Pick(scores map[string]float64, candidates []node.State) node.Node {
// 	minCost := 0.00

// 	var bestNode node.Node
// 	for idx, node := range candidates {
// 		if idx == 0 {
// 			minCost = scores[node.Name]
// 			bestNode = node
// 			continue
// 		}
// 		if scores[node.Name] < minCost {
// 			minCost = scores[node.Name]
// 			bestNode = node
// 		}
// 	}

// 	return bestNode
// }
