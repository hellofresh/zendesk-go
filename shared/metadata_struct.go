package shared

import "time"

// Cc struct represent the copy to message
type Cc struct {
	Name         *string `json:"name,omitempty"`
	Address      *string `json:"address,omitempty"`
	OriginalName *string `json:"original_name,omitempty"`
}

// System struct represents the zendesk system info
type System struct {
	IPAddress           *string  `json:"ip_address,omitempty"`
	MessageID           *string  `json:"message_id,omitempty"`
	RawEmailIdentifier  *string  `json:"raw_email_identifier,omitempty"`
	JSONEmailIdentifier *string  `json:"json_email_identifier,omitempty"`
	Ccs                 []Cc     `json:"ccs,omitempty"`
	Client              *string  `json:"client,omitempty"`
	Location            *string  `json:"location,omitempty"`
	Latitude            *float64 `json:"latitude,omitempty"`
	Longitude           *float64 `json:"longitude,omitempty"`
}

// Custom represents SDK metadata
type Custom struct {
	TimeSpent *string            `json:"time_spent,omitempty"`
	SDK       *map[string]string `json:"sdk,omitempty"`
}

// DecorationSource struct
type DecorationSource struct {
	ID        *string `json:"id,omitempty"`
	ZendeskID *int    `json:"zendesk_id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// The Author of data
type Author struct {
	ID       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	PhotoURL *string `json:"photo_url,omitempty"`
	Locale   *string `json:"locale,omitempty"`
}

// Channels of Zendesk
type Channels struct {
	AllowChannelback *bool `json:"allow_channelback,omitempty"`
}

// Decoration struct
type Decoration struct {
	CreatedAt    *time.Time        `json:"created_at,omitempty"`
	Source       *DecorationSource `json:"source,omitempty"`
	ExternalID   *string           `json:"external_id,omitempty"`
	ResourceType *string           `json:"resource_type,omitempty"`
	Version      *int64            `json:"version,omitempty"`
	Author       *Author           `json:"author,omitempty"`
	Type         *string           `json:"type,omitempty"`
	Link         *string           `json:"link,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Channels     *Channels         `json:"channels,omitempty"`
}

// FlagOption struct
type FlagOption struct {
	Message *map[string]string `json:"message,omitempty"`
	Trusted *bool              `json:"trusted,omitempty"`
}

// MetaData struct
type MetaData struct {
	Custom                     *Custom                `json:"custom,omitempty"`
	System                     *System                `json:"system,omitempty"`
	SuspensionTypeID           *string                `json:"suspension_type_id,omitempty"`
	Decoration                 *Decoration            `json:"decoration,omitempty"`
	NotificationsSuppressedFor []string               `json:"notifications_suppressed_for,omitempty"`
	Flags                      []int                  `json:"flags,omitempty"`
	Trusted                    *bool                  `json:"trusted,omitempty"`
	FlagOptions                *map[string]FlagOption `json:"flags_options,omitempty"`
}
