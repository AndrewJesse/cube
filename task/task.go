package task

type State int

const (
		Pending State = iota
		Running
		Completed
		Failed
)

import (
	"github.com/google/uuid"
	"github.com/docker/go-connections/nat"
	"time"
)	

type Task struct {
	ID    uuid.UUID
	Name  string
	State State
	Image string
	Memory int
	Disk   int
	ExposedPorts nat.PortSet
	PortBindings map[string]string
	RestartPolicy string
	StartTime time.Time
	EndTime   time.Time
}

type TaskEvent struct {
	ID uuid.UUID
	State State
	Timestamp time.Time
	Task Task
}