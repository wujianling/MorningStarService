package shopcenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(ShopCenter), "*"),
	wire.Struct(new(dal.Shop), "*"),
	wire.Struct(new(biz.Shop), "*"),
	wire.Struct(new(api.Shop), "*"),
)
