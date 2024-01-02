package mapper

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type BlogBannerMapper struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	FileName  null.String `json:"file_name,omitempty"`
	Path      null.String `json:"path,omitempty"`
	Url       null.String `json:"url,omitempty"`
	MimeType  null.String `json:"mime_type,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
