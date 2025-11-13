package schema

import (
	"time"

	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Product` struct.
type Product struct {
	ID        string    `json:"id" gorm:"size:20;primaryKey;"` // Unique ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
}

// Defining the query parameters for the `Product` struct.
type ProductQueryParam struct {
	util.PaginationParam
}

// Defining the query options for the `Product` struct.
type ProductQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Product` struct.
type ProductQueryResult struct {
	Data       Products
	PageResult *util.PaginationResult
}

// Defining the slice of `Product` struct.
type Products []*Product

// Defining the data structure for creating a `Product` struct.
type ProductForm struct {
}

// A validation function for the `ProductForm` struct.
func (a *ProductForm) Validate() error {
	return nil
}

// Convert `ProductForm` to `Product` object.
func (a *ProductForm) FillTo(product *Product) error {
	return nil
}
