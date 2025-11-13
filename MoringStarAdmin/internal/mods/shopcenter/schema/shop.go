package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Shop` struct.
type Shop struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Shop` struct.
type ShopQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Shop` struct.
type ShopQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Shop` struct.
type ShopQueryResult struct {
	Data       Shops
	PageResult *util.PaginationResult
}

// Defining the slice of `Shop` struct.
type Shops []*Shop

// Defining the data structure for creating a `Shop` struct.
type ShopForm struct {
}

// A validation function for the `ShopForm` struct.
func (a *ShopForm) Validate() error {
	return nil
}

// Convert `ShopForm` to `Shop` object.
func (a *ShopForm) FillTo(shop *Shop) error {
	return nil
}
