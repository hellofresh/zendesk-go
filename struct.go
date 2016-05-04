package zendesk

type ZendeskConfiguration struct {
	ApiVersion string `yaml:"api_version"`
	Domain     string `yaml:"domain"`
	Email      string `yaml:"email"`
	Token      string `yaml:"token"`
	Password   string `yaml:"password"`
}

type Attachment struct{}
type Via struct{}

type ManyUsers struct {
	Users []User
}

type User struct {
	Id                  int`json:"id,omitempty"`
	Url                 string `json:"url,omitempty"`
	Name                string `json:"name,omitempty"`
	ExternalId          string `json:"external_id,omitempty"`
	Alias               string `json:"alias,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
	Active              bool `json:"active,omitempty"`
	Verified            bool `json:"verified,omitempty"`
	Shared              bool `json:"shared,omitempty"`
	SharedAgent         bool `json:"shared_agent,omitempty"`
	Locale              string `json:"locale,omitempty"`
	LocaleId            int `json:"locale_id,omitempty"`
	TimeZone            string `json:"time_zone,omitempty"`
	LastLoginAt         string `json:"last_login_at,omitempty"`
	Email               string `json:"email,omitempty"`
	Phone               string `json:"phone,omitempty"`
	Signature           string `json:"signature,omitempty"`
	Details             string `json:"details,omitempty"`
	Notes               string `json:"notes,omitempty"`
	OrganizationId      int `json:"organization_id,omitempty"`
	Role                string `json:"role,omitempty"`
	CustomRoleId        string `json:"custom_role_id,omitempty"`
	Moderator           bool `json:"moderator,omitempty"`
	TicketRestriction   string `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments bool `json:"only_private_comments,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	Suspended           bool `json:"suspended,omitempty"`
	RestrictedAgent     bool `json:"restricted_agent,omitempty"`
	Photo               *Attachment `json:"photo,omitempty"`
}

type Ticket struct {
	Id                  int `json:"id,omitempty"`
	Url                 string `json:"url,omitempty"`
	ExternalId          string `json:"external_id,omitempty"`
	Type                string `json:"type,omitempty"`
	Subject             string `json:"subject,omitempty"`
	RawSubject          string `json:"raw_subject,omitempty"`
	Description         string `json:"description,omitempty"`
	Priority            string `json:"priority,omitempty"`
	Status              string `json:"status,omitempty"`
	Recipient           string `json:"recipient,omitempty"`
	RequesterId         int `json:"requester_id,omitempty"`
	SubmitterId         int `json:"submitter_id,omitempty"`
	AssigneeId          int `json:"assignee_id,omitempty"`
	OrganizationId      int `json:"organization_id,omitempty"`
	GroupId             int `json:"group_id,omitempty"`
	CollaboratorsId     []int `json:"collaborators_id,omitempty"`
	ForumTopicId        int `json:"forum_topic_id,omitempty"`
	ProblemId           int `json:"problem_id,omitempty"`
	HasIncidents        bool `json:"has_incidents,omitempty"`
	DueAt               string `json:"due_at,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	Via                 *Via `json:"via,omitempty"`
	CustomFields        []string `json:"custom_fields,omitempty"`
	SatisfactionRating  []string `json:"satisfaction_rating,omitempty"`
	SharingAgreementIds []int `json:"sharing_agreement_ids,omitempty"`
	FollowupIds         []int `json:"followup_ids,omitempty"`
	TicketFormId        int `json:"ticket_form_id,omitempty"`
	BrandId             int `json:"brand_id,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
}

type Job struct {
	JobStatus `json:"job_status,omitempty"`
}

type JobStatus struct {
	Id       string `json:"id,omitempty"`
	Url      string `json:"url,omitempty"`
	Total    string `json:"total,omitempty"`
	Progress string `json:"progress,omitempty"`
	Status   string `json:"status,omitempty"`
	Message  string `json:"mesage,omitempty"`
	Results  string `json:"results,omitempty"`
}