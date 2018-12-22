package zendesk

/*
(\s+)\[JsonProperty\("([^"]+)".+\n\s+public ([^ ]+) ([^ ]+).+
$1$4 $3 `json:"$2,omitempty"`
*/
type a struct {
	Address            string   `json:"address,omitempty"`
	Name               string   `json:"name,omitempty"`
	FormattedPhone     string   `json:"formatted_phone,omitempty"`
	Phone              string   `json:"phone,omitempty"`
	ProfileURL         string   `json:"profile_url,omitempty"`
	FacebookID         string   `json:"facebook_id,omitempty"`
	TicketID           int      `json:"ticket_id,omitempty"`
	Subject            string   `json:"subject,omitempty"`
	OriginalRecipients []string `json:"original_recipients,omitempty"`
}
