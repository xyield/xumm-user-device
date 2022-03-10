package device

import "fmt"

type PayloadNotFoundError struct {
	UUID string
}

func (e *PayloadNotFoundError) Error() string {
	return fmt.Sprintf("Payload with UUID %v not found", e.UUID)
}
