package shared

// Attachment info
type Attachment struct {
	*Individual

	ArticleID        *int         `json:"article_id,omitempty"`
	FileName         *string      `json:"file_name,omitempty"`
	ContentURL       *string      `json:"content_url,omitempty"`
	ContentType      *string      `json:"content_type,omitempty"`
	MappedContentURL *string      `json:"mapped_content_url,omitempty"`
	Size             *int         `json:"size,omitempty"`
	Inline           *bool        `json:"inline,omitempty"`
	Thumbnails       []Attachment `json:"thumbnails,omitempty"`
	Width            *int         `json:"width,omitempty"`
	Height           *int         `json:"height,omitempty"`
}
