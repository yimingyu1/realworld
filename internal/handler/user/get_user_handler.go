package user

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"realworld/cmd/api/internal/logic/user"
	"realworld/cmd/api/internal/svc"
)

// 获取用户信息
func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
