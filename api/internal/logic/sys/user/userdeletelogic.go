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

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserDelete godoc
//
//	@Id				UserDelete
//	@Summary		user delete logic
//	@Description	user delete
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body	    types.DeleteUserReq	 true	"user delete"
//	@Success		200		{object}	types.DeleteUserResp
//	@Router			/api/sys/user/delete [post]
func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	_, err = l.svcCtx.Sys.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Id:           req.Id,
		LastUpdateBy: "Admin",
	})
	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除用户失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除用户失败")
	}

	return &types.DeleteUserResp{
		Code:    "000000",
		Message: "删除用户成功",
	}, nil
}
