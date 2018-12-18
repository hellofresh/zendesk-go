package zendesk

import (
	"fmt"
	"log"

	resty "gopkg.in/resty.v0"
)

type SearchApiHandler struct {
	client Client
}

func (s SearchApiHandler) SearchTickets(searchString string, sortBy string, sortOrder string) (TicketSearch, error) {
	urlString := fmt.Sprintf("/search.json?query=type:ticket %s", searchString)

	if len(sortBy) > 0 {
		urlString = fmt.Sprintf(urlString+"&sort_by=%s", sortBy)
	}

	if len(sortBy) > 0 {
		urlString = fmt.Sprintf(urlString+"&sort_order=%s", sortOrder)
	}

	response, err := s.client.get(
		urlString,
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s SearchApiHandler) NextPage(t TicketSearch) (TicketSearch, error) {
	response, err := s.client.get(
		s.client.toPath(t.NextPage),
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}

	return s.parseResults(response), err
}

func (s SearchApiHandler) PrevPage(t TicketSearch) (TicketSearch, error) {
	response, err := s.client.get(
		s.client.toPath(t.PrevPage),
		nil,
	)

	if err != nil {
		return TicketSearch{}, err
	}

	return s.parseResults(response), err
}

func (s SearchApiHandler) parseResults(response *resty.Response) TicketSearch {
	var object TicketSearch

	s.client.parseResponseToInterface(response, &object)

	return object
}
