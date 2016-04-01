package zendesk

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type SingleUser struct {
	User User `json:"user"`
}

type MultipleUser struct {
	User []User `json:"users"`
}

func (api *Api) GetUser(id int) (User, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/api/v2/users/%d.json", id),
		nil,
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) GetUsers() ([]User, error) {
	response, err := api.getHttpRequest(
		"/api/v2/users.json",
		nil,
		MultipleUser{},
	)

	var object MultipleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) GetUsersByGroup(groupId string) ([]User, error) {
	response, err := api.getHttpRequest(
		fmt.Sprintf("/api/v2/groups/%s/users.json", groupId),
		nil,
		MultipleUser{},
	)

	var object MultipleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) CreateOrUpdateUser(user User) (User, error) {
	response, err := api.postHttpRequest(
		"/api/v2/users/create_or_update.json",
		map[string]User{"user": user},
		SingleUser{},
	)

	var object = SingleUser{}
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) CreateUser(user User) (User, error) {
	response, err := api.postHttpRequest(
		"/api/v2/users.json",
		map[string]User{"user": user},
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) UpdateUser(user User) (User, error) {
	response, err := api.updateHttpRequest(
		fmt.Sprintf("/api/v2/users/%d.json", user.Id),
		map[string]User{"user": user},
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}

func (api *Api) DeleteUser(id int) (User, error) {
	response, err := api.deleteHttpRequest(
		fmt.Sprintf("/api/v2/users/%d.json", id),
		SingleUser{},
	)

	var object SingleUser
	err = mapstructure.Decode(response, &object)

	if err != nil {
		panic(err)
	}

	return object.User, err;
}