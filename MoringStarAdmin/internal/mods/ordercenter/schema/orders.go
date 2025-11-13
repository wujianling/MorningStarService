package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Orders` struct.
type Orders struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Orders` struct.
type OrdersQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Orders` struct.
type OrdersQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Orders` struct.
type OrdersQueryResult struct {
	Data       Orders
	PageResult *util.PaginationResult
}

// Defining the slice of `Orders` struct.
type Orders []*Orders

// Defining the data structure for creating a `Orders` struct.
type OrdersForm struct {
}

// A validation function for the `OrdersForm` struct.
func (a *OrdersForm) Validate() error {
	return nil
}

// Convert `OrdersForm` to `Orders` object.
func (a *OrdersForm) FillTo(orders *Orders) error {
	return nil
}
