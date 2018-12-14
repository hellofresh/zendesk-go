package zendesk

import "time"

type SatisfactionRatings struct {
	Count               int                  `json:"count,omitempty"`
	NextPage            string               `json:"next_page,omitempty"`
	PrevPage            string               `json:"prev_page,omitempty"`
	SatisfactionRatings []SatisfactionRating `json:"satisfaction_ratings,omitempty"`
}

type SatisfactionRating struct {
	Id          int        `json:"id,omitempty"`
	Url         string     `json:"url,omitempty"`
	AssigneeId  int        `json:"assignee_id,omitempty"`
	GroupId     int        `json:"group_id,omitempty"`
	RequesterId int        `json:"requester_id,omitempty"`
	TicketId    int        `json:"ticket_id,omitempty"`
	Score       string     `json:"score,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
