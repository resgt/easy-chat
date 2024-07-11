/*
@Time : 7/11/2024 11:57 AM
@Author : ZhengXiangy
@File : errors.go
@Software: GoLand
*/
package xerr

import "github.com/zeromicro/x/errors"

// 生成错误
func New(code int, msg string) error {
	return errors.New(code, msg)
}

func NewMsgErr(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

func NewCodeErr(code int) error {
	return errors.New(code, ErrMsg(code))
}

func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}

func NewDBErr() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewReqParamErr() error {
	return errors.New(REQUEST_PARAM_ERROR, ErrMsg(REQUEST_PARAM_ERROR))
}
