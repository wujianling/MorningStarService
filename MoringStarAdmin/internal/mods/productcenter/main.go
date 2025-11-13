package productcenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/schema"
	"gorm.io/gorm"
)

type ProductCenter struct {
	DB         *gorm.DB
	ProductAPI *api.Product
}

func (a *ProductCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Product))
}

func (a *ProductCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *ProductCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	product := v1.Group("products")
	{
		product.GET("", a.ProductAPI.Query)
		product.GET(":id", a.ProductAPI.Get)
		product.POST("", a.ProductAPI.Create)
		product.PUT(":id", a.ProductAPI.Update)
		product.DELETE(":id", a.ProductAPI.Delete)
	}
	return nil
}

func (a *ProductCenter) Release(ctx context.Context) error {
	return nil
}
