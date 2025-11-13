package paycenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/schema"
	"gorm.io/gorm"
)

type PayCenter struct {
	DB         *gorm.DB
	BalanceAPI *api.Balance
}

func (a *PayCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Balance))
}

func (a *PayCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *PayCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	balance := v1.Group("balances")
	{
		balance.GET("", a.BalanceAPI.Query)
		balance.GET(":id", a.BalanceAPI.Get)
		balance.POST("", a.BalanceAPI.Create)
		balance.PUT(":id", a.BalanceAPI.Update)
		balance.DELETE(":id", a.BalanceAPI.Delete)
	}
	return nil
}

func (a *PayCenter) Release(ctx context.Context) error {
	return nil
}
