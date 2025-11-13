package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Product` business logic.
type Product struct {
	Trans      *util.Trans
	ProductDAL *dal.Product
}

// Query products from the data access object based on the provided parameters and options.
func (a *Product) Query(ctx context.Context, params schema.ProductQueryParam) (*schema.ProductQueryResult, error) {
	params.Pagination = true

	result, err := a.ProductDAL.Query(ctx, params, schema.ProductQueryOptions{
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

// Get the specified product from the data access object.
func (a *Product) Get(ctx context.Context, id string) (*schema.Product, error) {
	product, err := a.ProductDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if product == nil {
		return nil, errors.NotFound("", "Product not found")
	}
	return product, nil
}

// Create a new product in the data access object.
func (a *Product) Create(ctx context.Context, formItem *schema.ProductForm) (*schema.Product, error) {
	product := &schema.Product{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(product); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ProductDAL.Create(ctx, product); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Update the specified product in the data access object.
func (a *Product) Update(ctx context.Context, id string, formItem *schema.ProductForm) error {
	product, err := a.ProductDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if product == nil {
		return errors.NotFound("", "Product not found")
	}

	if err := formItem.FillTo(product); err != nil {
		return err
	}
	product.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ProductDAL.Update(ctx, product); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified product from the data access object.
func (a *Product) Delete(ctx context.Context, id string) error {
	exists, err := a.ProductDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Product not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ProductDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
