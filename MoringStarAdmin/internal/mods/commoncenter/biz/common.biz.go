package biz

import (
	"context"
	"time"

	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/dal"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/errors"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Common` business logic.
type Common struct {
	Trans     *util.Trans
	CommonDAL *dal.Common
}

// Query commons from the data access object based on the provided parameters and options.
func (a *Common) Query(ctx context.Context, params schema.CommonQueryParam) (*schema.CommonQueryResult, error) {
	params.Pagination = true

	result, err := a.CommonDAL.Query(ctx, params, schema.CommonQueryOptions{
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

// Get the specified common from the data access object.
func (a *Common) Get(ctx context.Context, id string) (*schema.Common, error) {
	common, err := a.CommonDAL.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if common == nil {
		return nil, errors.NotFound("", "Common not found")
	}
	return common, nil
}

// Create a new common in the data access object.
func (a *Common) Create(ctx context.Context, formItem *schema.CommonForm) (*schema.Common, error) {
	common := &schema.Common{
		ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if err := formItem.FillTo(common); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommonDAL.Create(ctx, common); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return common, nil
}

// Update the specified common in the data access object.
func (a *Common) Update(ctx context.Context, id string, formItem *schema.CommonForm) error {
	common, err := a.CommonDAL.Get(ctx, id)
	if err != nil {
		return err
	} else if common == nil {
		return errors.NotFound("", "Common not found")
	}

	if err := formItem.FillTo(common); err != nil {
		return err
	}
	common.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommonDAL.Update(ctx, common); err != nil {
			return err
		}
		return nil
	})
}

// Delete the specified common from the data access object.
func (a *Common) Delete(ctx context.Context, id string) error {
	exists, err := a.CommonDAL.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "Common not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.CommonDAL.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	})
}
