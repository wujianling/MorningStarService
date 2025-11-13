package commoncenter

import (
	"github.com/google/wire"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/api"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/dal"
)

var Set = wire.NewSet(
	wire.Struct(new(CommonCenter), "*"),
	wire.Struct(new(dal.Common), "*"),
	wire.Struct(new(biz.Common), "*"),
	wire.Struct(new(api.Common), "*"),
)
