package logic

import (
	"context"

	"MoringStarService/appGateWay/internal/svc"
	"MoringStarService/appGateWay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppGateWayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppGateWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppGateWayLogic {
	return &AppGateWayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppGateWayLogic) AppGateWay(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
