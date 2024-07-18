package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"realworld/cmd/api/internal/svc"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
