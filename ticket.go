package zendesk

import (
	"fmt"
	"strings"

	resty "gopkg.in/resty.v0"
)

type TicketApiHandler struct {
	client Client
}

type SingleTicket struct {
	Response Ticket `json:"ticket"`
}

type MultipleTicket struct {
	Response     []Ticket `json:"tickets"`
	NextPage     string   `json:"next_page,omitempty"`
	PreviousPage string   `json:"previous_page,omitempty"`
	Count        int      `json:"count"`
}

func (t TicketApiHandler) GetById(id int) (Ticket, error) {
	response, err := t.client.get(
		fmt.Sprintf("/tickets/%d.json", id),
		nil,
	)

	if err != nil {

	}

	return t.parseSingleObject(response), err
}

func (t TicketApiHandler) GetAll(path string) ([]Ticket, error) {
	if path == "" {
		path = "/tickets.json"
	}

	response, err := t.client.get(
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return t.parseMultiObjects(response)
}

func (t TicketApiHandler) Create(v Ticket) (Ticket, error) {
	var object SingleTicket

	_, err := t.client.post(
		"/tickets.json",
		map[string]Ticket{"ticket": v},
		&object,
	)

	return object.Response, err
}

func (t TicketApiHandler) CreateOrUpdate(v Ticket) (Ticket, error) {
	return v, nil
}

func (t TicketApiHandler) CreateOrUpdateMany(v []Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (t TicketApiHandler) Update(v Ticket) (Ticket, error) {
	var object SingleTicket

	_, err := t.client.put(
		fmt.Sprintf("/tickets/%d.json", v.Id),
		map[string]Ticket{"ticket": v},
		&object,
	)

	return object.Response, err
}

func (t TicketApiHandler) Delete(id int) (int, error) {
	response, err := t.client.delete(
		fmt.Sprintf("/tickets/%d.json", id),
	)

	return response.StatusCode(), err
}

func (t TicketApiHandler) parseMultiObjects(response *resty.Response) ([]Ticket, error) {
	tickets := ManyTickets{}
	var er error

	var object MultipleTicket

	t.client.parseResponseToInterface(response, &object)

	tickets.ContcatTickets(object.Response)
	if object.NextPage != "" {
		slices := strings.Split(object.NextPage, "/")
		path := "/" + slices[len(slices)-1]

		tckts, err := t.GetAll(path)

		if err != nil {
			er = err
		} else {
			tickets.ContcatTickets(tckts)
		}
	}

	return tickets.Tickets, er
}

func (t TicketApiHandler) parseSingleObject(response *resty.Response) Ticket {
	var object SingleTicket

	t.client.parseResponseToInterface(response, &object)

	return object.Response
}
