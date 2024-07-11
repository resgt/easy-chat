/*
@Time : 7/11/2024 9:09 AM
@Author : ZhengXiangy
@File : token.go
@Software: GoLand
*/
package ctxdata

import "github.com/golang-jwt/jwt/v4"

const Identify = "imooc.com"

// 密钥 生成时间 到期时间 uid
func GetJwtToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
