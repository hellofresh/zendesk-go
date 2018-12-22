package shared

import "time"

// From info
type From struct {
	Address            *string  `json:"address,omitempty"`
	Name               *string  `json:"name,omitempty"`
	FormattedPhone     *string  `json:"formatted_phone,omitempty"`
	Phone              *string  `json:"phone,omitempty"`
	ProfileURL         *string  `json:"profile_url,omitempty"`
	FacebookID         *string  `json:"facebook_id,omitempty"`
	TicketID           *int     `json:"ticket_id,omitempty"`
	Subject            *string  `json:"subject,omitempty"`
	OriginalRecipients []string `json:"original_recipients,omitempty"`
}

// To info
type To struct {
	*From
}

// Via channel info
type Via struct {
	Channel *string `json:"channel,omitempty"`
	Source  *Source `json:"source,omitempty"`
}

// Source info
type Source struct {
	From *From   `json:"from,omitempty"`
	To   *To     `json:"to,omitempty"`
	Rel  *string `json:"rel,omitempty"`
}

// FacebookPage description
type FacebookPage struct {
	Name    *string `json:"name,omitempty"`
	GraphID *string `json:"graph_id,omitempty"`
}

// Event data
type Event struct {
	ID                   *int          `json:"id,omitempty"`
	Type                 *string       `json:"type,omitempty"`
	Body                 *string       `json:"body,omitempty"`
	Public               *bool         `json:"public,omitempty"`
	Attachments          []Attachment  `json:"attachments,omitempty"`
	AuthorID             *int          `json:"author_id,omitempty"`
	HTMLBody             *string       `json:"html_body,omitempty"`
	Trusted              *bool         `json:"trusted,omitempty"`
	Data                 *Data         `json:"data,omitempty"`
	FormattedFrom        *string       `json:"formatted_from,omitempty"`
	TranscriptionVisible *string       `json:"transcription_visible,omitempty"`
	CommentID            *int          `json:"comment_id,omitempty"`
	FieldName            *string       `json:"field_name,omitempty"`
	Value                *string       `json:"value,omitempty"`
	PreviousValue        *string       `json:"previous_value,omitempty"`
	Subject              *string       `json:"subject,omitempty"`
	Via                  *Via          `json:"via,omitempty"`
	Recipients           []string      `json:"recipients,omitempty"`
	Message              *string       `json:"message,omitempty"`
	Success              *string       `json:"success,omitempty"`
	Resource             *string       `json:"resource,omitempty"`
	Communication        *int          `json:"communication,omitempty"`
	TicketVia            *string       `json:"ticket_via,omitempty"`
	FacebookPage         *FacebookPage `json:"page,omitempty"`
	ValueReference       *string       `json:"value_reference,omitempty"`
	AssigneeID           *string       `json:"assignee_id,omitempty"`
	Score                *string       `json:"score,omitempty"`
	DirectMessage        *string       `json:"direct_message,omitempty"`
	PhoneNumber          *string       `json:"phone_number,omitempty"`
	RecipientID          *string       `json:"recipient_id,omitempty"`
	AgreementID          *string       `json:"agreement_id,omitempty"`
	Action               *string       `json:"action,omitempty"`
}

// Audit info
type Audit struct {
	ID        *int       `json:"id,omitempty"`
	TicketID  *string    `json:"ticket_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	AuthorID  *int       `json:"author_id,omitempty"`
	MetaData  *MetaData  `json:"metadata,omitempty"`
	Via       *Via       `json:"via,omitempty"`
	Events    []Event    `json:"events,omitempty"`
}
