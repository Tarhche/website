package scheduler

// import (
// 	"math"

// 	"github.com/khanzadimahdi/testproject/domain/runner/node"
// 	"github.com/khanzadimahdi/testproject/domain/runner/task"
// )

// type Greedy struct {
// 	Name string
// }

// func (g *Greedy) NominatingNodes(t task.Task, nodes []node.Node) []node.Node {
// 	var candidates []node.Node
// 	for node := range nodes {
// 		if checkDisk(t, nodes[node].Disk.Available) {
// 			candidates = append(candidates, nodes[node])
// 		}
// 	}

// 	return candidates
// }

// func (g *Greedy) Score(t task.Task, nodes []node.Node) map[string]float64 {
// 	nodeScores := make(map[string]float64)

// 	for _, node := range nodes {
// 		cpuUsage := calculateCpuUsage(node)
// 		cpuLoad := calculateLoad(float64(cpuUsage), math.Pow(2, 0.8))
// 		nodeScores[node.Name] = cpuLoad
// 	}

// 	return nodeScores
// }

// func (g *Greedy) Pick(candidates map[string]float64, nodes []node.Node) node.Node {
// 	minCpu := 0.00

// 	var bestNode node.Node
// 	for idx, node := range nodes {
// 		if idx == 0 {
// 			minCpu = candidates[node.Name]
// 			bestNode = node
// 			continue
// 		}

// 		if candidates[node.Name] < minCpu {
// 			minCpu = candidates[node.Name]
// 			bestNode = node
// 		}
// 	}

// 	return bestNode
// }
