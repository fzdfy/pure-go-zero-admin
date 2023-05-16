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

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.UserAddResp, error) {
	insert, _ := l.svcCtx.UserModel.Insert(l.ctx, &sysmodel.SysUser{
		Name:           in.Name,
		NickName:       in.NickName,
		Avatar:         in.Avatar,
		Password:       "123456",
		Salt:           "123456",
		Email:          in.Email,
		Mobile:         in.Mobile,
		Status:         1,
		DeptId:         in.DeptId,
		CreateBy:       "admin",
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
		DelFlag:        0,
		JobId:          in.JobId,
	})

	insert.LastInsertId()
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

	return &sys.UserAddResp{}, nil
}
