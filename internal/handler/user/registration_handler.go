package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
	"realworld/cmd/api/internal/logic/user"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"
)

// RegistrationHandler 注册用户
func RegistrationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistrationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegistrationLogic(r.Context(), svcCtx)
		resp, err := l.Registration(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
