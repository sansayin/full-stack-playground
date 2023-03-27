package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rest/api/internal/logic"
	"go-rest/api/internal/svc"
)

func getUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetUsers()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
