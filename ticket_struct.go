package zendesk

type Via struct{}

type Ticket struct {
	Id                  int           `json:"id,omitempty"`
	Url                 string        `json:"url,omitempty"`
	ExternalId          string        `json:"external_id,omitempty"`
	Type                string        `json:"type,omitempty"`
	Subject             string        `json:"subject,omitempty"`
	RawSubject          string        `json:"raw_subject,omitempty"`
	Description         string        `json:"description,omitempty"`
	Priority            string        `json:"priority,omitempty"`
	Status              string        `json:"status,omitempty"`
	Recipient           string        `json:"recipient,omitempty"`
	RequesterId         int64         `json:"requester_id,omitempty"`
	SubmitterId         int           `json:"submitter_id,omitempty"`
	AssigneeId          int           `json:"assignee_id,omitempty"`
	OrganizationId      int           `json:"organization_id,omitempty"`
	GroupId             int           `json:"group_id,omitempty"`
	CollaboratorsId     []int         `json:"collaborators_id,omitempty"`
	ForumTopicId        int           `json:"forum_topic_id,omitempty"`
	ProblemId           int           `json:"problem_id,omitempty"`
	HasIncidents        bool          `json:"has_incidents,omitempty"`
	DueAt               string        `json:"due_at,omitempty"`
	Tags                []string      `json:"tags,omitempty"`
	Via                 *Via          `json:"via,omitempty"`
	CustomFields        []CustomField `json:"custom_fields,omitempty"`
	SatisfactionRating  []string      `json:"satisfaction_rating,omitempty"`
	SharingAgreementIds []int         `json:"sharing_agreement_ids,omitempty"`
	FollowupIds         []int         `json:"followup_ids,omitempty"`
	TicketFormId        int           `json:"ticket_form_id,omitempty"`
	BrandId             int           `json:"brand_id,omitempty"`
	CreatedAt           string        `json:"created_at,omitempty"`
	UpdatedAt           string        `json:"updated_at,omitempty"`
}

type CustomField struct {
	ID    int    `json:"id"`
	Value interface{} `json:"value"`
}
