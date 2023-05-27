package logic

import (
	"context"
	"encoding/json"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/rpc/sys/sys"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleAdd godoc
//
//	@Id				RoleAdd
//	@Summary		role add logic
//	@Description	role add
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	types.AddRoleReq	 true	"role add"
//	@Success		200		{object}	types.AddRoleResp
//	@Router			/api/sys/role/add [post]
func (l *RoleAddLogic) RoleAdd(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {
	_, err = l.svcCtx.Sys.RoleAdd(l.ctx, &sys.RoleAddReq{
		Name:     req.Name,
		Remark:   req.Remark,
		CreateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加角色信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加角色失败")
	}

	return &types.AddRoleResp{
		Code:    "000000",
		Message: "添加角色成功",
	}, nil
}
