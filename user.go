package zendesk

import (
	"fmt"
	"gopkg.in/resty.v0"
)

type SingleUser struct {
	Response User `json:"user"`
}

type MultipleUser struct {
	Response []User `json:"users"`
}

func (api *Api) GetUser(id int) (User, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/users/%d.json", id),
		nil,
	)

	return api.returnSingleUser(response), err
}

func (api *Api) GetUsers() ([]User, error) {
	response, err := api.getHttpRequest(
		"/users.json",
		nil,
	)

	return api.returnMultipleUser(response), err
}

func (api *Api) GetUsersByGroup(groupId string) ([]User, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/groups/%s/users.json", groupId),
		nil,
	)

	return api.returnMultipleUser(response), err
}

func (api *Api) CreateOrUpdateUser(user User) (User, error) {
	var object SingleUser

	_, err := api.postHttpRequest(
		"/users/create_or_update.json",
		map[string]User{"user": user},
		&object,
	)

	return object.Response, err
}

func (api *Api) CreateOrUpdateManyUsers(users []User) (Job, error) {
	var object Job

	_, err := api.postHttpRequest(
		"/users/create_or_update_many.json",
		map[string][]User{"users": users},
		&object,
	)

	return object, err
}

func (api *Api) CreateUser(user User) (User, error) {
	var object SingleUser

	_, err := api.postHttpRequest(
		"/users.json",
		map[string]User{"user": user},
		&object,
	)

	return object.Response, err
}

func (api *Api) UpdateUser(user User) (User, error) {
	var object SingleUser

	_, err := api.updateHttpRequest(
		fmt.Sprintf("/users/%d.json", user.Id),
		map[string]User{"user": user},
		&object,
	)

	return object.Response, err
}

func (api *Api) DeleteUser(id int) (int, error) {
	response, err := api.deleteHttpRequest(
		fmt.Sprintf("/users/%d.json", id),
	)

	return response.StatusCode(), err
}

func (api *Api) returnMultipleUser(response *resty.Response) ([]User) {
	var object MultipleUser

	api.parseResponseToInterface(response, &object)

	return object.Response
}

func (api *Api) returnSingleUser(response *resty.Response) (User) {
	var object SingleUser

	api.parseResponseToInterface(response, &object)

	return object.Response
}

func (users *ManyUsers) AppendUsers(user User) []User {
	users.Users = append(users.Users, user)

	return users.Users
}