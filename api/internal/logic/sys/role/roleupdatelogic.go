package logic

import (
	"context"
	"encoding/json"
	"pure-go-zero-admin/api/internal/common/errorx"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/rpc/sys/sysclient"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleUpdate godoc
//
//	@Id				RoleUpdate
//	@Summary		role update logic
//	@Description	role update
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	    types.UpdateRoleReq	 true	"role update"
//	@Success		200		{object}	types.UpdateRoleResp
//	@Router			/api/sys/role/update [post]
func (l *RoleUpdateLogic) RoleUpdate(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	_, err = l.svcCtx.Sys.RoleUpdate(l.ctx, &sysclient.RoleUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		Remark:       req.Remark,
		LastUpdateBy: "admin",
		Status:       req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新角色信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新角色失败")
	}

	return &types.UpdateRoleResp{
		Code:    "000000",
		Message: "更新角色成功",
	}, nil
}
