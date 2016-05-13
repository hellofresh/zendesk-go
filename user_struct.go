package zendesk

import (
	"github.com/ttacon/libphonenumber"
	"strings"
)

type ManyUsers struct {
	Users []User
}

type User struct {
	Id                  int`json:"id,omitempty"`
	Url                 string `json:"url,omitempty"`
	Name                string `json:"name,omitempty"`
	ExternalId          string `json:"external_id,omitempty"`
	Alias               string `json:"alias,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
	Active              bool `json:"active,omitempty"`
	Verified            bool `json:"verified,omitempty"`
	Shared              bool `json:"shared,omitempty"`
	SharedAgent         bool `json:"shared_agent,omitempty"`
	Locale              string `json:"locale,omitempty"`
	LocaleId            int `json:"locale_id,omitempty"`
	TimeZone            string `json:"time_zone,omitempty"`
	LastLoginAt         string `json:"last_login_at,omitempty"`
	Email               string `json:"email,omitempty"`
	Phone               string `json:"phone,omitempty"`
	Signature           string `json:"signature,omitempty"`
	Details             string `json:"details,omitempty"`
	Notes               string `json:"notes,omitempty"`
	OrganizationId      int `json:"organization_id,omitempty"`
	Role                string `json:"role,omitempty"`
	CustomRoleId        string `json:"custom_role_id,omitempty"`
	Moderator           bool `json:"moderator,omitempty"`
	TicketRestriction   string `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments bool `json:"only_private_comments,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	Suspended           bool `json:"suspended,omitempty"`
	RestrictedAgent     bool `json:"restricted_agent,omitempty"`
	Photo               *Attachment `json:"photo,omitempty"`
}

type Attachment struct{}

func (u *User) FormatPhone(country string) {
	if u.Phone != "" {
		num, _ := libphonenumber.Parse(u.Phone, strings.ToUpper(country))
		u.Phone = libphonenumber.Format(num, libphonenumber.E164)
	}
}

func (users *ManyUsers) AppendUsers(user User) []User {
	users.Users = append(users.Users, user)

	return users.Users
}