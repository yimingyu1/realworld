package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"realworld/common/constant"
	"realworld/common/token"
	"realworld/common/tool"
	"realworld/model"

	"realworld/cmd/api/internal/svc"
	"realworld/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册用户
func NewRegistrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistrationLogic {
	return &RegistrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistrationLogic) Registration(req *types.RegistrationReq) (resp *types.User, err error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		logx.Errorf("find user by email failed: %v", err)
		err = errors.New("find user by email failed")
		return
	}
	if user != nil && user.Email != req.Email {
		err = errors.New("user already exists")
		return
	}
	var insertRest sql.Result
	var userId int64
	if user == nil {
		// 保存用户
		insertRest, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
			UserName: req.UserName,
			Email:    req.Email,
			Password: req.Password,
			Bio:      "",
			DelState: constant.DelStateNo,
		})
		if err != nil {
			logx.Errorf("insert user failed: %v", err)
			err = errors.New("insert user failed")
			return
		}
		userId, _ = insertRest.LastInsertId()
	} else {
		userId = user.Id
	}
	// 生成token
	userToken, _ := token.GenerateToken(userId, l.svcCtx.Config.JwtAuth.AccessSecret, req.UserName, req.Email)
	// 保存token到redis中
	err = l.svcCtx.RedisClient.Setex(tool.GetUserTokenCacheKey(userId), userToken, constant.UserLoginExpire)
	if err != nil {
		logx.Errorf("save user token to redis failed: %v", err)
		err = errors.New("save user token to redis failed")
		return
	}
	return &types.User{
		ID:       userId,
		UserName: req.UserName,
		Email:    req.Email,
		Token:    userToken,
		Bio:      "",
	}, nil
}
