package logic

import (
	"context"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"time"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.BaseResp, error) {
	err := l.svcCtx.RoleModel.Update(l.ctx, &sysmodel.SysRole{
		Id:             in.Id,
		Name:           in.Name,
		Remark:         in.Remark,
		LastUpdateBy:   in.LastUpdateBy,
		Status:         in.Status,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return &sys.BaseResp{}, nil
}
