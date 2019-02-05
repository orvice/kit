package utils

import (
	"time"

	"github.com/orvice/kit/log"
)

type Task struct {
	name   string
	t      time.Duration
	fn     func() error
	
	Logger log.Logger
}

type TaskOpt func(t *Task)

func NewTask(name string, t time.Duration, fn func() error, opts ...TaskOpt) *Task {
	task := &Task{
		name:   name,
		t:      t,
		fn:     fn,
		Logger: log.NewDefaultLogger(),
	}
	for _, opt := range opts {
		opt(task)
	}
	return task
}

func (t *Task) Run() {
	for {
		start := time.Now()
		err := t.fn()
		t.Logger.Infof("run task %s cost: %v error: %v  ", t.name, err, time.Since(start))
		time.Sleep(t.t)
	}
}
