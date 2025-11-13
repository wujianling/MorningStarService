package customercenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(CustomerCenter), "*"),
	wire.Struct(new(dal.Customer), "*"),
	wire.Struct(new(biz.Customer), "*"),
	wire.Struct(new(api.Customer), "*"),
)
