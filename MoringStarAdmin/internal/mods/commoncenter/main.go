package commoncenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/schema"
	"gorm.io/gorm"
)

type CommonCenter struct {
	DB        *gorm.DB
	CommonAPI *api.Common
}

func (a *CommonCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Common))
}

func (a *CommonCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *CommonCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	common := v1.Group("commons")
	{
		common.GET("", a.CommonAPI.Query)
		common.GET(":id", a.CommonAPI.Get)
		common.POST("", a.CommonAPI.Create)
		common.PUT(":id", a.CommonAPI.Update)
		common.DELETE(":id", a.CommonAPI.Delete)
	}
	return nil
}

func (a *CommonCenter) Release(ctx context.Context) error {
	return nil
}
