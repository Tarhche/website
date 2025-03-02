package scheduler

// import (
// 	"github.com/khanzadimahdi/testproject/domain/runner/node"
// 	"github.com/khanzadimahdi/testproject/domain/runner/task"
// )

// type RoundRobin struct {
// 	Name       string
// 	lastWorker int
// }

// func (r *RoundRobin) NominatingNodes(t task.Task, nodes []*node.Node) []*node.Node {
// 	return nodes
// }

// func (r *RoundRobin) Score(t task.Task, nodes []node.Node) map[string]float64 {
// 	nodeScores := make(map[string]float64)
	
// 	var newWorker int
// 	if r.lastWorker+1 < len(nodes) {
// 		newWorker = r.lastWorker + 1
// 		r.lastWorker++
// 	} else {
// 		newWorker = 0
// 		r.lastWorker = 0
// 	}

// 	for idx, node := range nodes {
// 		if idx == newWorker {
// 			nodeScores[node.Info().Name] = 0.1
// 		} else {
// 			nodeScores[node.Info().Name] = 1.0
// 		}
// 	}

// 	return nodeScores
// }

// func (r *RoundRobin) Pick(scores map[string]float64, candidates []node.Node) node.Node {
// 	var bestNode node.Node
// 	var lowestScore float64

// 	for idx, node := range candidates {
// 		nodeName := node.Info().Name

// 		if idx == 0 {
// 			bestNode = node
// 			lowestScore = scores[nodeName]
// 			continue
// 		}

// 		if scores[nodeName] < lowestScore {
// 			bestNode = node
// 			lowestScore = scores[nodeName]
// 		}
// 	}

// 	return bestNode
// }
