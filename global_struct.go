package zendesk

type ZendeskConfiguration struct {
	ApiVersion string `yaml:"api_version"`
	Domain     string `yaml:"domain"`
	Email      string `yaml:"email"`
	Token      string `yaml:"token"`
	Password   string `yaml:"password"`
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