package friend

import (
	"GoLearn/eazy-chat/apps/social/api/internal/svc"
	"GoLearn/eazy-chat/apps/social/api/internal/types"
	"GoLearn/eazy-chat/apps/social/rpc/socialclient"
	"GoLearn/eazy-chat/apps/user/rpc/userclient"
	"GoLearn/eazy-chat/pkg/ctxdata"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUid(l.ctx)

	// 查询好友
	friends, err := l.svcCtx.Social.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return &types.FriendListResp{}, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 保存所有好友id
	uids := make([]string, 0)
	for _, v := range friends.List {
		uids = append(uids, v.FriendUid)
	}

	// 根据uids查询信息
	users, err := l.svcCtx.User.FindUser(l.ctx, &userclient.FindUserReq{
		Ids: uids,
	})
	if err != nil {
		return &types.FriendListResp{}, err
	}
	userRecord := make(map[string]*userclient.UserEntity, len(friends.List))
	for _, user := range users.User {
		userRecord[user.Id] = user
	}
	respList := make([]*types.Friends, 0, len(friends.List))
	for _, v := range friends.List {
		friend := &types.Friends{
			Id:        v.Id,
			FriendUid: v.FriendUid,
		}

		if u, ok := userRecord[v.FriendUid]; ok {
			friend.Nickname = u.Nickname
			friend.Avatar = u.Avatar
		}

		respList = append(respList, friend)
	}
	return &types.FriendListResp{
		List: respList,
	}, nil
}
