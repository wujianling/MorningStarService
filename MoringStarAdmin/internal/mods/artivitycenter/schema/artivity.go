package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Artivity` struct.
type Artivity struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Artivity` struct.
type ArtivityQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Artivity` struct.
type ArtivityQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Artivity` struct.
type ArtivityQueryResult struct {
	Data       Artivities
	PageResult *util.PaginationResult
}

// Defining the slice of `Artivity` struct.
type Artivities []*Artivity

// Defining the data structure for creating a `Artivity` struct.
type ArtivityForm struct {
}

// A validation function for the `ArtivityForm` struct.
func (a *ArtivityForm) Validate() error {
	return nil
}

// Convert `ArtivityForm` to `Artivity` object.
func (a *ArtivityForm) FillTo(artivity *Artivity) error {
	return nil
}
