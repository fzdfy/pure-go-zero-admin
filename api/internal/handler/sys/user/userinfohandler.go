package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"pure-go-zero-admin/api/internal/logic/sys/user"
	"pure-go-zero-admin/api/internal/svc"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
