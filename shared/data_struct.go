package shared

import "time"

// Data information
type Data struct {
	From                *string    `json:"from,omitempty"`
	To                  *string    `json:"to,omitempty"`
	RecordingURL        *string    `json:"recording_url,omitempty"`
	CallID              *int       `json:"call_id,omitempty"`
	CallDuration        *int       `json:"call_duration,omitempty"`
	AnsweredByID        *int       `json:"answered_by_id,omitempty"`
	StartedAt           *time.Time `json:"started_at,omitempty"`
	Location            *string    `json:"location,omitempty"`
	AuthorID            *int       `json:"author_id,omitempty"`
	Public              *bool      `json:"public,omitempty"`
	BrandID             *int       `json:"brand_id,omitempty"`
	ViaID               *int       `json:"via_id,omitempty"`
	AnsweredByName      *string    `json:"answered_by_name,omitempty"`
	TranscriptionStatus *string    `json:"transcription_status,omitempty"`
}
