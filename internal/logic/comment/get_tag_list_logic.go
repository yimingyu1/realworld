package comment

import (
	"context"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取tagList
func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagListLogic) GetTagList() (resp []types.Tag, err error) {
	// todo: add your logic here and delete this line

	return
}
