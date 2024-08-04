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

// Ensure BasicJob implements the Job interface.
var _ Job = (*BasicJob)(nil)

// BasicJob is a basic implementation of the Job interface.
type BasicJob struct {
	handler  func() error
	interval time.Duration
	name     string
	onError  func(error) bool
}

// NewBasicJob initializes a BasicJob with default values.
func NewBasicJob() *BasicJob {
	return &BasicJob{
		onError: func(error) bool { return false }, // Default implementation: do not stop on error
	}
}

// WithHandler sets the handler function for the job.
func (j *BasicJob) WithHandler(handler func() error) *BasicJob {
	j.handler = handler
	return j
}

// WithInterval sets the interval for the job.
func (j *BasicJob) WithInterval(interval time.Duration) *BasicJob {
	j.interval = interval
	return j
}

// WithName sets the name for the job.
func (j *BasicJob) WithName(name string) *BasicJob {
	j.name = name
	return j
}

// WithOnError sets the error handling function for the job.
func (j *BasicJob) WithOnError(onError func(error) bool) *BasicJob {
	j.onError = onError
	return j
}

// Name returns the name of the job.
func (j *BasicJob) Name() string {
	return j.name
}

// Run executes the job's handler function.
func (j *BasicJob) Run() error {
	return j.handler()
}

// Interval returns the interval at which the job should be run.
func (j *BasicJob) Interval() time.Duration {
	return j.interval
}

// OnError handles errors that occur during the job's execution.
func (j *BasicJob) OnError(err error) bool {
	return j.onError(err)
}
