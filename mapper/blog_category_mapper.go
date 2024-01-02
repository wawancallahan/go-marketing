package mapper

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type BlogCategoryMapper struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	IsActive    bool        `json:"is_active"`
	Description null.String `json:"description,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
