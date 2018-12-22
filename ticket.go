package zendesk

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"./models"

	resty "gopkg.in/resty.v0"
)

// TicketAPIHandler struct
type TicketAPIHandler struct {
	client Client
}

// SingleTicket result
type SingleTicket struct {
	Response models.Ticket `json:"ticket"`
}

// MultipleTicket result
type MultipleTicket struct {
	Response     []models.Ticket `json:"tickets"`
	NextPage     string          `json:"next_page,omitempty"`
	PreviousPage string          `json:"previous_page,omitempty"`
	Count        int             `json:"count"`
}

// MultipleComment result
type MultipleComment struct {
	Response     []models.Comment `json:"comments"`
	NextPage     string           `json:"next_page,omitempty"`
	PreviousPage string           `json:"previous_page,omitempty"`
	Count        int              `json:"count"`
}

// GetByID find a ticket by it's own id
func (t TicketAPIHandler) GetByID(id int) (models.Ticket, error) {
	response, err := t.client.get(
		fmt.Sprintf("/tickets/%d.json", id),
		nil,
	)

	if err != nil {

	}

	return t.parseSingleObject(response), err
}

// GetAll find all tickets on zendesk
func (t TicketAPIHandler) GetAll() ([]models.Ticket, error) {
	path := "/tickets.json"

	return t.GetAllFrom(path)
}

// GetAllFrom find all tickets begining a specific page
func (t TicketAPIHandler) GetAllFrom(path string) ([]models.Ticket, error) {
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

// GetAllComments find all comments by ticket id
func (t TicketAPIHandler) GetAllComments(id int) ([]models.Comment, error) {
	path := fmt.Sprintf("/tickets/%d/comments.json", id)

	return t.GetAllCommentsFrom(path)
}

// GetAllCommentsFrom find all comments begining a specific page
func (t TicketAPIHandler) GetAllCommentsFrom(path string) ([]models.Comment, error) {
	if path == "" {
		panic("Path not found!")
	}

	response, err := t.client.get(
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return t.parseMultiObjectComments(response)
}

// Create a new ticket
func (t TicketAPIHandler) Create(v models.Ticket) (models.Ticket, error) {
	var object SingleTicket

	_, err := t.client.post(
		"/tickets.json",
		map[string]models.Ticket{"ticket": v},
		&object,
	)

	return object.Response, err
}

// CreateOrUpdate a ticket
func (t TicketAPIHandler) CreateOrUpdate(v models.Ticket) (models.Ticket, error) {
	return v, nil
}

// CreateOrUpdateMany tickets
func (t TicketAPIHandler) CreateOrUpdateMany(v []models.Ticket) (models.Ticket, error) {
	return models.Ticket{}, nil
}

// Update a ticket
func (t TicketAPIHandler) Update(v models.Ticket) (models.Ticket, error) {
	var object SingleTicket

	_, err := t.client.put(
		fmt.Sprintf("/tickets/%d.json", *v.ID),
		map[string]models.Ticket{"ticket": v},
		&object,
	)

	return object.Response, err
}

// Delete a ticket
func (t TicketAPIHandler) Delete(id int) (int, error) {
	response, err := t.client.delete(
		fmt.Sprintf("/tickets/%d.json", id),
	)

	return response.StatusCode(), err
}

// LoadTicketComments func
func (t TicketAPIHandler) LoadTicketComments(ticket *models.Ticket) (bool, error) {
	comments, err := t.GetAllComments(*ticket.ID)

	if err != nil {
		panic(err)
	}

	ticket.Comments = comments

	return (err != nil), err
}

// LoadComments for many tickets
func (t TicketAPIHandler) LoadComments(tickets []models.Ticket) (bool, error) {
	procs := runtime.NumCPU()

	var wg sync.WaitGroup

	var er error

	if procs > 1 {
		procs--
	}

	wg.Add(len(tickets))

	runtime.GOMAXPROCS(procs)

	for index := 0; index < len(tickets); index++ {
		go func(t TicketAPIHandler, index int, er *error) {
			_, err := t.LoadTicketComments(&tickets[index])

			defer wg.Done()

			if err != nil {
				*er = err
			}
		}(t, index, &er)
	}

	wg.Wait()

	return (er != nil), er
}

// parseMultipleObjects
func (t TicketAPIHandler) parseMultiObjects(response *resty.Response) ([]models.Ticket, error) {
	tickets := models.ManyTickets{}
	var er error

	var object MultipleTicket

	t.client.parseResponseToInterface(response, &object)

	tickets.ContcatTickets(object.Response)
	if object.NextPage != "" {
		slices := strings.Split(object.NextPage, "/")
		path := "/" + slices[len(slices)-1]

		tckts, err := t.GetAllFrom(path)

		if err != nil {
			er = err
		} else {
			tickets.ContcatTickets(tckts)
		}
	}

	return tickets.Tickets, er
}

// parseMultiObjectComments
func (t TicketAPIHandler) parseMultiObjectComments(response *resty.Response) ([]models.Comment, error) {
	comments := models.ManyComments{}
	var er error

	var object MultipleComment

	t.client.parseResponseToInterface(response, &object)

	comments.ContcatComments(object.Response)
	if object.NextPage != "" {
		slices := strings.Split(object.NextPage, "/")
		path := "/" + slices[len(slices)-1]

		cmmts, err := t.GetAllCommentsFrom(path)

		if err != nil {
			er = err
		} else {
			comments.ContcatComments(cmmts)
		}
	}

	return comments.Comments, er
}

// parseSingleObject
func (t TicketAPIHandler) parseSingleObject(response *resty.Response) models.Ticket {
	var object SingleTicket

	t.client.parseResponseToInterface(response, &object)

	return object.Response
}

// AppendUser from users
func (t TicketAPIHandler) AppendUser(ticket *models.Ticket, users []models.User) error {
	var author, assign *models.User

	for _, user := range users {
		if user.ID == ticket.RequesterID {
			author = &user
		}

		if user.ID == ticket.AssigneeID {
			assign = &user
		}
	}

	for _, comment := range ticket.Comments {
		comment.Author = author
	}

	ticket.User = assign

	return nil
}
