package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tag/internal/logic"
	"tag/internal/svc"
	"tag/internal/types"
)

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserDetailLogic(r.Context(), svcCtx)
		err := l.GetUserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
