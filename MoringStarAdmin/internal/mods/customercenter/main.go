package customercenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/schema"
	"gorm.io/gorm"
)

type CustomerCenter struct {
	DB          *gorm.DB
	CustomerAPI *api.Customer
}

func (a *CustomerCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Customer))
}

func (a *CustomerCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *CustomerCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	customer := v1.Group("customers")
	{
		customer.GET("", a.CustomerAPI.Query)
		customer.GET(":id", a.CustomerAPI.Get)
		customer.POST("", a.CustomerAPI.Create)
		customer.PUT(":id", a.CustomerAPI.Update)
		customer.DELETE(":id", a.CustomerAPI.Delete)
	}
	return nil
}

func (a *CustomerCenter) Release(ctx context.Context) error {
	return nil
}
