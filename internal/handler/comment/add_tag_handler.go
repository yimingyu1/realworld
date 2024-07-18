package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/logic/comment"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"
)

// 新增tag
func AddTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewAddTagLogic(r.Context(), svcCtx)
		resp, err := l.AddTag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
