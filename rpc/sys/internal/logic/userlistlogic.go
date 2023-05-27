package logic

import (
	"context"
	"encoding/json"
	"pure-go-zero-admin/rpc/model/sysmodel"

	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {

	query := sysmodel.SysUser{
		Name:     in.Name,
		NickName: in.NickName,
		Email:    in.Email,
		Mobile:   in.Mobile,
		Status:   in.Status,
	}
	all, err := l.svcCtx.UserModel.FindAll(l.ctx, &query, in.Current, in.PageSize)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询用户列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	count, _ := l.svcCtx.UserModel.Count(l.ctx, &query)

	var list []*sys.UserListData
	for _, user := range *all {
		list = append(list, &sys.UserListData{
			Id:             user.Id,
			Name:           user.Name,
			NickName:       user.NickName,
			Avatar:         user.Avatar,
			Password:       user.Password,
			Salt:           user.Salt,
			Email:          user.Email,
			Mobile:         user.Mobile,
			Status:         user.Status,
			CreateBy:       user.CreateBy,
			CreateTime:     user.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   user.LastUpdateBy,
			LastUpdateTime: user.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        user.DelFlag,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询用户列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.UserListResp{
		Total: count,
		List:  list,
	}, nil
	return &sys.UserListResp{}, nil
}
