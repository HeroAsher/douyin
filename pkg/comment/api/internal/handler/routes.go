// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	feed "douyin/pkg/comment/api/internal/handler/feed"
	publish "douyin/pkg/comment/api/internal/handler/publish"
	user "douyin/pkg/comment/api/internal/handler/user"
	userOpt "douyin/pkg/comment/api/internal/handler/userOpt"
	"douyin/pkg/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/favorite/action",
					Handler: userOpt.FavoriteOptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/comment/action",
					Handler: userOpt.CommentOptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/relation/action",
					Handler: userOpt.FollowOptHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/favorite/list",
					Handler: userOpt.GetFavoriteListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/comment/list",
					Handler: userOpt.GetCommentListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/relation/follow/list",
					Handler: userOpt.GetFollowListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/relation/follower/list",
					Handler: userOpt.GetFollowerListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthJWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: publish.PublishVideoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: publish.GetPublishVideoListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.IsLogin},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: feed.FeedVideoListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/feed"),
	)
}
