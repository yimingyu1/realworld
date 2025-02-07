package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"realworld/cmd/api/internal/config"
	"realworld/cmd/api/internal/middleware"
	"realworld/model"
)

type ServiceContext struct {
	Config                   config.Config
	AuthenticationMiddleware rest.Middleware
	UserModel                model.UserModel
	ArticleModel             model.ArticleModel
	RedisClient              *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.DB.DataSource)
	redisClient := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
	})
	return &ServiceContext{
		Config:                   c,
		AuthenticationMiddleware: middleware.NewAuthenticationMiddleware(redisClient, c).Handle,
		UserModel:                model.NewUserModel(mysqlConn),
		ArticleModel:             model.NewArticleModel(mysqlConn),
		RedisClient:              redisClient,
	}
}
