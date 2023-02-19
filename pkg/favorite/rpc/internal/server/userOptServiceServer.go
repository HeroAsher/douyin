// Code generated by goctl. DO NOT EDIT.
// Source: UserOptService.proto

package server

import (
	"context"

	"douyin/pkg/favorite/rpc/internal/logic"
	"douyin/pkg/favorite/rpc/internal/svc"
	"douyin/pkg/favorite/rpc/userOptPb"
)

type UserOptServiceServer struct {
	svcCtx *svc.ServiceContext
	userOptPb.UnimplementedUserOptServiceServer
}

func NewUserOptServiceServer(svcCtx *svc.ServiceContext) *UserOptServiceServer {
	return &UserOptServiceServer{
		svcCtx: svcCtx,
	}
}

// -----------------------userFavoriteList-----------------------
func (s *UserOptServiceServer) GetUserFavorite(ctx context.Context, in *userOptPb.GetUserFavoriteReq) (*userOptPb.GetUserFavoriteResp, error) {
	l := logic.NewGetUserFavoriteLogic(ctx, s.svcCtx)
	return l.GetUserFavorite(in)
}

func (s *UserOptServiceServer) UpdateFavoriteStatus(ctx context.Context, in *userOptPb.UpdateFavoriteStatusReq) (*userOptPb.UpdateFavoriteStatusResp, error) {
	l := logic.NewUpdateFavoriteStatusLogic(ctx, s.svcCtx)
	return l.UpdateFavoriteStatus(in)
}