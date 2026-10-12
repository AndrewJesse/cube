package worker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/golang-collections/collections/queue"

	"cube/task"
)

type Worker struct {
	Name string
	Queue *queue.Queue
	Db    map[uuid.UUID]*task.Task
	TaskCount int
}