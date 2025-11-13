package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get product storage instance
func GetProductDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Product))
}

// Defining the `Product` data access object.
type Product struct {
	DB *gorm.DB
}

// Query products from the database based on the provided parameters and options.
func (a *Product) Query(ctx context.Context, params schema.ProductQueryParam, opts ...schema.ProductQueryOptions) (*schema.ProductQueryResult, error) {
	var opt schema.ProductQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetProductDB(ctx, a.DB)

	var list schema.Products
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ProductQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified product from the database.
func (a *Product) Get(ctx context.Context, id string, opts ...schema.ProductQueryOptions) (*schema.Product, error) {
	var opt schema.ProductQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Product)
	ok, err := util.FindOne(ctx, GetProductDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified product exists in the database.
func (a *Product) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetProductDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new product.
func (a *Product) Create(ctx context.Context, item *schema.Product) error {
	result := GetProductDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified product in the database.
func (a *Product) Update(ctx context.Context, item *schema.Product) error {
	result := GetProductDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified product from the database.
func (a *Product) Delete(ctx context.Context, id string) error {
	result := GetProductDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Product))
	return errors.WithStack(result.Error)
}
