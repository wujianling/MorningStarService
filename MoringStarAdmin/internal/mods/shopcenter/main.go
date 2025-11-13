package shopcenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/schema"
	"gorm.io/gorm"
)

type ShopCenter struct {
	DB      *gorm.DB
	ShopAPI *api.Shop
}

func (a *ShopCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Shop))
}

func (a *ShopCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *ShopCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	shop := v1.Group("shops")
	{
		shop.GET("", a.ShopAPI.Query)
		shop.GET(":id", a.ShopAPI.Get)
		shop.POST("", a.ShopAPI.Create)
		shop.PUT(":id", a.ShopAPI.Update)
		shop.DELETE(":id", a.ShopAPI.Delete)
	}
	return nil
}

func (a *ShopCenter) Release(ctx context.Context) error {
	return nil
}
