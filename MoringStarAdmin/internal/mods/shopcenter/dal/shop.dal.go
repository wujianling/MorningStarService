package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get shop storage instance
func GetShopDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Shop))
}

// Defining the `Shop` data access object.
type Shop struct {
	DB *gorm.DB
}

// Query shops from the database based on the provided parameters and options.
func (a *Shop) Query(ctx context.Context, params schema.ShopQueryParam, opts ...schema.ShopQueryOptions) (*schema.ShopQueryResult, error) {
	var opt schema.ShopQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetShopDB(ctx, a.DB)

	var list schema.Shops
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ShopQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified shop from the database.
func (a *Shop) Get(ctx context.Context, id string, opts ...schema.ShopQueryOptions) (*schema.Shop, error) {
	var opt schema.ShopQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Shop)
	ok, err := util.FindOne(ctx, GetShopDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified shop exists in the database.
func (a *Shop) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetShopDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new shop.
func (a *Shop) Create(ctx context.Context, item *schema.Shop) error {
	result := GetShopDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified shop in the database.
func (a *Shop) Update(ctx context.Context, item *schema.Shop) error {
	result := GetShopDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified shop from the database.
func (a *Shop) Delete(ctx context.Context, id string) error {
	result := GetShopDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Shop))
	return errors.WithStack(result.Error)
}
