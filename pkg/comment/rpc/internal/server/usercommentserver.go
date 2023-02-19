// Code generated by goctl. DO NOT EDIT.
// Source: UserComment.proto

package server

import (
	"context"

	"douyin/pkg/comment/rpc/internal/logic"
	"douyin/pkg/comment/rpc/internal/svc"
	"douyin/pkg/comment/rpc/userCommentPb"
)

type UserCommentServer struct {
	svcCtx *svc.ServiceContext
	userCommentPb.UnimplementedUserCommentServer
}

func NewUserCommentServer(svcCtx *svc.ServiceContext) *UserCommentServer {
	return &UserCommentServer{
		svcCtx: svcCtx,
	}
}

// -----------------------userCommentStatus-----------------------
func (s *UserCommentServer) UpdateCommentStatus(ctx context.Context, in *userCommentPb.UpdateCommentStatusReq) (*userCommentPb.UpdateCommentStatusResp, error) {
	l := logic.NewUpdateCommentStatusLogic(ctx, s.svcCtx)
	return l.UpdateCommentStatus(in)
}

// -----------------------userCommentList-----------------------
func (s *UserCommentServer) GetVideoComment(ctx context.Context, in *userCommentPb.GetVideoCommentReq) (*userCommentPb.GetVideoCommentReqResp, error) {
	l := logic.NewGetVideoCommentLogic(ctx, s.svcCtx)
	return l.GetVideoComment(in)
}
