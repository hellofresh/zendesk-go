package zendesk

import (
	"time"
)

type TicketMetrics struct {
	Count         int            `json:"count,omitempty"`
	NextPage      string         `json:"next_page,omitempty"`
	PrevPage      string         `json:"prev_page,omitempty"`
	TicketMetrics []TicketMetric `json:"ticket_metrics,omitempty"`
}

type InMinutes struct {
	Calendar int `json:"calendar,omitempty"`
	Business int `json:"business,omitempty"`
}

type TicketMetric struct {
	Id                           int        `json:"id,omitempty"`
	TicketId                     int        `json:"ticket_id,omitempty"`
	Url                          string     `json:"url,omitempty"`
	GroupStations                int        `json:"group_stations,omitempty"`
	AssigneeStations             int        `json:"assignee_stations,omitempty"`
	Reopens                      int        `json:"reopens,omitempty"`
	Replies                      int        `json:"replies,omitempty"`
	AssigneeUpdatedAt            *time.Time `json:"assignee_updated_at,omitempty"`
	RequesterUpdatedAt           *time.Time `json:"requester_updated_at,omitempty"`
	StausUpdatedAt               *time.Time `json:"status_updated_at,omitempty"`
	InitiallyAssignedAt          *time.Time `json:"initially_assigned_at,omitempty"`
	AssignedAt                   *time.Time `json:"assigned_at,omitempty"`
	SolvedAt                     *time.Time `json:"solved_at,omitempty"`
	LastCommentAddedAt           *time.Time `json:"last_comment_added_at,omitempty"`
	FirstResolutionTimeInMinutes InMinutes  `json:"first_resolution_time_in_minutes,omitempty"`
	ReplyTimeInMinutes           InMinutes  `json:"reply_time_in_minutes,omitempty"`
	FullResolutionTimeInMinutes  InMinutes  `json:"full_resolution_time_in_minutes,omitempty"`
	AgentWaitTimeInMinutes       InMinutes  `json:"agent_wait_time_in_minutes,omitempty"`
	RequesterWaitTimeInMinutes   InMinutes  `json:"requester_wait_time_in_minutes,omitempty"`
	CreatedAt                    *time.Time `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time `json:"updated_at,omitempty"`
}
