package zendesk

import (
	"fmt"
	"gopkg.in/resty.v0"
)

type TicketApiHandler struct {
	client Client
}

type SingleTicket struct {
	Response Ticket `json:"ticket"`
}

type MultipleTicket struct {
	Response []Ticket `json:"tickets"`
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

func (t TicketApiHandler) GetAll() ([]Ticket, error) {
	response, err := t.client.get(
		"/tickets.json",
		nil,
	)

	if err != nil {
	}

	return t.parseMultiObjects(response), err
}

func (t TicketApiHandler) GetRequestedByUser(userId int64) ([]Ticket, error) {
	response, err := t.client.get(
		fmt.Sprintf("/users/%d/tickets/requested.json", userId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return t.parseMultiObjects(response), err
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

func (t TicketApiHandler) parseMultiObjects(response *resty.Response) []Ticket {
	var object MultipleTicket

	t.client.parseResponseToInterface(response, &object)

	return object.Response
}

func (t TicketApiHandler) parseSingleObject(response *resty.Response) Ticket {
	var object SingleTicket

	t.client.parseResponseToInterface(response, &object)

	return object.Response
}
