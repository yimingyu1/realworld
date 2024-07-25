package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"realworld/common/xtools"

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
	userId, ok := l.ctx.Value("UserId").(int64)
	if !ok {
		err = errors.New("用户不存在")
		return
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = errors.New("用户不存在")
		} else {
			logx.Errorf("查询用户失败: %v", err)
			err = errors.New("查询用户失败")
		}
		return
	}
	_, err = l.svcCtx.RedisClient.DelCtx(l.ctx, xtools.GetUserTokenCacheKey(user.Id))
	if err != nil {
		logx.Errorf("删除用户token失败: %v", err)
		err = errors.New("删除用户token失败")
		return
	}
	return "已退出登录", nil
}
