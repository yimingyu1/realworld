package profile

import (
	"context"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDelArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelArticleLogic {
	return &DelArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelArticleLogic) DelArticle(req *types.DeleteArticleReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
