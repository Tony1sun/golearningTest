package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"hello/internal/logic"
	"hello/internal/svc"
	"hello/internal/types"
)

func authorizationHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAuthorizationLogic(r.Context(), ctx)
		resp, err := l.Authorization(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
