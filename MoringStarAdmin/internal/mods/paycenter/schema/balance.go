package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Balance` struct.
type Balance struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Balance` struct.
type BalanceQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Balance` struct.
type BalanceQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Balance` struct.
type BalanceQueryResult struct {
	Data       Balances
	PageResult *util.PaginationResult
}

// Defining the slice of `Balance` struct.
type Balances []*Balance

// Defining the data structure for creating a `Balance` struct.
type BalanceForm struct {
}

// A validation function for the `BalanceForm` struct.
func (a *BalanceForm) Validate() error {
	return nil
}

// Convert `BalanceForm` to `Balance` object.
func (a *BalanceForm) FillTo(balance *Balance) error {
	return nil
}
