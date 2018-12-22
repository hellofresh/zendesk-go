package models

import (
	"time"

	"../shared"
)

// ManyTickets struct agreggator
type ManyTickets struct {
	Tickets []Ticket
}

// ManyComments struct agreggator
type ManyComments struct {
	Comments []Comment
}

// Field struct definition
type Field struct {
	ID    *int    `json:"id,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Custom Field Definition
type CustomField struct {
	*Field
}

// SatisfactionRating struct info
type SatisfactionRating struct {
	Score    *string `json:"score,omitempty"`
	Comment  *string `json:"comment,omitempty"`
	ID       *int    `json:"id,omitempty"`
	Reason   *string `json:"reason,omitempty"`
	ReasonID *int    `json:"reason_id,omitempty"`
}

// Comment data
type Comment struct {
	ID                   *int                `json:"id,omitempty"`
	Type                 *string             `json:"type,omitempty"`
	Body                 *string             `json:"body,omitempty"`
	HTMLBody             *string             `json:"html_body,omitempty"`
	PlainBody            *string             `json:"plain_body,omitempty"`
	Public               *bool               `json:"public,omitempty"`
	AuthorID             *int                `json:"author_id,omitempty"`
	Uploads              []string            `json:"uploads,omitempty"`
	Attachments          []shared.Attachment `json:"attachments,omitempty"`
	Via                  *shared.Via         `json:"via,omitempty"`
	CreatedAt            *time.Time          `json:"created_at,omitempty"`
	Author               *User               `json:"author,omitempty"`
	Data                 *shared.Data        `json:"data,omitempty"`
	FormattedFrom        *string             `json:"formatted_from,omitempty"`
	FormattedTo          *string             `json:"formatted_to,omitempty"`
	TranscriptionVisible *bool               `json:"transcription_visible,omitempty"`
	Trusted              *bool               `json:"trusted,omitempty"`
	Metadata             *shared.MetaData    `json:"metadata,omitempty"`
	AuditID              *int                `json:"audit_id,omitempty"`
}

// TimeSpanMetric struct
type TimeSpanMetric struct {
	Calendar *int64 `json:"calendar,omitempty"`
	Business *int64 `json:"business,omitempty"`
}

// TicketMetric struct
type TicketMetric struct {
	TicketID                     *int            `json:"ticket_id,omitempty"`
	GroupStations                *int            `json:"group_stations,omitempty"`
	AssigneeStations             *int            `json:"assignee_stations,omitempty"`
	Reopens                      *int            `json:"reopens,omitempty"`
	Replies                      *int            `json:"replies,omitempty"`
	AssigneeUpdatedAt            *time.Time      `json:"assignee_updated_at,omitempty"`
	RequesterUpdatedAt           *time.Time      `json:"requester_updated_at,omitempty"`
	StatusUpdatedAt              *time.Time      `json:"status_updated_at,omitempty"`
	InitiallyAssignedAt          *time.Time      `json:"initially_assigned_at,omitempty"`
	AssignedAt                   *time.Time      `json:"assigned_at,omitempty"`
	SolvedAt                     *time.Time      `json:"solved_at,omitempty"`
	LatestCommentAddedAt         *time.Time      `json:"latest_comment_added_at,omitempty"`
	FirstResolutionTimeInMinutes *TimeSpanMetric `json:"first_resolution_time_in_minutes,omitempty"`
	ReplyTimeInMinutes           *TimeSpanMetric `json:"reply_time_in_minutes,omitempty"`
	FullResolutionTimeInMinutes  *TimeSpanMetric `json:"full_resolution_time_in_minutes,omitempty"`
	AgentWaitTimeInMinutes       *TimeSpanMetric `json:"agent_wait_time_in_minutes,omitempty"`
	RequesterWaitTimeInMinutes   *TimeSpanMetric `json:"requester_wait_time_in_minutes,omitempty"`
	OnHoldTimeInMinutes          *TimeSpanMetric `json:"on_hold_time_in_minutes,omitempty"`
}

// Requester info
type Requester struct {
	LocaleID *int    `json:"locale_id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
}

// Ticket struct data
type Ticket struct {
	*shared.Individual

	ExternalID          *string             `json:"external_id,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Description         *string             `json:"description,omitempty"`
	Subject             *string             `json:"subject,omitempty"`
	Priority            *string             `json:"priority,omitempty"`
	Status              *string             `json:"status,omitempty"`
	Recipient           *string             `json:"recipient,omitempty"`
	RequesterID         *int                `json:"requester_id,omitempty"`
	SubmitterID         *int                `json:"submitter_id,omitempty"`
	AssigneeID          *int                `json:"assignee_id,omitempty"`
	OrganizationID      *int                `json:"organization_id,omitempty"`
	GroupID             *int                `json:"group_id,omitempty"`
	CollaboratorIds     []int               `json:"collaborator_ids,omitempty"`
	CollaboratorEmails  []string            `json:"collaborators,omitempty"`
	ForumTopicID        *string             `json:"forum_topic_id,omitempty"`
	ProblemID           *string             `json:"problem_id,omitempty"`
	HasIncidents        *bool               `json:"has_incidents,omitempty"`
	DueAt               *time.Time          `json:"due_at,omitempty"`
	Tags                []string            `json:"tags,omitempty"`
	CustomFields        []CustomField       `json:"custom_fields,omitempty"`
	SatisfactionRating  *SatisfactionRating `json:"satisfaction_rating,omitempty"`
	BrandID             *int                `json:"brand_id,omitempty"`
	Via                 *shared.Via         `json:"via,omitempty"`
	Comment             *Comment            `json:"comment,omitempty"`
	Requester           *Requester          `json:"requester,omitempty"`
	TicketFormID        *int                `json:"ticket_form_id,omitempty"`
	IsPublic            *bool               `json:"is_public,omitempty"`
	ViaFollowupSourceID *int                `json:"via_followup_source_id,omitempty"`
	Comments            []Comment           `json:"comments,omitempty"`
	User                *User               `json:"user,omitempty"`
	Audits              []shared.Audit      `json:"audits,omitempty"`
	TicketMetric        *TicketMetric       `json:"ticket_metric,omitempty"`
	FollowerIDs         []int               `json:"follower_ids,omitempty"`
	EmailCcIDs          []string            `json:"email_cc_ids,omitempty"`
	SharingAgreementIDs []string            `json:"sharing_agreement_ids,omitempty"`
	Fields              []Field             `json:"fields,omitempty"`
	FollowupIDs         []string            `json:"followup_ids,omitempty"`
	AllowChannelback    *bool               `json:"allow_channelback,omitempty"`
	AllowAttachments    *bool               `json:"allow_attachments,omitempty"`
	RawSubject          *string             `json:"raw_subject,omitempty"`
}

// ContcatTickets function concats many tickets with many tickets
func (tickets *ManyTickets) ContcatTickets(manyTickets []Ticket) []Ticket {
	tickets.Tickets = append(tickets.Tickets, manyTickets...)

	return tickets.Tickets
}

// ContcatComments function concats many comments with many comments
func (comments *ManyComments) ContcatComments(manyComments []Comment) []Comment {
	comments.Comments = append(comments.Comments, manyComments...)

	return comments.Comments
}
