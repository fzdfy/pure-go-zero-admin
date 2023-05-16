package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	logic "pure-go-zero-admin/api/internal/logic/sys/menu"
	"pure-go-zero-admin/api/internal/svc"
	"pure-go-zero-admin/api/internal/types"
)

func MenuListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMenuListLogic(r.Context(), ctx)
		resp, err := l.MenuList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
