package logic

import (
	"context"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
