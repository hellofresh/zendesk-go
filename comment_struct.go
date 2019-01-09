package zendesk

type Comment struct {
	Id        int64  `json:"id,omitempty"`
	AuthorId  int64  `json:"author_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
