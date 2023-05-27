package logic

import (
	"context"
	"encoding/json"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/rpc/sys/sysclient"

	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuUpdate godoc
//
//	@Id				MenuUpdate
//	@Summary		menu update logic
//	@Description	menu update
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	    types.UpdateMenuReq	 true	"menu update"
//	@Success		200		{object}	types.UpdateMenuResp
//	@Router			/api/sys/menu/update [post]
func (l *MenuUpdateLogic) MenuUpdate(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	_, err = l.svcCtx.Sys.MenuUpdate(l.ctx, &sysclient.MenuUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		ParentId:     req.ParentId,
		Path:         req.Path,
		Perms:        req.Perms,
		Type:         req.Type,
		Icon:         req.Icon,
		OrderNum:     req.OrderNum,
		Status:       req.Status,
		LastUpdateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新菜单失败")
	}

	return &types.UpdateMenuResp{
		Code:    "000000",
		Message: "更新菜单成功",
	}, nil
}
