package zendesk

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
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
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) GetUsers() ([]User, error) {
	response, err := api.getHttpRequest(
		"/users.json",
		nil,
		MultipleUser{},
	)

	var object MultipleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) GetUsersByGroup(groupId string) ([]User, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/groups/%s/users.json", groupId),
		nil,
		MultipleUser{},
	)

	var object MultipleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) CreateOrUpdateUser(user User) (User, error) {
	response, err := api.postHttpRequest(
		"/users/create_or_update.json",
		map[string]User{"user": user},
		SingleUser{},
	)

	var object = SingleUser{}
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) CreateUser(user User) (User, error) {
	response, err := api.postHttpRequest(
		"/users.json",
		map[string]User{"user": user},
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) UpdateUser(user User) (User, error) {
	response, err := api.updateHttpRequest(
		fmt.Sprintf("/users/%d.json", user.Id),
		map[string]User{"user": user},
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.Response, err
}

func (api *Api) DeleteUser(id int) (int, error) {
	response, err := api.deleteHttpRequest(
		fmt.Sprintf("/users/%d.json", id),
	)

	return response.StatusCode(), err
}