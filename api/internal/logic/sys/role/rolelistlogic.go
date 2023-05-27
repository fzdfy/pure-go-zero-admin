package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"pure-go-zero-admin/api/internal/common/errorx"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/rpc/sys/sysclient"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleList godoc
//
//	@Id				RoleList
//	@Summary		role list logic
//	@Description	role list
//	@Tags			pure
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"list role"
//	@Param			status		query	int64	false	"list role"
//	@Success		200		{object}	types.ListRoleResp
//	@Router			/api/sys/role/list [get]
func (l *RoleListLogic) RoleList(req *types.ListRoleReq) (*types.ListRoleResp, error) {
	roleListReq := sysclient.RoleListReq{}
	_ = copier.Copy(&roleListReq, &req)
	resp, err := l.svcCtx.Sys.RoleList(l.ctx, &roleListReq)

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询角色列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询角色失败")
	}

	var list []*types.ListRoleData = []*types.ListRoleData{}

	for _, role := range resp.List {
		list = append(list, &types.ListRoleData{
			Id:             role.Id,
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreateTime,
			LastUpdateBy:   role.LastUpdateBy,
			LastUpdateTime: role.LastUpdateTime,
			DelFlag:        role.DelFlag,
			Status:         role.Status,
		})
	}

	return &types.ListRoleResp{
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
