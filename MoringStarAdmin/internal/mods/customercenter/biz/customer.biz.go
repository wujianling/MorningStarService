package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Customer` business logic.
type Customer struct {
	Trans       *util.Trans
	CustomerDAL *dal.Customer
}

// Query customers from the data access object based on the provided parameters and options.
func (a *Customer) Query(ctx context.Context, params schema.CustomerQueryParam) (*schema.CustomerQueryResult, error) {
	params.Pagination = true

	result, err := a.CustomerDAL.Query(ctx, params, schema.CustomerQueryOptions{
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

// Get the specified customer from the data access object.
func (a *Customer) Get(ctx context.Context, id string) (*schema.Customer, error) {
	customer, err := a.CustomerDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if customer == nil {
		return nil, errors.NotFound("", "Customer not found")
	}
	return customer, nil
}

// Create a new customer in the data access object.
func (a *Customer) Create(ctx context.Context, formItem *schema.CustomerForm) (*schema.Customer, error) {
	customer := &schema.Customer{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(customer); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CustomerDAL.Create(ctx, customer); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// Update the specified customer in the data access object.
func (a *Customer) Update(ctx context.Context, id string, formItem *schema.CustomerForm) error {
	customer, err := a.CustomerDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if customer == nil {
		return errors.NotFound("", "Customer not found")
	}

	if err := formItem.FillTo(customer); err != nil {
		return err
	}
	customer.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CustomerDAL.Update(ctx, customer); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified customer from the data access object.
func (a *Customer) Delete(ctx context.Context, id string) error {
	exists, err := a.CustomerDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Customer not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CustomerDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
