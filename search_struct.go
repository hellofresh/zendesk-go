package zendesk

type TicketSearch struct {
	Count    int      `json:"count,omitempty"`
	NextPage string   `json:"next_page,omitempty"`
	PrevPage string   `json:"prev_page,omitempty"`
	Tickets  []Ticket `json:"results,omitempty"`
}
