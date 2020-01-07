package main

import (
	"flag"

	"github.com/romwod/toq/task"
	"github.com/romwod/toq/producer"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	redisPool := NewRedisPool()
	p := producer.NewProducer(redisPool)
	for i := 0; i < 15; i++ {
		taskID := UUID4()
		logrus.Infof("enqueue task %s", taskID)
		t := task.Task{ID: taskID, Retry: true, MaxRetries: 1, Key: "test_key", Args: "{}"}
		p.Enqueue("test_toq_queue", t)
	}
}
