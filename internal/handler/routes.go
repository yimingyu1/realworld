// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	comment "realworld/cmd/api/internal/handler/comment"
	profile "realworld/cmd/api/internal/handler/profile"
	user "realworld/cmd/api/internal/handler/user"
	"realworld/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware},
			[]rest.Route{
				{
					// 获取评论列表
					Method:  http.MethodGet,
					Path:    "/api/articles/:articleId/comments",
					Handler: comment.GetCommentListHandler(serverCtx),
				},
				{
					// 删除评论
					Method:  http.MethodDelete,
					Path:    "/api/articles/:articleId/comments/:commentId",
					Handler: comment.DeleteCommentHandler(serverCtx),
				},
				{
					// 新增评论
					Method:  http.MethodPost,
					Path:    "/articles/:articleId/comments",
					Handler: comment.AddCommentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware},
			[]rest.Route{
				{
					// 新增tag
					Method:  http.MethodPost,
					Path:    "/tag",
					Handler: comment.AddTagHandler(serverCtx),
				},
				{
					// 获取tagList
					Method:  http.MethodGet,
					Path:    "/tags",
					Handler: comment.GetTagListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware},
			[]rest.Route{
				{
					// 查看用户主页
					Method:  http.MethodGet,
					Path:    "/profiles/:userId",
					Handler: profile.GetProfileHandler(serverCtx),
				},
				{
					// 关注用户
					Method:  http.MethodPost,
					Path:    "/profiles/:userId/follow",
					Handler: profile.FollowUserHandler(serverCtx),
				},
				{
					// 取消关注用户
					Method:  http.MethodDelete,
					Path:    "/profiles/:userId/follow",
					Handler: profile.UnfollowUserHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware},
			[]rest.Route{
				{
					// 获取文章列表
					Method:  http.MethodGet,
					Path:    "/articles",
					Handler: profile.GetArticleListHandler(serverCtx),
				},
				{
					// 创建文章
					Method:  http.MethodPost,
					Path:    "/articles",
					Handler: profile.CreateArticleHandler(serverCtx),
				},
				{
					// 查看文章
					Method:  http.MethodGet,
					Path:    "/articles/:articleId",
					Handler: profile.GetArticleHandler(serverCtx),
				},
				{
					// 更新文章
					Method:  http.MethodPut,
					Path:    "/articles/:articleId",
					Handler: profile.UpdateArticleHandler(serverCtx),
				},
				{
					// 删除文章
					Method:  http.MethodDelete,
					Path:    "/articles/:articleId",
					Handler: profile.DelArticleHandler(serverCtx),
				},
				{
					// 关注文章
					Method:  http.MethodPost,
					Path:    "/articles/:articleId/favorite",
					Handler: profile.FavoriteArticleHandler(serverCtx),
				},
				{
					// 取消关注文章
					Method:  http.MethodDelete,
					Path:    "/articles/:articleId/favorite",
					Handler: profile.UnfavoriteArticleHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 注册用户
				Method:  http.MethodPost,
				Path:    "/user/registration",
				Handler: user.RegistrationHandler(serverCtx),
			},
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/users/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				// 登出
				Method:  http.MethodPost,
				Path:    "/users/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthenticationMiddleware},
			[]rest.Route{
				{
					// 获取用户信息
					Method:  http.MethodGet,
					Path:    "/user",
					Handler: user.GetUserHandler(serverCtx),
				},
				{
					// 更新用户信息
					Method:  http.MethodPut,
					Path:    "/user",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					// 修改密码
					Method:  http.MethodPut,
					Path:    "/user/update-password",
					Handler: user.UpdatePasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}