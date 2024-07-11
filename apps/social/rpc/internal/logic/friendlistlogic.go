package logic

import (
	"GoLearn/eazy-chat/pkg/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"GoLearn/eazy-chat/apps/social/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	// todo: add your logic here and delete this line
	friendsList, err := l.svcCtx.FriendsModel.ListByUserid(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "list friends by uid err %v, uid %v", err, in.UserId)
	}

	// 将数据库结构转化成go对象结构
	var respList []*social.Friends
	copier.Copy(&respList, &friendsList)
	return &social.FriendListResp{
		List: respList,
	}, nil
}
