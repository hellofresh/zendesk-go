package test

import (
	"testing"
	"zendesk-go"
)

var client = zendesk.FromToken(
	zendesk.LoadConfiguration("./../config/configuration.yml"),
)

var id int

// --------- USERS --------

func TestUserApiHandler_GetAll(t *testing.T) {
	users, err := client.User().GetAll()

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}

	for _, user := range users {
		id = int(user.Id)
		break
	}
}

func TestUserApiHandler_GetById(t *testing.T) {
	_, err := client.User().GetById(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_GetAllAgents(t *testing.T) {
	_, err := client.User().GetAllAgents()

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_Create(t *testing.T) {
	user := zendesk.User{
		Name:  "Felipe Pieretti Umpierre",
		Email: "fum@hellofresh.com",
	}

	_, err := client.User().Create(user)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_CreateOrUpdate(t *testing.T) {
	user := zendesk.User{
		Name:  "Felipe Pieretti Umpierre = Updated",
		Email: "fum@hellofresh.com",
	}

	_, err := client.User().CreateOrUpdate(user)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_CreateOrUpdateMany(t *testing.T) {
	var many zendesk.ManyUsers

	many.AppendUsers(zendesk.User{
		Name:  "User 1",
		Email: "user-1@hellofresh.com",
	})

	many.AppendUsers(zendesk.User{
		Name:  "User-2",
		Email: "user-2@hellofresh.com",
	})

	_, err := client.User().CreateOrUpdateMany(many.Users)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_Delete(t *testing.T) {
	_, err := client.User().Delete(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_Update(t *testing.T) {
	user := zendesk.User{
		Id:    int64(id),
		Name:  "Felipe Pieretti Umpierre - hallo",
		Email: "fum@hellofresh.com",
	}

	_, err := client.User().Update(user)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

// --------- TICKET --------

func TestTicketApiHandler_GetAll(t *testing.T) {
	tickets, err := client.Ticket().GetAll()

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}

	for _, ticket := range tickets {
		id = ticket.Id
		break
	}
}

func TestTicketApiHandler_GetById(t *testing.T) {
	_, err := client.Ticket().GetById(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_Create(t *testing.T) {
	ticket := zendesk.Ticket{
		Description: "Test ticket - Created",
	}

	_, err := client.Ticket().Create(ticket)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_CreateOrUpdate(t *testing.T) {
	ticket := zendesk.Ticket{
		Description: "Test ticket - Created OR Updated",
	}

	_, err := client.Ticket().CreateOrUpdate(ticket)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_Delete(t *testing.T) {
	_, err := client.Ticket().Delete(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_Update(t *testing.T) {
	ticket := zendesk.Ticket{
		Id:          id,
		Description: "Test ticket",
	}

	_, err := client.Ticket().Update(ticket)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}
