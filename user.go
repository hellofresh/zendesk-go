package zendesk

import (
	"fmt"
	"gopkg.in/resty.v0"
)

type UserApiHandler struct {
	client Client
}

type SingleUser struct {
	Response User `json:"user"`
}

type MultipleUser struct {
	Response []User `json:"users"`
}

func (u UserApiHandler) GetById(id int) (User, error) {
	response, err := u.client.get(
		fmt.Sprintf("/users/%d.json", id),
		nil,
	)

	if err != nil {

	}

	return u.parseSingleObject(response), err
}

func (u UserApiHandler) GetAll() ([]User, error) {
	response, err := u.client.get(
		"/users.json",
		nil,
	)

	if err != nil {

	}

	return u.parseMultiObjects(response), err
}

func (u UserApiHandler) GetAllAgents() ([]User, error) {
	response, err := u.client.get(
		"/users.json",
		map[string]string{"role": "admin"},
	)

	if err != nil {
		return nil, err
	}

	users := u.parseMultiObjects(response)

	// This could be done in a single API call with role[]=agent&role[]=admin but the current interface makes doing this difficult
	response, err = u.client.get(
		"/users.json",
		map[string]string{"role": "agent"},
	)

	if err != nil {
		return nil, err
	}

	adminUsers := u.parseMultiObjects(response)
	users = append(users, adminUsers...)

	return users, err
}


func (u UserApiHandler) Create(v User) (User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users.json",
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) CreateOrUpdate(v User) (User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users/create_or_update.json",
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) CreateOrUpdateMany(v []User) (Job, error) {
	var object Job

	_, err := u.client.post(
		"/users/create_or_update_many.json",
		map[string][]User{"users": v},
		&object,
	)

	return object, err
}

func (u UserApiHandler) Update(v User) (User, error) {
	var object SingleUser

	_, err := u.client.put(
		fmt.Sprintf("/users/%d.json", v.Id),
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) Delete(id int) (int, error) {
	response, err := u.client.delete(
		fmt.Sprintf("/users/%d.json", id),
	)

	return response.StatusCode(), err
}

func (u UserApiHandler) parseMultiObjects(response *resty.Response) []User {
	var object MultipleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}

func (u UserApiHandler) parseSingleObject(response *resty.Response) User {
	var object SingleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}
