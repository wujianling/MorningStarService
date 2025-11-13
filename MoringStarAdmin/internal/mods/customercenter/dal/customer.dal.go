package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get customer storage instance
func GetCustomerDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Customer))
}

// Defining the `Customer` data access object.
type Customer struct {
	DB *gorm.DB
}

// Query customers from the database based on the provided parameters and options.
func (a *Customer) Query(ctx context.Context, params schema.CustomerQueryParam, opts ...schema.CustomerQueryOptions) (*schema.CustomerQueryResult, error) {
	var opt schema.CustomerQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetCustomerDB(ctx, a.DB)

	var list schema.Customers
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.CustomerQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified customer from the database.
func (a *Customer) Get(ctx context.Context, id string, opts ...schema.CustomerQueryOptions) (*schema.Customer, error) {
	var opt schema.CustomerQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Customer)
	ok, err := util.FindOne(ctx, GetCustomerDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified customer exists in the database.
func (a *Customer) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetCustomerDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new customer.
func (a *Customer) Create(ctx context.Context, item *schema.Customer) error {
	result := GetCustomerDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified customer in the database.
func (a *Customer) Update(ctx context.Context, item *schema.Customer) error {
	result := GetCustomerDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified customer from the database.
func (a *Customer) Delete(ctx context.Context, id string) error {
	result := GetCustomerDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Customer))
	return errors.WithStack(result.Error)
}
