package article

import (
	"context"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfavoriteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消关注文章
func NewUnfavoriteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfavoriteArticleLogic {
	return &UnfavoriteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfavoriteArticleLogic) UnfavoriteArticle(req *types.UnfavoriteArticleReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
