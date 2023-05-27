package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/rpc/model/sysmodel"
	"pure-go-zero-admin/rpc/sys/sys"
	"time"

	//"pure-go-zero-admin/api/internal/svc"
	//"pure-go-zero-admin/api/internal/types"
	"pure-go-zero-admin/rpc/sys/internal/svc"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.BaseResp, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &sysmodel.SysUser{
		Name:           in.Name,
		NickName:       in.NickName,
		Email:          in.Email,
		Mobile:         in.Mobile,
		Status:         1,
		CreateBy:       "admin",
		CreateTime:     time.Now(),
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
	})

	//insert.LastInsertId()
	//id, _ := insert.LastInsertId()
	//_ = l.svcCtx.UserRoleModel.Delete(id)
	//
	//_, _ = l.svcCtx.UserRoleModel.Insert(sysmodel.SysUserRole{
	//	UserId:         id,
	//	RoleId:         in.RoleId,
	//	CreateBy:       "admin",
	//	CreateTime:     time.Now(),
	//	LastUpdateBy:   "admin",
	//	LastUpdateTime: time.Now(),
	//})

	if err != nil {
		return nil, err
	}

	return &sys.BaseResp{}, nil
}
