package article

import (
	"context"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 关注文章
func NewFavoriteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteArticleLogic {
	return &FavoriteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteArticleLogic) FavoriteArticle(req *types.FavoriteArticleReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
