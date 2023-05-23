package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"pure-go-zero-admin/api/internal/logic/sys/user"
	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"
)

func UserLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), ctx)
		resp, err := l.UserLogin(req, r.RemoteAddr)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
