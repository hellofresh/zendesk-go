package zendesk

import (
	"testing"
)

var client = FromToken(
	LoadConfiguration("./config/configuration.yml"),
);

var getUserId int

func TestGetUsers(t *testing.T) {
	user, _ := client.ZendeskApi().GetUsers();

	if user[0].Id == 0 {
		t.Fail()
	}

	getUserId = user[0].Id
}

func TestGetSingleUser(t *testing.T) {
	user, _ := client.ZendeskApi().GetUser(getUserId)

	if getUserId != user.Id {
		t.Fail()
	}
}

var createdUserId int

func TestCreateUser(t *testing.T) {
	user := User {
		Name: "My user",
		Email: "my@user.com",
		Verified: true,
		Phone: "9928893",
	}

	result, _ := client.ZendeskApi().CreateUser(user)

	if result.Email != "my@user.com" {
		t.Fail()
	}

	createdUserId = result.Id
}

func TestUpdateUser(t *testing.T) {
	user, _ := client.ZendeskApi().GetUser(createdUserId)
	user.Name = "Updated user"
	user.Email = "updated@user.com"

	result, _ := client.ZendeskApi().UpdateUser(user)

	if result.Email != "updated@user.com" {
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {
	statusCode, _ := client.ZendeskApi().DeleteUser(createdUserId)

	if statusCode != 200 {
		t.Fail()
	}
}