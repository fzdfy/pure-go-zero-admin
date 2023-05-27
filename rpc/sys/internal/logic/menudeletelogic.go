package logic

import (
	"context"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"time"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDeleteLogic) MenuDelete(in *sys.MenuDeleteReq) (*sys.BaseResp, error) {
	err := l.svcCtx.MenuModel.DeleteMenu(l.ctx, &sysmodel.SysMenu{
		Id:             in.Id,
		DelFlag:        -1,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return &sys.BaseResp{}, nil
}
