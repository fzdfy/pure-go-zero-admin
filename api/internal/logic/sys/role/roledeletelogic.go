package logic

import (
	"context"
	"encoding/json"
	"pure-go-zero-admin/api/internal/common/errorx"
	sysclient "pure-go-zero-admin/rpc/sys/sys"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleDelete godoc
//
//	@Id				RoleDelete
//	@Summary		role delete logic
//	@Description	role delete
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	    types.DeleteRoleReq	 true	"role delete"
//	@Success		200		{object}	types.DeleteRoleResp
//	@Router			/api/sys/role/delete [post]
func (l *RoleDeleteLogic) RoleDelete(req *types.DeleteRoleReq) (resp *types.DeleteRoleResp, err error) {
	_, err = l.svcCtx.Sys.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Id:           req.Id,
		LastUpdateBy: "Admin",
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除角色失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除角色失败")
	}

	return &types.DeleteRoleResp{
		Code:    "000000",
		Message: "删除角色成功",
	}, nil
}
