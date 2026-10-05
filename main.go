package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Execute() error
}

type JobQueue struct {
	jobs    chan Job
	workers int
	wg      sync.WaitGroup
}

type SampleJob struct {
	ID int
}

func NewJobQueue(buffer int, workers int) *JobQueue {
	q := &JobQueue{
		jobs:    make(chan Job, buffer),
		workers: workers,
	}

	for i := 0; i < q.workers; i++ {
		go func(workerID int) {
			for job := range q.jobs {
				fmt.Printf("Worker %d picked a job", workerID)
				job.Execute()
				q.wg.Done()
			}
		}(i + 1)
	}

	return q
}

func (s *SampleJob) Execute() error {
	startTime := time.Now()
	fmt.Printf("Job %d started at %s", s.ID, startTime.Format("15:04:05.000"))

	time.Sleep(2 * time.Second) // simulate slowness

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Job %d finished at %s (duration: %v)", s.ID, endTime.Format("15:04:05.000"), duration)
	return nil
}
func (q *JobQueue) Submit(job Job) {
	q.wg.Add(1)
	q.jobs <- job
}

func (q *JobQueue) Wait() {
	q.wg.Wait()
}

func (q *JobQueue) ShutDown() {
	close(q.jobs)
}

func main() {
	start := time.Now()

	buffer, workers := 10, 3

	q := job.NewJobQueue(buffer, workers)

	for i := 1; i <= 5; i++ {
		q.Submit(&job.SampleJob{ID: i})
	}

	q.Wait()
	q.Shutdown()

	fmt.Printf("Total execution time: %.2fs", time.Since(start).Seconds())
}
