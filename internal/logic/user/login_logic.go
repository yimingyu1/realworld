package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"realworld/common/xtoken"
	"realworld/common/xtools"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewLoginLogic 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			err = errors.New("user not found")
			return
		}
		logx.Errorf("find user by email failed: %v", err)
		err = errors.New("find user by email failed")
		return
	}
	if user.Password != req.Password {
		err = errors.New("password incorrect")
		return
	}
	// 刷新token
	// 生成token
	userToken, _ := xtoken.GenerateToken(user.Id, l.svcCtx.Config.JwtAuth.AccessSecret, user.UserName, req.Email)
	// 保存token到redis中
	err = l.svcCtx.RedisClient.Setex(xtools.GetUserTokenCacheKey(user.Id), userToken, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		logx.Errorf("save user xtoken to redis failed: %v", err)
		err = errors.New("save user xtoken to redis failed")
		return
	}
	return &types.UserResp{
		ID:       user.Id,
		UserName: user.UserName,
		Email:    user.Email,
		Bio:      user.Bio,
	}, nil
}
