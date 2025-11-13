package ordercenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/schema"
	"gorm.io/gorm"
)

type OrderCenter struct {
	DB        *gorm.DB
	OrdersAPI *api.Orders
}

func (a *OrderCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Orders))
}

func (a *OrderCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *OrderCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	orders := v1.Group("orders")
	{
		orders.GET("", a.OrdersAPI.Query)
		orders.GET(":id", a.OrdersAPI.Get)
		orders.POST("", a.OrdersAPI.Create)
		orders.PUT(":id", a.OrdersAPI.Update)
		orders.DELETE(":id", a.OrdersAPI.Delete)
	}
	return nil
}

func (a *OrderCenter) Release(ctx context.Context) error {
	return nil
}
