package logic

import (
	"context"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"time"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.BaseResp, error) {
	_, err := l.svcCtx.RoleModel.Insert(l.ctx, &sysmodel.SysRole{
		Id:             0,
		Name:           in.Name,
		Remark:         in.Remark,
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
