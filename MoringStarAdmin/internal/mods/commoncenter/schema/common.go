package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Common` struct.
type Common struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Common` struct.
type CommonQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Common` struct.
type CommonQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Common` struct.
type CommonQueryResult struct {
	Data       Commons
	PageResult *util.PaginationResult
}

// Defining the slice of `Common` struct.
type Commons []*Common

// Defining the data structure for creating a `Common` struct.
type CommonForm struct {
}

// A validation function for the `CommonForm` struct.
func (a *CommonForm) Validate() error {
	return nil
}

// Convert `CommonForm` to `Common` object.
func (a *CommonForm) FillTo(common *Common) error {
	return nil
}
