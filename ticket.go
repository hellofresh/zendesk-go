package zendesk

import (
	"fmt"
	"strings"
)

type SingleTicket struct {
	Response Ticket `json:"ticket"`
}

type MultipleTicket struct {
	Response []Ticket `json:"tickets"`
}

func (api *Api) GetTicket(id int) (Ticket, error) {
	var object SingleTicket

	response, err := api.getHttpRequest(
		fmt.Sprintf("/tickets/%d.json", id),
		nil,
	)

	api.parseResponseToInterface(response, &object)

	return object.Response, err
}

func (api *Api) GetTickets() ([]Ticket, error) {
	return api.getMultiple("/tickets.json");
}

func (api *Api) GetRecentTickets() ([]Ticket, error) {
	return api.getMultiple("/tickets/recent.json");
}

func (api *Api) GetTicketsFromOrganization(id int) ([]Ticket, error) {
	return api.getMultiple(
		fmt.Sprintf("/organizations/%d/tickets.json", id),
	)
}

func (api *Api) GetManyTickets(ids []string) ([]Ticket, error) {
	return api.getMultiple(
		fmt.Sprintf("/show_many.json?ids=%s", strings.Join(ids, ",")),
	)
}

func (api *Api) GetRequestedTicketsFromUser(id int) ([]Ticket, error) {
	return api.getMultiple(
		fmt.Sprintf("users/%d/tickets/requested.json", id),
	)
}

func (api *Api) GetCcdTicketsFromUser(id int) ([]Ticket, error) {
	return api.getMultiple(
		fmt.Sprintf("users/%d/tickets/ccd.json", id),
	)
}

func (api *Api) GetAssignedTicketsFromUser(id int) ([]Ticket, error) {
	return api.getMultiple(
		fmt.Sprintf("users/%d/tickets/assigned.json", id),
	)
}

func (api *Api) CreateTicket(ticket Ticket) (Ticket, error) {
	var object SingleTicket

	_, err := api.postHttpRequest(
		"/tickets.json",
		map[string]Ticket{"ticket": ticket},
		&object,
	)

	return object.Response, err
}

func (api *Api) UpdateTicket(ticket Ticket) (Ticket, error) {
	var object SingleTicket

	_, err := api.updateHttpRequest(
		fmt.Sprintf("/tickets/%d.json", ticket.Id),
		map[string]Ticket{"ticket": ticket},
		&object,
	)

	return object.Response, err
}

func (api *Api) UpdateTicketMarkAsSpam(id int) (int, error) {
	response, err := api.updateHttpRequestGetResponse(
		fmt.Sprintf("/tickets/%d.json", id),
		nil,
	)

	return response.StatusCode(), err
}

func (api *Api) DeleteTicket(id int) (int, error) {
	response, err := api.deleteHttpRequest(
		fmt.Sprintf("/tickets/%d.json", id),
	)

	return response.StatusCode(), err
}

func (api *Api) getMultiple(url string) ([]Ticket, error) {
	var object MultipleTicket

	response, err := api.getHttpRequest(
		url,
		nil,
	)

	api.parseResponseToInterface(response, &object)

	return object.Response, err
}