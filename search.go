package zendesk

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

type SearchApiHandler struct {
	client Client
}

func (s SearchApiHandler) SearchTickets(searchString string) (TicketSearch, error) {
	response, err := s.client.get(
		fmt.Sprintf("/search.json?query=type:ticket %s", searchString),
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s SearchApiHandler) parseResults(response *resty.Response) TicketSearch {
	var object TicketSearch

	s.client.parseResponseToInterface(response, &object)

	return object
}
