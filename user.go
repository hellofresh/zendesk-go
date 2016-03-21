package zendesk

import (
	"golang.org/x/net/context"
	"fmt"
)

type Attachment struct {}

type User struct {
	Id                  int                    `json:"id,omitempty"`
	Url                 string                 `json:"url,omitempty"`
	Name                string                 `json:"name,omitempty"`
	ExternalId          string                 `json:"external_id,omitempty"`
	Alias               string                 `json:"alias,omitempty"`
	CreatedAt           string                 `json:"created_at,omitempty"`
	UpdatedAt           string                 `json:"updated_at,omitempty"`
	Active              bool                   `json:"active,omitempty"`
	Verified            bool                   `json:"verified,omitempty"`
	Shared              bool                   `json:"shared,omitempty"`
	SharedAgent         bool                   `json:"shared_agent,omitempty"`
	Locale              string                 `json:"locale,omitempty"`
	LocaleId            int                    `json:"locale_id,omitempty"`
	TimeZone            string                 `json:"time_zone,omitempty"`
	LastLoginAt         string                 `json:"last_login_at,omitempty"`
	Email               string                 `json:"email,omitempty"`
	Phone               string                 `json:"phone,omitempty"`
	Signature           string                 `json:"signature,omitempty"`
	Details             string                 `json:"details,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	OrganizationId      int                    `json:"organization_id,omitempty"`
	Role                string                 `json:"role,omitempty"`
	CustomRoleId        string                 `json:"custom_role_id,omitempty"`
	Moderator           bool                   `json:"moderator,omitempty"`
	TicketRestriction   string                 `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments bool                   `json:"only_private_comments,omitempty"`
	Tags                []string               `json:"tags,omitempty"`
	Suspended           bool                   `json:"suspended,omitempty"`
	RestrictedAgent     bool                   `json:"restricted_agent,omitempty"`
	Photo               *Attachment            `json:"photo,omitempty"`
	UserFields          map[string]interface{} `json:"user_fields,omitempty"`
}

type UserApi struct {
	client  *Client
	context context.Context
}

func (api *UserApi) getUser(path string, params map[string]string) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	_, err := api.client.get(path, params, &response)

	if err != nil {
		return User{}, err
	}

	return response.User, nil
}

func (api *UserApi) getUsers(path string, params map[string]string) ([]User, error) {
	response := struct {
		User []User `json:"users"`
	}{}

	_, err := api.client.get(path, params, &response)

	if err != nil {
		return nil, err
	}

	return response.User, nil
}

func (api *UserApi) postUser(path string, payload interface{}) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	_, err := api.client.post(path, payload, &response)

	if err != nil {
		return User{}, err
	}

	return response.User, nil
}

func (api *UserApi) deleteUser(path string) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	_, err := api.client.delete(path, &response)

	if err != nil {
		return User{}, err
	}

	return response.User, nil
}

func (api *UserApi) updateUser(path string, payload interface{}) (User, error) {
	response := struct {
		User User `json:"user"`
	}{}

	_, err := api.client.put(path, payload, &response)

	if err != nil {
		return User{}, err
	}

	return response.User, nil
}

//

func (api *UserApi) GetUser(id int) (User, error) {
	return api.getUser(fmt.Sprintf("/api/v2/users/%d.json", id), nil)
}

func (api *UserApi) GetUsers() ([]User, error) {
	return api.getUsers("/api/v2/users.json", nil)
}

func (api *UserApi) CreateOrUpdateUser(user User) (User, error) {
	return api.postUser("/api/v2/users/create_or_update.json", user)
}

func (api *UserApi) CreateUser(user User) (User, error) {
	return api.postUser("/api/v2/users.json", user)
}

func (api *UserApi) UpdateUser(user User) (User, error) {
	return api.updateUser(fmt.Sprintf("/api/v2/users/%d.json", user.Id), user)
}

func (api *UserApi) DeleteUser(id int) (User, error) {
	return api.deleteUser(fmt.Sprintf("/api/v2/users/%d.json", id))
}