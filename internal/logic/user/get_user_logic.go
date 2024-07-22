package user

import (
	"context"
	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser() (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	userID, ok := l.ctx.Value("UserId").(int64)
	if !ok {
		return
	}
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, userID)
	if err != nil {
		return
	}
	logx.Info("sadfasdfasf")
	return &types.UserResp{
		ID:       one.Id,
		UserName: one.UserName,
		Email:    one.Email,
		Bio:      one.Bio,
	}, nil
}
