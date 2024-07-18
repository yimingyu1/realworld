package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/logic/profile"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"
)

// 删除文章
func DelArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewDelArticleLogic(r.Context(), svcCtx)
		resp, err := l.DelArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
