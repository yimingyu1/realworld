package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"realworld/cmd/api/internal/logic/comment"
	"realworld/cmd/api/internal/svc"
)

// 获取tagList
func GetTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := comment.NewGetTagListLogic(r.Context(), svcCtx)
		resp, err := l.GetTagList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
