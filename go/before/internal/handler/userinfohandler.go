package handler

import (
	"net/http"

	"before/internal/logic"
	"before/internal/svc"
	"before/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
