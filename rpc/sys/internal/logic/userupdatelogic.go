package logic

import (
	"context"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"time"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.BaseResp, error) {
	err := l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:             in.Id,
		Name:           in.Name,
		NickName:       in.NickName,
		Email:          in.Email,
		Mobile:         in.Mobile,
		LastUpdateBy:   in.LastUpdateBy,
		Status:         in.Status,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return &sys.BaseResp{}, nil
}
