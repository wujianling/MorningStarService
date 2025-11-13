package ordercenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(OrderCenter), "*"),
	wire.Struct(new(dal.Orders), "*"),
	wire.Struct(new(biz.Orders), "*"),
	wire.Struct(new(api.Orders), "*"),
)
