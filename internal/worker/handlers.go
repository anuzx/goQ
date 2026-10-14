package worker

import (
	"context"
	"encoding/json"
	"fmt"
)

type PrintPayload struct {
	Message string `json:"message"`
}

type PrintHandler struct {
}

// accepts a json, decodes the payload into a typed struct and returns errors instead of crashing when the payload is invalid
func (h *PrintHandler) Execute(
	ctx context.Context,
	payload json.RawMessage,
) error {
	var data PrintPayload

	if err := json.Unmarshal(payload, &data); err != nil {
		return fmt.Errorf("decode print payload: %w", err)
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	fmt.Println(data.Message)
	return nil
}
