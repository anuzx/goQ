package worker

import (
	"fmt"
	"sync"

	"github.com/anuzx/goQ/internal/job"
)

type Registry struct {
	mu       sync.RWMutex
	handlers map[string]job.Handler
}

func NewRegistry() *Registry {
	return &Registry{
		handlers: make(map[string]job.Handler),
	}
}

//adds a handler and rejects duplicate job types instead of silently replacing an existing handler 
func (r *Registry) Register(jobType string, handler job.Handler) error {

	if jobType == "" {
		return fmt.Errorf("job type cannot be empty")
	}

	if handler == nil {
		return fmt.Errorf("handler cannot be nil")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.handlers[jobType]; exists {
		return fmt.Errorf("handler already registered for job type %q", jobType)
	}

	r.handlers[jobType] = handler
	return nil
}

//retrieves a handler and returns a boolean indication whether it exists 
func (r *Registry) Get(jobType string) (job.Handler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	handler, exists := r.handlers[jobType]
	return handler, exists
}
