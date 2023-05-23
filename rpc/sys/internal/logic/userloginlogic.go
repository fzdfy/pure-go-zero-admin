package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"pure-go-zero-admin/rpc/sys/internal/svc"
	"pure-go-zero-admin/rpc/sys/sys"
	"pure-go-zero-admin/util/jwtx"
	"time"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户名和密码登录
func (l *UserLoginLogic) UserLogin(in *sys.UserLoginReq) (*sys.UserLoginResp, error) {
	logx.WithContext(l.ctx).Info("rpc Login")
	logx.Info("rpc Login")
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, err
	}

	if userInfo.Password != in.Password {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &sys.UserLoginResp{
		Status:           "ok",
		CurrentAuthority: "admin",
		Id:               userInfo.Id,
		UserName:         userInfo.Name,
		AccessToken:      jwtToken,
		AccessExpire:     now + accessExpire,
		RefreshAfter:     now + accessExpire/2,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
	return resp, nil
}
