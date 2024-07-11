package logic

import (
	"GoLearn/eazy-chat/apps/user/models"
	"GoLearn/eazy-chat/apps/user/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/user/rpc/user"
	"GoLearn/eazy-chat/pkg/ctxdata"
	"GoLearn/eazy-chat/pkg/encrypt"
	"GoLearn/eazy-chat/pkg/wuid"
	"GoLearn/eazy-chat/pkg/xerr"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegister = errors.New("手机号码已经注册过！")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	//1. 验证用户是否已经注册 根据手机号码
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil && err != models.ErrNotFound {
		return nil, errors.WithStack(err)
	}

	if userEntity != nil {
		return nil, errors.WithStack(ErrPhoneIsRegister)
	}
	//2. 定义用户数据
	userEntity = &models.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex),
			Valid: true,
		},
	}

	// 如果设置了密码
	if len(in.Password) > 0 {
		genPasswordHash, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		userEntity.Password = sql.NullString{
			String: string(genPasswordHash),
			Valid:  true,
		}
	}
	//3. 插入数据库
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	//4. 设置token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, now+l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "ctxdata get jwt token err %v", err)
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
