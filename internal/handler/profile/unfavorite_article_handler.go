package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/logic/profile"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"
)

// 取消关注文章
func UnfavoriteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnfavoriteArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewUnfavoriteArticleLogic(r.Context(), svcCtx)
		resp, err := l.UnfavoriteArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
