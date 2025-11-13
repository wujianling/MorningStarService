package mods

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter"
	"github.com/wujianling/moringstaradmin/internal/mods/rbac"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter"
)

const (
	apiPrefix = "/api/"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	rbac.Set,
	customercenter.Set,
	ordercenter.Set,
	commoncenter.Set,
	shopcenter.Set,
	productcenter.Set,
	paycenter.Set,
	artivitycenter.Set,
)

type Mods struct {
	RBAC           *rbac.RBAC
	CustomerCenter *customercenter.CustomerCenter
	OrderCenter    *ordercenter.OrderCenter
	CommonCenter   *commoncenter.CommonCenter
	ShopCenter     *shopcenter.ShopCenter
	ProductCenter  *productcenter.ProductCenter
	PayCenter      *paycenter.PayCenter
	ArtivityCenter *artivitycenter.ArtivityCenter
}

func (a *Mods) Init(ctx context.Context) error {
	if err := a.RBAC.Init(ctx); err != nil {
		return err
	}
	if err := a.CustomerCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.OrderCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.CommonCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.ShopCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.ProductCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.PayCenter.Init(ctx); err != nil {
		return err
	}
	if err := a.ArtivityCenter.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Mods) RouterPrefixes() []string {
	return []string{
		apiPrefix,
	}
}

func (a *Mods) RegisterRouters(ctx context.Context, e *gin.Engine) error {
	gAPI := e.Group(apiPrefix)
	v1 := gAPI.Group("v1")

	if err := a.RBAC.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.CustomerCenter.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.OrderCenter.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.CommonCenter.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.ShopCenter.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}
	if err := a.ProductCenter.RegisterV1Routers(
		ctx,
		v1); err != nil {
		return err
	}
	if err := a.PayCenter.RegisterV1Routers(ctx,

		v1); err != nil {
		return err
	}
	if err := a.ArtivityCenter.RegisterV1Routers(
		ctx, v1); err != nil {
		return err
	}

	return nil
}

func (a *Mods) Release(ctx context.Context) error {
	if err := a.RBAC.Release(ctx); err != nil {
		return err
	}
	if err := a.CustomerCenter.
		Release(ctx); err != nil {
		return err
	}
	if err := a.OrderCenter.
		Release(ctx); err != nil {
		return err
	}
	if err := a.CommonCenter.
		Release(
			ctx); err != nil {
		return err
	}
	if err := a.ShopCenter.
		Release(ctx); err != nil {
		return err
	}
	if err := a.ProductCenter.
		Release(ctx); err != nil {
		return err
	}
	if err := a.PayCenter.Release(ctx); err != nil {
		return err
	}
	if err := a.ArtivityCenter.
		Release(ctx); err != nil {
		return err
	}

	return nil
}
