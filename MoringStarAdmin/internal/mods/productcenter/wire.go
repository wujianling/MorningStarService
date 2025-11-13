package productcenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(ProductCenter), "*"),
	wire.Struct(new(dal.Product), "*"),
	wire.Struct(new(biz.Product), "*"),
	wire.Struct(new(api.Product), "*"),
)
