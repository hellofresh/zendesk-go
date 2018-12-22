package shared

import "time"

// The Individual implementation for all responses
type Individual struct {
	ID        *int       `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
