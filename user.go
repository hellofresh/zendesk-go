package zendesk

import (
	"fmt"
	"strings"

	"./models"

	resty "gopkg.in/resty.v0"
)

// UserAPIHandler struct
type UserAPIHandler struct {
	client Client
}

// SingleUser result
type SingleUser struct {
	Response models.User `json:"user"`
}

// MultipleUser result
type MultipleUser struct {
	Response     []models.User `json:"users"`
	NextPage     string        `json:"next_page,omitempty"`
	PreviousPage string        `json:"previous_page,omitempty"`
	Count        int           `json:"count"`
}

// GetByID find an user by id
func (u UserAPIHandler) GetByID(id int) (models.User, error) {
	response, err := u.client.get(
		fmt.Sprintf("/users/%d.json", id),
		nil,
	)

	if err != nil {

	}

	return u.parseSingleObject(response), err
}

// GetAll find all users
func (u UserAPIHandler) GetAll() ([]models.User, error) {
	path := "/users.json"

	return u.GetAllFrom(path)
}

// GetAllFrom find all users begining by page
func (u UserAPIHandler) GetAllFrom(path string) ([]models.User, error) {
	if path == "" {
		path = "/users.json"
	}

	response, err := u.client.get(
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return u.parseMultiObjects(response)
}

// Create an user
func (u UserAPIHandler) Create(v models.User) (models.User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users.json",
		map[string]models.User{"user": v},
		&object,
	)

	return object.Response, err
}

// CreateOrUpdate an user
func (u UserAPIHandler) CreateOrUpdate(v models.User) (models.User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users/create_or_update.json",
		map[string]models.User{"user": v},
		&object,
	)

	return object.Response, err
}

// CreateOrUpdateMany users
func (u UserAPIHandler) CreateOrUpdateMany(v []models.User) (Job, error) {
	var object Job

	_, err := u.client.post(
		"/users/create_or_update_many.json",
		map[string][]models.User{"users": v},
		&object,
	)

	return object, err
}

// Update an user
func (u UserAPIHandler) Update(v models.User) (models.User, error) {
	var object SingleUser

	_, err := u.client.put(
		fmt.Sprintf("/users/%d.json", *v.ID),
		map[string]models.User{"user": v},
		&object,
	)

	return object.Response, err
}

// Delete an user
func (u UserAPIHandler) Delete(id int) (int, error) {
	response, err := u.client.delete(
		fmt.Sprintf("/users/%d.json", id),
	)

	return response.StatusCode(), err
}

// parseMultiObjects
func (u UserAPIHandler) parseMultiObjects(response *resty.Response) ([]models.User, error) {
	users := models.ManyUsers{}
	var er error

	var object MultipleUser

	u.client.parseResponseToInterface(response, &object)

	users.ContcatUsers(object.Response)
	if object.NextPage != "" {
		slices := strings.Split(object.NextPage, "/")
		path := "/" + slices[len(slices)-1]

		usrs, err := u.GetAllFrom(path)

		if err != nil {
			er = err
		} else {
			users.ContcatUsers(usrs)
		}
	}

	return users.Users, er
}

// parseSingleObject
func (u UserAPIHandler) parseSingleObject(response *resty.Response) models.User {
	var object SingleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}
