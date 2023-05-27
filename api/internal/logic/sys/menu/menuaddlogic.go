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

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuAdd godoc
//
//	@Id				MenuAdd
//	@Summary		menu add logic
//	@Description	menu add
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	types.AddMenuReq	 true	"menu add"
//	@Success		200		{object}	types.AddMenuResp
//	@Router			/api/sys/menu/add [post]
func (l *MenuAddLogic) MenuAdd(req *types.AddMenuReq) (resp *types.AddMenuResp, err error) {

	_, err = l.svcCtx.Sys.MenuAdd(l.ctx, &sys.MenuAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		Path:     req.Path,
		Perms:    req.Perms,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		CreateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加菜单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功",
	}, nil
}
