package profile

import (
	"context"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.GetArticleListReq) (resp *types.ArticleList, err error) {
	// todo: add your logic here and delete this line

	return
}
