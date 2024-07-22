package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"net/http"
	"realworld/cmd/api/internal/config"
	"realworld/common/tool"
	"realworld/common/xtoken"
)

type AuthenticationMiddleware struct {
	RedisClient *redis.Redis
	Config      config.Config
}

func NewAuthenticationMiddleware(redisClient *redis.Redis, config config.Config) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		RedisClient: redisClient,
		Config:      config,
	}
}

func (m *AuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析token
		requestToken, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
		if err != nil {
			logx.Errorf("parse token error: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized, err)
			return
		}
		tok, err := jwt.ParseWithClaims(requestToken, &xtoken.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.Config.JwtAuth.AccessSecret), nil
		})
		if err != nil {
			logx.Errorf("parse xtoken error: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized, err)
			return
		}

		if !tok.Valid {
			logx.Errorf("parse xtoken error: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized, err)
			return
		}

		claims, ok := tok.Claims.(*xtoken.Claims)
		if !ok {
			logx.Errorf("parse xtoken error: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized, err)
			return
		}
		tokenFromRedis, err := m.RedisClient.GetCtx(context.Background(), tool.GetUserTokenCacheKey(claims.UserId))
		if err != nil {
			logx.Errorf("get user xtoken error: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusInternalServerError, err)
			return
		}
		if tokenFromRedis == "" || tok.Raw != tokenFromRedis {
			logx.Errorf("xtoken not match: %v", err)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized, errors.New(http.StatusUnauthorized, "登录已过期，请重新登录"))
			return
		}
		ctx := context.WithValue(r.Context(), "UserId", claims.UserId)
		ctx = context.WithValue(ctx, "UserName", claims.UserName)
		ctx = context.WithValue(ctx, "Email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
