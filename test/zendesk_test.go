package test

import (
	"testing"

	zendesk ".."
	models "../models"
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
		id = *user.ID
		break
	}
}

func TestUserApiHandler_GetByID(t *testing.T) {
	_, err := client.User().GetByID(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_Create(t *testing.T) {
	user := models.User{}

	*user.Name = "Felipe Pieretti Umpierre"
	*user.Email = "fum@hellofresh.com"

	_, err := client.User().Create(user)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_CreateOrUpdate(t *testing.T) {
	user := models.User{}

	*user.Name = "Felipe Pieretti Umpierre = Updated"
	*user.Email = "fum@hellofresh.com"

	_, err := client.User().CreateOrUpdate(user)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestUserApiHandler_CreateOrUpdateMany(t *testing.T) {
	var many models.ManyUsers

	user := models.User{}

	*user.Name = "User 1"
	*user.Email = "user-1@hellofresh.com"

	many.AppendUsers(user)

	*user.Name = "User 2"
	*user.Email = "user-2@hellofresh.com"

	many.AppendUsers(user)

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
	user := models.User{}
	*user.ID = id
	*user.Name = "Felipe Pieretti Umpierre - hallo"
	*user.Email = "fum@hellofresh.com"

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
		id = *ticket.ID
		break
	}
}

func TestTicketApiHandler_GetById(t *testing.T) {
	_, err := client.Ticket().GetByID(id)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_Create(t *testing.T) {
	ticket := models.Ticket{}
	*ticket.Description = "Test ticket - Created"

	_, err := client.Ticket().Create(ticket)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}

func TestTicketApiHandler_CreateOrUpdate(t *testing.T) {
	ticket := models.Ticket{}
	*ticket.Description = "Test ticket - Created OR Updated"

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
	ticket := models.Ticket{}
	*ticket.Description = "Test ticket"
	ticket.ID = &id

	_, err := client.Ticket().Update(ticket)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.Fail()
	}
}
