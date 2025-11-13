package artivitycenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(ArtivityCenter), "*"),
	wire.Struct(new(dal.Artivity), "*"),
	wire.Struct(new(biz.Artivity), "*"),
	wire.Struct(new(api.Artivity), "*"),
)
