package handler

import (
	"net/http"

	"tag/internal/logic"
	"tag/internal/svc"
	"tag/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHeaderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var header types.UserHeader
		if err := httpx.Parse(r, &header); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetHeaderLogic(r.Context(), svcCtx)
		err := l.GetHeader(header)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
