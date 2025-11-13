package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Orders` business logic.
type Orders struct {
	Trans     *util.Trans
	OrdersDAL *dal.Orders
}

// Query orders from the data access object based on the provided parameters and options.
func (a *Orders) Query(ctx context.Context, params schema.OrdersQueryParam) (*schema.OrdersQueryResult, error) {
	params.Pagination = true

	result, err := a.OrdersDAL.Query(ctx, params, schema.OrdersQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified orders from the data access object.
func (a *Orders) Get(ctx context.Context, id string) (*schema.Orders, error) {
	orders, err := a.OrdersDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if orders == nil {
		return nil, errors.NotFound("", "Orders not found")
	}
	return orders, nil
}

// Create a new orders in the data access object.
func (a *Orders) Create(ctx context.Context, formItem *schema.OrdersForm) (*schema.Orders, error) {
	orders := &schema.Orders{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(orders); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.OrdersDAL.Create(ctx, orders); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// Update the specified orders in the data access object.
func (a *Orders) Update(ctx context.Context, id string, formItem *schema.OrdersForm) error {
	orders, err := a.OrdersDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if orders == nil {
		return errors.NotFound("", "Orders not found")
	}

	if err := formItem.FillTo(orders); err != nil {
		return err
	}
	orders.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.OrdersDAL.Update(ctx, orders); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified orders from the data access object.
func (a *Orders) Delete(ctx context.Context, id string) error {
	exists, err := a.OrdersDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Orders not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.OrdersDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
