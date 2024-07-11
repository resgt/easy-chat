/*
@Time : 7/11/2024 10:36 AM
@Author : ZhengXiangy
@File : login_test.go
@Software: GoLand
*/
package logic

import (
	"GoLearn/eazy-chat/apps/user/rpc/internal/config"
	"GoLearn/eazy-chat/apps/user/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
)

var svcCtx *svc.ServiceContext

func init() {

	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)

}
