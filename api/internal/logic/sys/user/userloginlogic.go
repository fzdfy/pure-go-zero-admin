package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/common/prefix"
	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"
	"pure-go-zero-admin/rpc/sys/sysclient"
	"strings"
	"time"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserLogin godoc
//
//	@Id				UserLogin
//	@Summary		user login logic
//	@Description	user login
//	@Tags			pure
//	@Accept			json
//	@Produce		json
//	@Param			body	body		types.LoginReq	true	"Login info"
//	@Success		200		{object}	types.LoginResp
//	@Router			/api/sys/user/login [post]
func (l *UserLoginLogic) UserLogin(req types.LoginReq, ip string) (*types.LoginResp, error) {
	logx.Info("UserLogin")
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errorx.NewDefaultError("用户名或密码不能为空")
	}

	resp, err := l.svcCtx.Sys.UserLogin(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, errorx.NewDefaultError("登录失败")
	}

	//l.svcCtx.Cache.ge
	////保存登录日志
	//_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
	//	UserName: resp.UserName,
	//	Status:   "login",
	//	Ip:       ip,
	//	CreateBy: resp.UserName,
	//})

	// 缓存当前用户的token数据
	cacheKey := fmt.Sprintf("%s%v", prefix.CacheSysUserTokenPrefix, resp.Id)
	times := time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)
	_ = l.svcCtx.Cache.SetWithExpire(cacheKey, resp.AccessToken, times)

	return &types.LoginResp{
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
		Id:               resp.Id,
		UserName:         resp.UserName,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
