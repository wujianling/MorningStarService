package paycenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(PayCenter), "*"),
	wire.Struct(new(dal.Balance), "*"),
	wire.Struct(new(biz.Balance), "*"),
	wire.Struct(new(api.Balance), "*"),
)
