package docker

// import (
// 	"context"
// 	"io"
// 	"log"
// 	"math"
// 	"os"

// 	"github.com/khanzadimahdi/testproject/domain/runner/port"
// 	"github.com/khanzadimahdi/testproject/domain/runner/task"

// 	"github.com/docker/docker/api/types/container"
// 	"github.com/docker/docker/api/types/image"
// 	"github.com/docker/docker/client"
// 	"github.com/docker/go-connections/nat"
// 	"github.com/moby/moby/pkg/stdcopy"
// )

// type Config struct {
// 	Name         string
// 	AttachStdin  bool
// 	AttachStdout bool
// 	AttachStderr bool
// 	ExposedPorts port.PortSet
// 	Cmd          []string
// 	Image        string
// 	Cpu          float64
// 	// Memory in MiB
// 	Memory int64
// 	// Disk in GiB
// 	Disk          int64
// 	Env           []string
// 	RestartPolicy string
// }

// func NewConfig(t *task.Task) *Config {
// 	return &Config{
// 		Name:          t.Name,
// 		AttachStdin:   t.AttachStdin,
// 		AttachStdout:  t.AttachStdout,
// 		AttachStderr:  t.AttachStderr,
// 		ExposedPorts:  t.ExposedPorts,
// 		Cmd:           t.Cmd,
// 		Image:         t.Image,
// 		Cpu:           t.Cpu,
// 		Memory:        t.Memory,
// 		Disk:          t.Disk,
// 		Env:           t.Env,
// 		RestartPolicy: t.RestartPolicy,
// 	}
// }

// type Docker struct {
// 	Client *client.Client
// 	Config Config
// }

// var _ task.RunStopper = &Docker{}

// func NewDocker(c *Config) *Docker {
// 	dc, _ := client.NewClientWithOpts(client.FromEnv)
// 	return &Docker{
// 		Client: dc,
// 		Config: *c,
// 	}
// }

// func (d *Docker) Run() (string, error) {
// 	ctx := context.Background()
// 	reader, err := d.Client.ImagePull(ctx, d.Config.Image, image.PullOptions{})
// 	if err != nil {
// 		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
// 		return "", err
// 	}
// 	io.Copy(os.Stdout, reader)

// 	rp := container.RestartPolicy{
// 		Name: container.RestartPolicyMode(d.Config.RestartPolicy),
// 	}

// 	r := container.Resources{
// 		Memory:   d.Config.Memory,
// 		NanoCPUs: int64(d.Config.Cpu * math.Pow(10, 9)),
// 	}

// 	hc := container.HostConfig{
// 		RestartPolicy:   rp,
// 		Resources:       r,
// 		PublishAllPorts: true,
// 	}

// 	exposedPorts := make(nat.PortSet, len(d.Config.ExposedPorts))
// 	for i := range d.Config.ExposedPorts {
// 		exposedPorts[nat.Port(i)] = d.Config.ExposedPorts[i]
// 	}

// 	resp, err := d.Client.ContainerCreate(ctx, &container.Config{
// 		Image:        d.Config.Image,
// 		Tty:          false,
// 		Env:          d.Config.Env,
// 		ExposedPorts: exposedPorts,
// 	}, &hc, nil, nil, d.Config.Name)
// 	if err != nil {
// 		log.Printf("Error creating container using image %s: %v\n", d.Config.Image, err)
// 		return "", err
// 	}

// 	if err := d.Client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
// 		log.Printf("Error starting container %s: %v\n", resp.ID, err)
// 		return "", err
// 	}

// 	out, err := d.Client.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
// 	if err != nil {
// 		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
// 		return "", err
// 	}

// 	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

// 	return resp.ID, nil
// }

// func (d *Docker) Stop(containerID string) error {
// 	log.Printf("Attempting to stop container %v", containerID)
// 	ctx := context.Background()
// 	err := d.Client.ContainerStop(ctx, containerID, container.StopOptions{})
// 	if err != nil {
// 		log.Printf("Error stopping container %s: %v\n", containerID, err)
// 		return err
// 	}

// 	return d.Remove(containerID)
// }

// func (d *Docker) Remove(containerID string) error {
// 	log.Printf("Attempting to remove container %v", containerID)
// 	ctx := context.Background()

// 	err := d.Client.ContainerRemove(ctx, containerID, container.RemoveOptions{
// 		RemoveVolumes: true,
// 		RemoveLinks:   false,
// 		Force:         false,
// 	})

// 	return err
// }

// func (d *Docker) Inspect(containerID string) (task.Container, error) {
// 	ctx := context.Background()

// 	resp, err := d.Client.ContainerInspect(ctx, containerID)
// 	if err != nil {
// 		log.Printf("Error inspecting container: %s\n", err)
// 		return task.Container{ID: resp.ID}, err
// 	}

// 	return task.Container{
// 		ID:            resp.ID,
// 		Name:          resp.Name,
// 		Image:         resp.Config.Image,
// 		Status:        resp.State.Status,
// 		Cpu:           float64(resp.HostConfig.NanoCPUs) / math.Pow(10, 9),
// 		Memory:        resp.HostConfig.Memory,
// 		RestartPolicy: string(resp.HostConfig.RestartPolicy.Name),
// 		Env:           resp.Config.Env,
// 		Cmd:           resp.Config.Cmd,
// 	}, nil
// }
