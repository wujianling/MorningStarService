package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get common storage instance
func GetCommonDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Common))
}

// Defining the `Common` data access object.
type Common struct {
	DB *gorm.DB
}

// Query commons from the database based on the provided parameters and options.
func (a *Common) Query(ctx context.Context, params schema.CommonQueryParam, opts ...schema.CommonQueryOptions) (*schema.CommonQueryResult, error) {
	var opt schema.CommonQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetCommonDB(ctx, a.DB)

	var list schema.Commons
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.CommonQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified common from the database.
func (a *Common) Get(ctx context.Context, id string, opts ...schema.CommonQueryOptions) (*schema.Common, error) {
	var opt schema.CommonQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Common)
	ok, err := util.FindOne(ctx, GetCommonDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified common exists in the database.
func (a *Common) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetCommonDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new common.
func (a *Common) Create(ctx context.Context, item *schema.Common) error {
	result := GetCommonDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified common in the database.
func (a *Common) Update(ctx context.Context, item *schema.Common) error {
	result := GetCommonDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified common from the database.
func (a *Common) Delete(ctx context.Context, id string) error {
	result := GetCommonDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Common))
	return errors.WithStack(result.Error)
}
