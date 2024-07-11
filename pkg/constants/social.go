/*
@Time : 7/11/2024 5:49 PM
@Author : ZhengXiangy
@File : social.go
@Software: GoLand
*/
package constants

// 定义好友申请状态

type HandlerResult int

const (
	NoHandlerResult HandlerResult = iota + 1
	PassHandlerResult
	RejectHandlerResult
	CancelHandlerResult
)
