package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/logic/comment"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"
)

// 新增评论
func AddCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewAddCommentLogic(r.Context(), svcCtx)
		resp, err := l.AddComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
