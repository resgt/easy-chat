/*
@Time : 7/11/2024 10:40 AM
@Author : ZhengXiangy
@File : registerlogic_test.go
@Software: GoLand
*/
package logic

import (
	"GoLearn/eazy-chat/apps/user/rpc/internal/svc"
	"GoLearn/eazy-chat/apps/user/rpc/user"
	"context"
	"reflect"
	"testing"
)

func TestNewRegisterLogic(t *testing.T) {
	type args struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *RegisterLogic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterLogic(tt.args.ctx, tt.args.svcCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterLogic_Register(t *testing.T) {

	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"1", args{in: &user.RegisterReq{
				Phone:    "13572383042",
				Nickname: "nick",
				Password: "123456",
				Avatar:   "a.jpg",
				Sex:      1,
			}}, true, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				t.Log(tt.name, got)
			}
		})
	}
}
