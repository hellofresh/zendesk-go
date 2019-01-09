package zendesk

type PhoneNumber struct {
	Id            int64  `json:"id,omitempty"`
	DisplayNumber string `json:"display_number,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
}
