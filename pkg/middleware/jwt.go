package middleware

import (
	"gin_api_frame/pkg/e"
	"gin_api_frame/pkg/utils/tokenutil"

	"github.com/gin-gonic/gin"

	"time"
)

//JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		var claims *tokenutil.Claims
		if token == "" {
			code = e.ErrorTokenIsNUll
		} else {
			var err error
			claims, err = tokenutil.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		tokenutil.SetTokenClaimsToContext(c, claims)
		c.Next()
	}
}
