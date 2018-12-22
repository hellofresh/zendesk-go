package models

// Photo struct represents an image
type Photo struct {
	ID               *int    `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	ContentURL       *string `json:"content_url,omitempty"`
	ContentType      *string `json:"content_type,omitempty"`
	Size             *int    `json:"size,omitempty"`
	Thumbnails       []Photo `json:"thumbnails,omitempty"`
	URL              *string `json:"url,omitempty"`
	FileName         *string `json:"file_name,omitempty"`
	MappedContentURL *string `json:"mapped_content_url,omitempty"`
	Width            *int    `json:"width,omitempty"`
	Height           *int    `json:"height,omitempty"`
	Inline           *bool   `json:"inline,omitempty"`
}
