package zendesk

// Configuration struct
type Configuration struct {
	APIVersion string `yaml:"api_version"`
	Domain     string `yaml:"domain"`
	Email      string `yaml:"email"`
	Token      string `yaml:"token"`
	Password   string `yaml:"password"`
}

// Job struct
type Job struct {
	JobStatus `json:"job_status,omitempty"`
}

// JobStatus struct
type JobStatus struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Total    string `json:"total,omitempty"`
	Progress string `json:"progress,omitempty"`
	Status   string `json:"status,omitempty"`
	Message  string `json:"mesage,omitempty"`
	Results  string `json:"results,omitempty"`
}
