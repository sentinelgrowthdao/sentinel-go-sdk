package cron

import (
	"time"
)

// Job defines the interface for a scheduler job.
type Job interface {
	Run() error
	Name() string
	Interval() time.Duration
	OnError(err error) (stop bool)
}
