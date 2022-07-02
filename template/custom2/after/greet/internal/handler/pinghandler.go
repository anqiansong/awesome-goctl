package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"greet/internal/logic"
	"greet/internal/svc"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
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
			httpx.Ok(w)
		}
	}
}
