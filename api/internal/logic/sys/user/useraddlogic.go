package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"
	sys "pure-go-zero-admin/rpc/sys/sysclient"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAddLogic {
	return UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserAdd godoc
//
//	@Id				UserAdd
//	@Summary		user add logic
//	@Description	user add
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	types.AddUserReq	 true	"user add"
//	@Success		200		{object}	types.AddUserResp
//	@Router			/api/sys/user/add [post]
func (l *UserAddLogic) UserAdd(req types.AddUserReq) (*types.AddUserResp, error) {
	_, err := l.svcCtx.Sys.UserAdd(l.ctx, &sys.UserAddReq{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Name:     req.Name,
		NickName: req.NickName,
		CreateBy: "admin",
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加用户失败")
	}

	return &types.AddUserResp{
		Code:    "000000",
		Message: "添加用户成功",
	}, nil
}
