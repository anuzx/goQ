package job

import (
	"context"
	"encoding/json"
	"time"
)

type Status string

const (
	StatusPending Status = "PENDING"
	StatusRunning Status = "RUNNING"
	StatusFailed  Status = "FAILED"
	StatusSuccess Status = "SUCCESS"
)

type Job struct {
	ID          string          `json:"id"`
	Type        string          `json:"type"`
	Payload     json.RawMessage `json:"payload"`
	Status      Status          `json:"status"`
	Attempts    int             `json:"attempts"`
	MaxAttempts int             `json:"max_attempts"`
	LastError   string          `json:"last_error,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// Handler contains the behavior for a particular kind of job.
type Handler interface {
	Execute(ctx context.Context, payload json.RawMessage) error
}
