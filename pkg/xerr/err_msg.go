/*
@Time : 7/11/2024 11:52 AM
@Author : ZhengXiangy
@File : err_msg.go
@Software: GoLand
*/
package xerr

var codeText = map[int]string{
	SERVER_COMMON_ERROR: "服务异常，请稍后处理",
	REQUEST_PARAM_ERROR: "参数不正确",
	TOKEN_EXPIRE_ERROR:  "token失效，请重新登陆",
	DB_ERROR:            "数据库繁忙，请稍后再试",
}

// 获取错误码对应的错误信息
func ErrMsg(code int) string {
	if msg, ok := codeText[code]; ok {
		return msg
	}
	return codeText[SERVER_COMMON_ERROR]
}
