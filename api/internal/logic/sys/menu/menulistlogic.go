package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"
	"pure-go-zero-admin/rpc/sys/sysclient"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuListLogic {
	return MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MenuList godoc
//
//	@Id				MenuList
//	@Summary		menu list logic
//	@Description	menu list
//	@Tags			pure
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"List Menu"
//	@Param			path	query		string	false	"List Menu"
//	@Param			status		query	int64	false	"List Menu"
//	@Success		200		{object}	types.ListMenuResp
//	@Router			/api/sys/menu/list [get]
func (l *MenuListLogic) MenuList(req types.ListMenuReq) (*types.ListMenuResp, error) {
	menuListReq := sysclient.MenuListReq{}
	_ = copier.Copy(&menuListReq, &req)
	resp, err := l.svcCtx.Sys.MenuList(l.ctx, &menuListReq)

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询菜单列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询菜单失败")
	}

	var list []*types.ListMenuData = []*types.ListMenuData{}

	for _, menu := range resp.List {
		list = append(list, &types.ListMenuData{
			Id:             menu.Id,
			Name:           menu.Name,
			ParentId:       menu.ParentId,
			Path:           menu.Path,
			Perms:          menu.Perms,
			Type:           menu.Type,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime,
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime,
			DelFlag:        menu.DelFlag,
		})
	}

	return &types.ListMenuResp{
		Data:    list,
		Success: true,
		Total:   resp.Total,
	}, nil
}
