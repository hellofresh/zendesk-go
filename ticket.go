package zendesk

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type SingleTicket struct {
	Response Ticket `json:"ticket"`
}

type MultipleTicket struct {
	Response []Ticket `json:"tickets"`
}

func (api *Api) GetTicket(id int) (Ticket, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/tickets/%d.json", id),
		nil,
		SingleTicket{},
	)

	var object SingleTicket
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

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
	response, err := api.postHttpRequest(
		"/tickets.json",
		map[string]Ticket{"ticket": ticket},
		SingleTicket{},
	)

	var object SingleTicket
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) UpdateTicket(ticket Ticket) (Ticket, error) {
	response, err := api.updateHttpRequest(
		fmt.Sprintf("/tickets/%d.json", ticket.Id),
		map[string]User{"ticket": ticket},
		SingleTicket{},
	)

	var object SingleTicket
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

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
	response, err := api.getHttpRequest(
		url,
		nil,
		MultipleTicket{},
	)

	var object MultipleTicket
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}