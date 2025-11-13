package dal

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
	"gorm.io/gorm"
)

// Get artivity storage instance
func GetArtivityDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(schema.Artivity))
}

// Defining the `Artivity` data access object.
type Artivity struct {
	DB *gorm.DB
}

// Query artivities from the database based on the provided parameters and options.
func (a *Artivity) Query(ctx context.Context, params schema.ArtivityQueryParam, opts ...schema.ArtivityQueryOptions) (*schema.ArtivityQueryResult, error) {
	var opt schema.ArtivityQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetArtivityDB(ctx, a.DB)

	var list schema.Artivities
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.ArtivityQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified artivity from the database.
func (a *Artivity) Get(ctx context.Context, id string, opts ...schema.ArtivityQueryOptions) (*schema.Artivity, error) {
	var opt schema.ArtivityQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.Artivity)
	ok, err := util.FindOne(ctx, GetArtivityDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists checks if the specified artivity exists in the database.
func (a *Artivity) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetArtivityDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new artivity.
func (a *Artivity) Create(ctx context.Context, item *schema.Artivity) error {
	result := GetArtivityDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified artivity in the database.
func (a *Artivity) Update(ctx context.Context, item *schema.Artivity) error {
	result := GetArtivityDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified artivity from the database.
func (a *Artivity) Delete(ctx context.Context, id string) error {
	result := GetArtivityDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.Artivity))
	return errors.WithStack(result.Error)
}
