package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rest/zero/api/internal/logic"
	"go-rest/zero/api/internal/svc"
)

func getAllUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAllUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUsers()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
