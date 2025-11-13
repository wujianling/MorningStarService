package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get orders storage instance
func GetOrdersDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Orders))
}

// Defining the `Orders` data access object.
type Orders struct {
	DB *gorm.DB
}

// Query orders from the database based on the provided parameters and options.
func (a *Orders) Query(ctx context.Context, params schema.OrdersQueryParam, opts ...schema.OrdersQueryOptions) (*schema.OrdersQueryResult, error) {
	var opt schema.OrdersQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetOrdersDB(ctx, a.DB)

	var list schema.Orders
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.OrdersQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified orders from the database.
func (a *Orders) Get(ctx context.Context, id string, opts ...schema.OrdersQueryOptions) (*schema.Orders, error) {
	var opt schema.OrdersQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Orders)
	ok, err := util.FindOne(ctx, GetOrdersDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified orders exists in the database.
func (a *Orders) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetOrdersDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new orders.
func (a *Orders) Create(ctx context.Context, item *schema.Orders) error {
	result := GetOrdersDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified orders in the database.
func (a *Orders) Update(ctx context.Context, item *schema.Orders) error {
	result := GetOrdersDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified orders from the database.
func (a *Orders) Delete(ctx context.Context, id string) error {
	result := GetOrdersDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Orders))
	return errors.WithStack(result.Error)
}
