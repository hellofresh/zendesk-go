package models

import (
	"strings"
	"time"

	"github.com/ttacon/libphonenumber"

	"../shared"
)

// ManyUsers struct is used to aggregate some users
type ManyUsers struct {
	Users []User
}

// User struct model represents response from api
type User struct {
	*shared.Individual

	Name                 *string                 `json:"name,omitempty"`
	ExternalID           *string                 `json:"external_id,omitempty"`
	Alias                *string                 `json:"alias,omitempty"`
	Active               *bool                   `json:"active,omitempty"`
	Verified             *bool                   `json:"verified,omitempty"`
	Shared               *bool                   `json:"shared,omitempty"`
	LocaleID             *int                    `json:"locale_id,omitempty"`
	TimeZone             *string                 `json:"time_zone,omitempty"`
	IanaTimeZone         *string                 `json:"iana_time_zone,omitempty"`
	LastLoginAt          *time.Time              `json:"last_login_at,omitempty"`
	Email                *string                 `json:"email,omitempty"`
	Phone                *string                 `json:"phone,omitempty"`
	Signature            *string                 `json:"signature,omitempty"`
	Details              *string                 `json:"details,omitempty"`
	Notes                *string                 `json:"notes,omitempty"`
	OrganizationID       *int                    `json:"organization_id,omitempty"`
	Role                 *string                 `json:"role,omitempty"`
	CustomRoleID         *int                    `json:"custom_role_id,omitempty"`
	Moderator            *bool                   `json:"moderator,omitempty"`
	TicketRestriction    *string                 `json:"ticket_restriction,omitempty"`
	OnlyPrivateComments  *bool                   `json:"only_private_comments,omitempty"`
	Tags                 []string                `json:"tags,omitempty"`
	Suspended            *bool                   `json:"suspended,omitempty"`
	Photo                *Photo                  `json:"photo,omitempty"`
	RemotePhotoURL       *string                 `json:"remote_photo_url,omitempty"`
	UserFields           *map[string]interface{} `json:"user_fields,omitempty"`
	Identities           []UserIdentity          `json:"identities,omitempty"`
	SharedPhoneNumber    *bool                   `json:"shared_phone_number,omitempty"`
	Locale               *string                 `json:"locale,omitempty"`
	SharedAgent          *bool                   `json:"shared_agent,omitempty"`
	TwoFactorAuthEnabled *bool                   `json:"two_factor_auth_enabled,omitempty"`
	RoleType             *string                 `json:"role_type,omitempty"`
	RestrictedAgent      *bool                   `json:"restricted_agent,omitempty"`
	ChatOnly             *bool                   `json:"chat_only,omitempty"`
	DefaultGroupID       *string                 `json:"default_group_id,omitempty"`
	ReportCsv            *bool                   `json:"report_csv,omitempty"`
}

// UserIdentity struct represents an user unique identifier
type UserIdentity struct {
	ID        *int       `json:"id,omitempty"`
	UserID    *int       `json:"user_id,omitempty"`
	Type      *string    `json:"type,omitempty"`
	Value     *string    `json:"value,omitempty"`
	Verified  *bool      `json:"verified,omitempty"`
	Primary   *bool      `json:"primary,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// FormatPhone convert a string to formatted phone representation
func (u *User) FormatPhone(country string) {
	if u.Phone != nil {
		num, _ := libphonenumber.Parse(*u.Phone, strings.ToUpper(country))
		*u.Phone = libphonenumber.Format(num, libphonenumber.E164)
	}
}

// AppendUsers concat a user with many other users
func (users *ManyUsers) AppendUsers(user User) []User {
	users.Users = append(users.Users, user)

	return users.Users
}

// ContcatUsers function concats many users with many users
func (users *ManyUsers) ContcatUsers(manyUsers []User) []User {
	users.Users = append(users.Users, manyUsers...)

	return users.Users
}
