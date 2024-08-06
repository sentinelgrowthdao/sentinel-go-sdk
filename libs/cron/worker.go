package cron

import (
	"time"
)

// Worker defines the interface for a scheduler worker.
type Worker interface {
	Interval() time.Duration       // Returns the interval at which the worker should be run.
	MaxRetries() int               // Returns the maximum number of retry attempts for the worker.
	Name() string                  // Returns the name of the worker.
	OnError(err error) (stop bool) // Handles errors that occur during worker execution. Returns true to stop the worker, false otherwise.
	Run() error                    // Executes the worker and returns an error if it fails.
}

// Ensure BasicWorker implements the Worker interface.
var _ Worker = (*BasicWorker)(nil)

// BasicWorker provides a basic implementation of the Worker interface.
type BasicWorker struct {
	handler    func() error
	interval   time.Duration
	maxRetries int
	name       string
	onError    func(error) bool
}

// NewBasicWorker creates a new BasicWorker with default settings.
func NewBasicWorker() *BasicWorker {
	return &BasicWorker{
		onError: func(error) bool { return false }, // Default behavior: do not stop on error.
	}
}

// WithHandler sets the handler function for the worker.
func (bw *BasicWorker) WithHandler(handler func() error) *BasicWorker {
	bw.handler = handler
	return bw
}

// WithInterval sets the interval at which the worker should be run.
func (bw *BasicWorker) WithInterval(interval time.Duration) *BasicWorker {
	bw.interval = interval
	return bw
}

// WithMaxRetries sets the maximum number of retry attempts for the worker.
func (bw *BasicWorker) WithMaxRetries(maxRetries int) *BasicWorker {
	bw.maxRetries = maxRetries
	return bw
}

// WithName sets the name of the worker.
func (bw *BasicWorker) WithName(name string) *BasicWorker {
	bw.name = name
	return bw
}

// WithOnError sets the function to handle errors that occur during worker execution.
func (bw *BasicWorker) WithOnError(onError func(error) bool) *BasicWorker {
	bw.onError = onError
	return bw
}

// Interval returns the interval at which the worker should be executed.
func (bw *BasicWorker) Interval() time.Duration {
	return bw.interval
}

// MaxRetries returns the maximum number of retry attempts for the worker.
func (bw *BasicWorker) MaxRetries() int {
	return bw.maxRetries
}

// Name returns the name of the worker.
func (bw *BasicWorker) Name() string {
	return bw.name
}

// OnError processes errors encountered during worker execution.
func (bw *BasicWorker) OnError(err error) bool {
	return bw.onError(err)
}

// Run executes the worker's handler function and returns any error encountered.
func (bw *BasicWorker) Run() error {
	return bw.handler()
}
