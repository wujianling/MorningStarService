package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Shop` business logic.
type Shop struct {
	Trans   *util.Trans
	ShopDAL *dal.Shop
}

// Query shops from the data access object based on the provided parameters and options.
func (a *Shop) Query(ctx context.Context, params schema.ShopQueryParam) (*schema.ShopQueryResult, error) {
	params.Pagination = true

	result, err := a.ShopDAL.Query(ctx, params, schema.ShopQueryOptions{
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

// Get the specified shop from the data access object.
func (a *Shop) Get(ctx context.Context, id string) (*schema.Shop, error) {
	shop, err := a.ShopDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if shop == nil {
		return nil, errors.NotFound("", "Shop not found")
	}
	return shop, nil
}

// Create a new shop in the data access object.
func (a *Shop) Create(ctx context.Context, formItem *schema.ShopForm) (*schema.Shop, error) {
	shop := &schema.Shop{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(shop); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ShopDAL.Create(ctx, shop); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return shop, nil
}

// Update the specified shop in the data access object.
func (a *Shop) Update(ctx context.Context, id string, formItem *schema.ShopForm) error {
	shop, err := a.ShopDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if shop == nil {
		return errors.NotFound("", "Shop not found")
	}

	if err := formItem.FillTo(shop); err != nil {
		return err
	}
	shop.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ShopDAL.Update(ctx, shop); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified shop from the data access object.
func (a *Shop) Delete(ctx context.Context, id string) error {
	exists, err := a.ShopDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Shop not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ShopDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
