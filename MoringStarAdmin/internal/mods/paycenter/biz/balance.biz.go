package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Balance` business logic.
type Balance struct {
	Trans      *util.Trans
	BalanceDAL *dal.Balance
}

// Query balances from the data access object based on the provided parameters and options.
func (a *Balance) Query(ctx context.Context, params schema.BalanceQueryParam) (*schema.BalanceQueryResult, error) {
	params.Pagination = true

	result, err := a.BalanceDAL.Query(ctx, params, schema.BalanceQueryOptions{
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

// Get the specified balance from the data access object.
func (a *Balance) Get(ctx context.Context, id string) (*schema.Balance, error) {
	balance, err := a.BalanceDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if balance == nil {
		return nil, errors.NotFound("", "Balance not found")
	}
	return balance, nil
}

// Create a new balance in the data access object.
func (a *Balance) Create(ctx context.Context, formItem *schema.BalanceForm) (*schema.Balance, error) {
	balance := &schema.Balance{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(balance); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.BalanceDAL.Create(ctx, balance); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// Update the specified balance in the data access object.
func (a *Balance) Update(ctx context.Context, id string, formItem *schema.BalanceForm) error {
	balance, err := a.BalanceDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if balance == nil {
		return errors.NotFound("", "Balance not found")
	}

	if err := formItem.FillTo(balance); err != nil {
		return err
	}
	balance.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.BalanceDAL.Update(ctx, balance); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified balance from the data access object.
func (a *Balance) Delete(ctx context.Context, id string) error {
	exists, err := a.BalanceDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Balance not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.BalanceDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
