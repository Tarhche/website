package worker

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/golang-collections/collections/queue"
// 	"github.com/khanzadimahdi/testproject/domain/runner/stats"
// 	"github.com/khanzadimahdi/testproject/domain/runner/task"
// 	"github.com/khanzadimahdi/testproject/infrastructure/docker"
// )

// type Worker struct {
// 	Name           string
// 	Queue          queue.Queue
// 	TaskRepository task.Repository
// 	Stats          *stats.Stats
// 	StatsCollector stats.Collector
// 	TaskCount      int
// }

// func New(
// 	name string,
// 	taskRepository task.Repository,
// 	statsCollector stats.Collector,
// ) *Worker {
// 	return &Worker{
// 		Name:           name,
// 		Queue:          *queue.New(),
// 		TaskRepository: taskRepository,
// 		StatsCollector: statsCollector,
// 	}
// }

// func (w *Worker) GetTasks(offset uint, limit uint) []task.Task {
// 	taskList, err := w.TaskRepository.GetAll(offset, limit)
// 	if err != nil {
// 		log.Printf("error getting list of tasks: %v", err)
// 		return nil
// 	}

// 	return taskList
// }

// func (w *Worker) CollectStats() {
// 	var (
// 		stats stats.Stats
// 		err   error
// 	)

// 	for {
// 		log.Println("Collecting stats")
// 		if stats, err = w.StatsCollector.Collect(); err != nil {
// 			continue
// 		}
// 		w.Stats = &stats

// 		time.Sleep(15 * time.Second)
// 	}
// }

// func (w *Worker) AddTask(t task.Task) {
// 	w.Queue.Enqueue(t)
// }

// func (w *Worker) RunTasks() {
// 	for {
// 		if w.Queue.Len() != 0 {
// 			if _, err := w.runTask(); err != nil {
// 				log.Printf("Error running task: %v\n", err)
// 			}
// 		} else {
// 			log.Printf("No tasks to process currently.\n")
// 		}
// 		log.Println("Sleeping for 10 seconds.")
// 		time.Sleep(10 * time.Second)
// 	}
// }

// func (w *Worker) runTask() (string, error) {
// 	t := w.Queue.Dequeue()
// 	if t == nil {
// 		log.Println("[worker] No tasks in the queue")

// 		return "", errors.New("no tasks in the queue")
// 	}

// 	taskQueued := t.(task.Task)
// 	fmt.Printf("[worker] Found task in queue: %v:\n", taskQueued)

// 	UUID, err := w.TaskRepository.Save(&taskQueued)
// 	if err != nil {
// 		msg := fmt.Errorf("error storing task %s: %v", taskQueued.ID, err)
// 		log.Println(msg)
// 		return UUID, err
// 	}

// 	taskPersisted, err := w.TaskRepository.GetOne(UUID)
// 	if err != nil {
// 		msg := fmt.Errorf("error getting task %s from database: %v", taskQueued.ID, err)
// 		log.Println(msg)
// 		return UUID, err
// 	}

// 	if taskPersisted.State == task.Completed {
// 		return "", w.StopTask(taskPersisted)
// 	}

// 	if task.ValidStateTransition(taskPersisted.State, taskQueued.State) {
// 		switch taskQueued.State {
// 		case task.Scheduled:
// 			if taskQueued.ContainerID != "" {
// 				if err := w.StopTask(taskQueued); err != nil {
// 					log.Printf("%v\n", err)
// 				}
// 			}
// 			return w.StartTask(taskQueued)
// 		default:
// 			fmt.Printf("This is a mistake. taskPersisted: %v, taskQueued: %v\n", taskPersisted, taskQueued)
// 			return "", errors.New("We should not get here")
// 		}
// 	}

// 	return "", fmt.Errorf("Invalid transition from %v to %v", taskPersisted.State, taskQueued.State)
// }

// func (w *Worker) StartTask(t task.Task) (string, error) {
// 	config := docker.NewConfig(&t)
// 	d := docker.NewDocker(config)
// 	containerID, err := d.Run()
// 	if err != nil {
// 		log.Printf("Err running task %v: %v\n", t.ID, err)
// 		t.State = task.Failed

// 		if _, err := w.TaskRepository.Save(&t); err != nil {
// 			return containerID, err
// 		}

// 		return containerID, err
// 	}

// 	t.ContainerID = containerID
// 	t.State = task.Running
// 	if _, err := w.TaskRepository.Save(&t); err != nil {
// 		return "", err
// 	}

// 	return containerID, nil
// }

// func (w *Worker) StopTask(t task.Task) error {
// 	config := docker.NewConfig(&t)
// 	d := docker.NewDocker(config)

// 	stopResult := d.Stop(t.ContainerID)
// 	if stopResult.Error != nil {
// 		log.Printf("%v\n", stopResult.Error)
// 	}
// 	removeResult := d.Remove(t.ContainerID)
// 	if removeResult.Error != nil {
// 		log.Printf("%v\n", removeResult.Error)
// 	}

// 	t.FinishTime = time.Now().UTC()
// 	t.State = task.Completed
// 	log.Printf("Stopped and removed container %v for task %v\n", t.ContainerID, t.ID)

// 	if _, err := w.TaskRepository.Save(&t); err != nil {
// 		return err
// 	}

// 	return removeResult
// }

// func (w *Worker) InspectTask(t task.Task) (task.Container, error) {
// 	config := docker.NewConfig(&t)
// 	d := docker.NewDocker(config)

// 	return d.Inspect(t.ContainerID)
// }

// func (w *Worker) UpdateTasks() {
// 	for {
// 		log.Println("Checking status of tasks")
// 		w.updateTasks()
// 		log.Println("Task updates completed")
// 		log.Println("Sleeping for 15 seconds")
// 		time.Sleep(15 * time.Second)
// 	}
// }

// func (w *Worker) updateTasks() {
// 	// for each task in the worker's datastore:
// 	// 1. call InspectTask method
// 	// 2. verify task is in running state
// 	// 3. if task is not in running state, or not running at all, mark task as `failed`
// 	tasks, err := w.TaskRepository.List()
// 	if err != nil {
// 		log.Printf("error getting list of tasks: %v", err)
// 		return
// 	}

// 	for _, t := range tasks.([]*task.Task) {
// 		if t.State == task.Running {
// 			inspect, err := w.InspectTask(*t)
// 			if err != nil {
// 				fmt.Printf("ERROR: %v", err)
// 			}

// 			if inspect.Container == nil {
// 				log.Printf("No container for running task %s", t.ID)
// 				t.State = task.Failed
// 				w.TaskRepository.Save(t)
// 			}

// 			if resp.Container.State.Status == "exited" {
// 				log.Printf("Container for task %s in non-running state %s", t.ID, resp.Container.State.Status)
// 				t.State = task.Failed
// 				w.TaskRepository.Save(t)
// 			}

// 			// task is running, update exposed ports
// 			t.HostPorts = resp.Container.NetworkSettings.NetworkSettingsBase.Ports
// 			w.TaskRepository.Save(t)
// 		}
// 	}
// }
