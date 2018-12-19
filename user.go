package zendesk

import (
	"fmt"
	"strings"

	resty "gopkg.in/resty.v0"
)

type UserApiHandler struct {
	client Client
}

type SingleUser struct {
	Response User `json:"user"`
}

type MultipleUser struct {
	Response     []User `json:"users"`
	NextPage     string `json:next_page,omitempty`
	PreviousPage string `json:previous_page,omitempty`
	Count        int    `json:count`
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

func (u UserApiHandler) GetAll(path string) ([]User, error) {
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

func (u UserApiHandler) parseMultiObjects(response *resty.Response) ([]User, error) {
	users := ManyUsers{}
	var er error

	var object MultipleUser

	u.client.parseResponseToInterface(response, &object)

	users.ContcatUsers(object.Response)
	if object.NextPage != "" {
		slices := strings.Split(object.NextPage, "/")
		path := "/" + slices[len(slices)-1]

		usrs, err := u.GetAll(path)

		if err != nil {
			er = err
		} else {
			users.ContcatUsers(usrs)
		}
	}

	return users.Users, er
}

func (u UserApiHandler) parseSingleObject(response *resty.Response) User {
	var object SingleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}
