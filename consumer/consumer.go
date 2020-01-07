package consumer

import (
	"github.com/romwod/toq/task"
)

type Dequeuer interface {
	Dequeue() error
}

type Worker func(task.Task) task.Result

type Consumer interface {
	RegisterWorker(string, Worker) error
	Run() error
	SetResult(task.Result)
}
