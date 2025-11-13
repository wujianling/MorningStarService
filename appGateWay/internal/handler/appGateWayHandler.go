package handler

import (
	"net/http"

	"MoringStarService/appGateWay/internal/logic"
	"MoringStarService/appGateWay/internal/svc"
	"MoringStarService/appGateWay/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AppGateWayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAppGateWayLogic(r.Context(), svcCtx)
		resp, err := l.AppGateWay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
