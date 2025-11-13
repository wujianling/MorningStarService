package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Artivity` business logic.
type Artivity struct {
	Trans       *util.Trans
	ArtivityDAL *dal.Artivity
}

// Query artivities from the data access object based on the provided parameters and options.
func (a *Artivity) Query(ctx context.Context, params schema.ArtivityQueryParam) (*schema.ArtivityQueryResult, error) {
	params.Pagination = true

	result, err := a.ArtivityDAL.Query(ctx, params, schema.ArtivityQueryOptions{
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

// Get the specified artivity from the data access object.
func (a *Artivity) Get(ctx context.Context, id string) (*schema.Artivity, error) {
	artivity, err := a.ArtivityDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if artivity == nil {
		return nil, errors.NotFound("", "Artivity not found")
	}
	return artivity, nil
}

// Create a new artivity in the data access object.
func (a *Artivity) Create(ctx context.Context, formItem *schema.ArtivityForm) (*schema.Artivity, error) {
	artivity := &schema.Artivity{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(artivity); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ArtivityDAL.Create(ctx, artivity); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return artivity, nil
}

// Update the specified artivity in the data access object.
func (a *Artivity) Update(ctx context.Context, id string, formItem *schema.ArtivityForm) error {
	artivity, err := a.ArtivityDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if artivity == nil {
		return errors.NotFound("", "Artivity not found")
	}

	if err := formItem.FillTo(artivity); err != nil {
		return err
	}
	artivity.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ArtivityDAL.Update(ctx, artivity); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified artivity from the data access object.
func (a *Artivity) Delete(ctx context.Context, id string) error {
	exists, err := a.ArtivityDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Artivity not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.ArtivityDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
