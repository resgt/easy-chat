/*
@Time : 7/11/2024 9:09 AM
@Author : ZhengXiangy
@File : data.go
@Software: GoLand
*/
package ctxdata

import "context"

// 从context中获取uid
func GetUid(ctx context.Context) string {
	if u, ok := ctx.Value(Identify).(string); ok {
		return u
	}
	return ""
}
