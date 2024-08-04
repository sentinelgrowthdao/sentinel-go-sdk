package cron

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Scheduler manages the scheduling and execution of jobs.
type Scheduler struct {
	isRunning  bool
	jobs       map[string]Job
	stopSignal chan struct{}
	mu         sync.Mutex
	wg         sync.WaitGroup
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs:       make(map[string]Job),
		stopSignal: make(chan struct{}),
	}
}

// Start begins executing all jobs in separate goroutines.
func (s *Scheduler) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return errors.New("scheduler is already started")
	}

	s.isRunning = true

	// Start existing jobs
	for _, job := range s.jobs {
		s.wg.Add(1)
		go s.runLoop(job)
	}

	return nil
}

// Stop halts the execution of all jobs and stops the scheduler.
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return
	}

	// Stop all jobs
	close(s.stopSignal)
	s.wg.Wait()
	s.isRunning = false
}

// RegisterJobs adds multiple jobs to the scheduler.
func (s *Scheduler) RegisterJobs(jobs ...Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return errors.New("cannot add new jobs while scheduler is running")
	}

	for _, job := range jobs {
		if _, exists := s.jobs[job.Name()]; exists {
			return fmt.Errorf("job with name %q already exists", job.Name())
		}

		s.jobs[job.Name()] = job
	}

	return nil
}

// runLoop executes a job's function in a loop and handles errors using the OnError method.
func (s *Scheduler) runLoop(j Job) {
	defer s.wg.Done()

	for {
		select {
		case <-s.stopSignal:
			return
		default:
			if err := j.Run(); err != nil {
				if j.OnError(err) {
					return
				}
			}

			// Prevent tight loop if job has a zero interval
			interval := j.Interval()
			if interval > 0 {
				time.Sleep(interval)
			}
		}
	}
}
