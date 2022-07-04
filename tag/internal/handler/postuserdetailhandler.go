package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tag/internal/logic"
	"tag/internal/svc"
	"tag/internal/types"
)

func PostUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostUserDetailLogic(r.Context(), svcCtx)
		err := l.PostUserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
