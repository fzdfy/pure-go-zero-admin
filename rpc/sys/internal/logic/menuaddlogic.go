package logic

import (
	"context"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"time"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.BaseResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(l.ctx, &sysmodel.SysMenu{
		Id:             0,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Path:           in.Path,
		Perms:          in.Perms,
		Type:           in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		CreateBy:       in.CreateBy,
		CreateTime:     time.Now(),
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.BaseResp{}, nil
}
