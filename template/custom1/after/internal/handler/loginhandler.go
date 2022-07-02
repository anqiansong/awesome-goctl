package handler

import (
	"net/http"

	"after/internal/logic"
	"after/internal/svc"
	"after/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			var body Body
			if err != nil {
				body.Code = -1
				body.Msg = err.Error()
			} else {
				body.Msg = "OK"
				body.Data = resp
			}
			httpx.OkJson(w, body)
		}
	}
}
