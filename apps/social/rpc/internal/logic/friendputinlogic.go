package logic

import (
	"GoLearn/eazy-chat/apps/social/socialmodels"
	"GoLearn/eazy-chat/pkg/constants"
	"GoLearn/eazy-chat/pkg/xerr"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"

	"GoLearn/eazy-chat/apps/social/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送好友请求
func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// todo: add your logic here and delete this line

	// 检查双方是否是好友关系
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friends by uid and fid err %v %v", err, in)
	}
	if friends != nil {
		return &social.FriendPutInResp{}, nil
	}
	// 检查是否已经发送过好友请求
	friendReq, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friendsRequest by rid and uid err %v req %v ", err, in)
	}

	if friendReq != nil {
		return &social.FriendPutInResp{}, nil
	}
	// 构造好友请求
	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &socialmodels.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			String: in.ReqMsg,
			Valid:  true,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert friendRequest err %v req %v ", err, in)
	}
	// 插入数据库
	return &social.FriendPutInResp{}, nil
}
