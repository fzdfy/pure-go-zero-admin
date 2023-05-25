package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/rpc/sys/sysclient"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserList godoc
//
//	@Id				UserList
//	@Summary		user list logic
//	@Description	user list
//	@Tags			pure
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			current		query	int64	false	"List User"
//	@Param			pageSize	query	int64	false	"List User"
//	@Param			name		query	string	false	"List User"
//	@Param			nickName	query	string	false	"List User"
//	@Param			mobile		query	string	false	"List User"
//	@Param			email		query	string	false	"List User"
//	@Param			status		query	int64	false	"List User"
//	@Success		200		{object}	types.ListUserResp
//	@Router			/api/sys/user/list [get]
func (l *UserListLogic) UserList(req *types.ListUserReq) (*types.ListUserResp, error) {
	userListReq := sysclient.UserListReq{}
	_ = copier.Copy(&userListReq, &req)
	resp, err := l.svcCtx.Sys.UserList(l.ctx, &userListReq)

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询用户列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询失败")
	}

	//var list []*types.ListUserData = make([]*types.ListUserData, 0)
	var list []*types.ListUserData = []*types.ListUserData{}

	for _, item := range resp.List {
		listUserData := types.ListUserData{}
		_ = copier.Copy(&listUserData, &item)
		list = append(list, &listUserData)
	}

	return &types.ListUserResp{
		//Code:     "000000",
		//Message:  "查询用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
