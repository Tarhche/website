package scheduler

// import (
// 	"time"

// 	"github.com/khanzadimahdi/testproject/domain/runner/node"
// 	"github.com/khanzadimahdi/testproject/domain/runner/task"
// )

// type Scheduler interface {
// 	NominateNodes(t task.Task, nodes []*node.Node) []*node.Node
// 	Score(t task.Task, nodes []*node.Node) map[string]float64
// 	Pick(scores map[string]float64, candidates []*node.Node) *node.Node
// }

// func checkDisk(t task.Task, diskAvailable uint64) bool {
// 	return t.Disk <= diskAvailable
// }

// func calculateLoad(usage float64, capacity float64) float64 {
// 	return usage / capacity
// }

// func calculateCpuUsage(node node.Node) float64 {
// 	stat1 := node.Stats()
// 	time.Sleep(3 * time.Second)
// 	stat2 := node.Stats()

// 	stat1Idle := stat1.CPU.Idle + stat1.CPU.IOWait
// 	stat2Idle := stat2.CPU.Idle + stat2.CPU.IOWait
// 	stat1NonIdle := stat1.CPU.User + stat1.CPU.Nice + stat1.CPU.System + stat1.CPU.IRQ + stat1.CPU.SoftIRQ + stat1.CPU.Steal
// 	stat2NonIdle := stat2.CPU.User + stat2.CPU.Nice + stat2.CPU.System + stat2.CPU.IRQ + stat2.CPU.SoftIRQ + stat2.CPU.Steal
// 	stat1Total := stat1Idle + stat1NonIdle
// 	stat2Total := stat2Idle + stat2NonIdle
// 	total := stat2Total - stat1Total
// 	idle := stat2Idle - stat1Idle

// 	var cpuPercentUsage float64
// 	if total == 0 && idle == 0 {
// 		cpuPercentUsage = 0.00
// 	} else {
// 		cpuPercentUsage = (float64(total) - float64(idle)) / float64(total)
// 	}

// 	return cpuPercentUsage
// }
