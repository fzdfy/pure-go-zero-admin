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

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuDelete godoc
//
//	@Id				MenuDelete
//	@Summary		menu delete logic
//	@Description	menu delete
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	    types.DeleteMenuReq	 true	"menu delete"
//	@Success		200		{object}	types.DeleteMenuResp
//	@Router			/api/sys/menu/delete [post]
func (l *MenuDeleteLogic) MenuDelete(req *types.DeleteMenuReq) (resp *types.DeleteMenuResp, err error) {
	_, err = l.svcCtx.Sys.MenuDelete(l.ctx, &sysclient.MenuDeleteReq{
		Id:           req.Id,
		LastUpdateBy: "Admin",
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除菜单失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除菜单失败")
	}

	return &types.DeleteMenuResp{
		Code:    "000000",
		Message: "删除菜单成功",
	}, nil
}
