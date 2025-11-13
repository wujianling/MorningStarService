package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get balance storage instance
func GetBalanceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Balance))
}

// Defining the `Balance` data access object.
type Balance struct {
	DB *gorm.DB
}

// Query balances from the database based on the provided parameters and options.
func (a *Balance) Query(ctx context.Context, params schema.BalanceQueryParam, opts ...schema.BalanceQueryOptions) (*schema.BalanceQueryResult, error) {
	var opt schema.BalanceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetBalanceDB(ctx, a.DB)

	var list schema.Balances
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.BalanceQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified balance from the database.
func (a *Balance) Get(ctx context.Context, id string, opts ...schema.BalanceQueryOptions) (*schema.Balance, error) {
	var opt schema.BalanceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Balance)
	ok, err := util.FindOne(ctx, GetBalanceDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified balance exists in the database.
func (a *Balance) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetBalanceDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new balance.
func (a *Balance) Create(ctx context.Context, item *schema.Balance) error {
	result := GetBalanceDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified balance in the database.
func (a *Balance) Update(ctx context.Context, item *schema.Balance) error {
	result := GetBalanceDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified balance from the database.
func (a *Balance) Delete(ctx context.Context, id string) error {
	result := GetBalanceDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Balance))
	return errors.WithStack(result.Error)
}
