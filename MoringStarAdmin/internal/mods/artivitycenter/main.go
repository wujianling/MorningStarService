package artivitycenter

import (
	"context"

	"github.com/wujianling/moringstaradmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/schema"
	"gorm.io/gorm"
)

type ArtivityCenter struct {
	DB          *gorm.DB
	ArtivityAPI *api.Artivity
}

func (a *ArtivityCenter) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(new(schema.Artivity))
}

func (a *ArtivityCenter) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *ArtivityCenter) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {
	artivity := v1.Group("artivities")
	{
		artivity.GET("", a.ArtivityAPI.Query)
		artivity.GET(":id", a.ArtivityAPI.Get)
		artivity.POST("", a.ArtivityAPI.Create)
		artivity.PUT(":id", a.ArtivityAPI.Update)
		artivity.DELETE(":id", a.ArtivityAPI.Delete)
	}
	return nil
}

func (a *ArtivityCenter) Release(ctx context.Context) error {
	return nil
}
